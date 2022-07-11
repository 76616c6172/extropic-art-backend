package main

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"project-exia-monorepo/website/exapi"
	"project-exia-monorepo/website/exdb"
	"project-exia-monorepo/website/exutils"

	"github.com/google/uuid"
)

const WEBSERVER_PORT = ":8091"  // Scheduler is listening on this port
const COLAB_TEST_WORKER = false //debug
const NGROK_IP = "8496-35-201-156-213.ngrok.io"

var WORKERDB *sql.DB //pointer used to connect to the db, initialized in main
var JOB_COMPLETE = true
var JOBDB *sql.DB
var SECRET string
var PANIC bool

// Initialize and connect to workerdb
// GPU_WORERS register themselves with the scheduler through this endpoint
// /api/0/registration
func handleWorkerRegistration(w http.ResponseWriter, r *http.Request) {
	println("Receiving new worker registration..")

	registrationSecret, err := r.Cookie("secret")
	if err != nil {
		log.Println("Error reading secret cookie: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if registrationSecret.Value == SECRET {
		println("Authenticated new worker..")
		newWorkerId := uuid.New().String()
		log.Println("New worker ID created:", newWorkerId)

		workerIpForDB := strings.Split(r.RemoteAddr, ":") // Parse the ip and edit the port
		err := exdb.RegisterNewWorker(WORKERDB, newWorkerId, workerIpForDB[0], exutils.P100_16GB_X1)
		if err != nil {
			log.Println("Error registering worker:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send OK response
		println("Registered new worker with: ", newWorkerId)
		w.WriteHeader(http.StatusAccepted)
		return
	}
	// Registration secret is wrong, send back bad response
	w.WriteHeader(http.StatusForbidden)

}

// GPU_WORKERS send images+metadata to this endpoint: /api/0/report
// FIXME: this function is too long
func handleUpdateFromWorker(w http.ResponseWriter, r *http.Request) {

	var maxBodySize int64 = 10 * 1024 * 1024

	// Extract the image from the request body
	err := r.ParseMultipartForm(maxBodySize)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	fileFromWebBody, _, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	log.Println("RECEIVING UPDATE")

	var worker exdb.Worker

	// Check if worker is Colab testworker
	// and use the ngrok tunnel if it is

	workerIp := strings.Split(r.RemoteAddr, ":")[0]

	if COLAB_TEST_WORKER {
		// Identify the worker based on ip and get info from the databases
		// https://7150-34-83-189-107.ngrok.io
		worker, err = exdb.GetWorkerByIP(WORKERDB, "https://"+NGROK_IP)
		if err != nil {
			log.Println("Error getting worker from DB", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		log.Println("Handling update for: ", worker.Worker_id, "at", worker.Worker_current_job)
	} else {
		println("+Looking up worker by IP: ", r.RemoteAddr)
		// Identify the worker based on ip and get info from the databases
		// TODO: Identify worker based on workerid instead!
		worker, err = exdb.GetWorkerByIP(WORKERDB, workerIp)
		if err != nil {
			log.Println("Error getting worker from DB", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		println("+Got worker from db: ", worker.Worker_ip, "job", worker.Worker_current_job)
	}

	// 3. Run logic based on the information we received
	// FIXME: This may cause unexpected behavior when trying to override the file
	jobString := strconv.Itoa(worker.Worker_current_job)
	log.Println("Receiving update for JOB:", jobString)
	filepath := exutils.PNG_PATH + jobString + ".png"
	println("SAVING JOB TO", filepath)
	emptyFile, err := os.Create(filepath)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer emptyFile.Close()

	_, err = io.Copy(emptyFile, fileFromWebBody)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Extract metadata out of the request using custom headers defined in exapi/headers.go
	jobIsDone, iterStatus := getIterationStatusAndJobStatusFromHeaders(w, r)

	// Update the jobdb
	// TODO: Add a safety measure where jobs already marked completed can't be updated by the worker
	if jobIsDone {
		println("HAPPENED: CASE A")
		exdb.UpdateJobById(JOBDB, jobString, "completed", iterStatus)
		exdb.UpdateWorkerByJobid(WORKERDB, jobString, true) // set worker to no longer busy
	} else {
		println("HAPPENED: CASE B")
		exdb.UpdateJobById(JOBDB, jobString, "processing", iterStatus)
	}

	w.WriteHeader(http.StatusAccepted)
}

// Helper func returns information extracted from request headers
func getIterationStatusAndJobStatusFromHeaders(w http.ResponseWriter, r *http.Request) (jobIsDone bool, status string) {

	iteration_status_from_header := r.Header.Values(exapi.HeaderJobIterationStatus)
	iteration_status := string(iteration_status_from_header[0])
	println(iteration_status)

	isJobDoneFromHeader := r.Header.Values(exapi.HeaderJobStatusComplete)
	isJobDone := string(isJobDoneFromHeader[0])
	if isJobDone == "1" {
		println("Scheduler detected Job is Done")
		return true, iteration_status
	}

	println("Scheduler detected Job is in progress")
	return false, iteration_status
}

// Keeps the scheduler running until CTRL-C or exit signal is received.
func KeepSchedulerRunningUntilExitSignalIsReceived() {
	fmt.Println("Scheduler is running..") // debug
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-channel
	fmt.Println("scheduler closed gracefully")
}

// Initalizes/creates the log file as needed.
func InitializeLogFile() {
	logFile, err := os.OpenFile(("./logs/scheduler.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	log.SetOutput(logFile)
}

// Run the scheduling loop posting jobs to the workers
func runSchedulingLoop(quit chan bool) {

	for {
		select {
		case <-quit:
			println("Exiting scheduling loop")
			return
		default:

			time.Sleep(5 * time.Second)

			// 1. Get the oldest queued job from the jobdb
			println("checking jobdb for oldest queued job")
			queuedJob, err := exdb.GetOldestQueuedJob(JOBDB)
			if err != nil {
				println("no queued job in db")
				//log.Println("error getting queued job from db")
				continue
			}
			println("Got queued job:", queuedJob.Jobid, queuedJob.Prompt)

			// 2. Get a worker from the workerdb that is not busy
			worker, err := exdb.GetFreeWorker(WORKERDB)
			if err != nil {
				println("error getting worker from db")
				log.Println("error getting worker from db")
				continue

			}
			println("Got worker: ", worker.Worker_id)

			// 3. Assign the job to the worker

			//recoverWrapperForPostJobToWorker(queuedJob, worker)

			err = postJobToWorker(queuedJob, worker)
			if err != nil {
				println("Error when trying to post job to worker", err)
				log.Println("Error when trying to post job to worker", err)
				continue
			}
			if PANIC {
				println("Error, panic when trying to post job to worker", err)
				log.Println("Error, panic when trying to post job to worker", err)
				PANIC = false
				continue
			}

			// 4. Update the Jobdb and the Workerdb
			exdb.UpdateJobById(JOBDB, queuedJob.Jobid, "processing", "1")
			exdb.UpdateWorkerByWorkerId(WORKERDB, queuedJob.Jobid, worker.Worker_id, 1)
		}
	}
}

// Sends job to worker and returns error if it fails
func postJobToWorker(job exdb.Job, worker exdb.Worker) error {

	// Intercept the panic that can happen when the request fails
	defer func() {
		if err := recover(); err != nil {
			println("Panic intercepted!", err)
			PANIC = true
		}
	}()

	workerUrl := "http://" + worker.Worker_ip + ":8090/api/0/worker"
	if COLAB_TEST_WORKER { // Ugly special case for colab testing
		workerUrl = "https://" + worker.Worker_ip + "/api/0/worker"

	}
	println("Posting job:\"", job.Prompt, "\"to worker: ", worker.Worker_id)

	client := &http.Client{}
	client.Timeout = time.Second * 10
	client.Transport = http.DefaultTransport

	// Marshal the job to json
	jsonData, err := json.Marshal(job)
	if err != nil {
		log.Println("error marshalling struct into json", err)
		return err
	}

	request, err := http.NewRequest("POST", workerUrl, bytes.NewBuffer(jsonData))
	request.Header.Set("Content-Type", "application/json; charset=UTF-8")
	if err != nil {
		log.Println("error creating request to send to worker")
		return err
	}

	// TODO: Scheduler is waiting on a response that doesn't come

	response, err := client.Do(request)
	if err != nil {
		log.Println("Error posting job to scheduler", err)
		return err
	}
	defer response.Body.Close()

	// Only return everything okay if correct status code

	if response.StatusCode == http.StatusAccepted {
		log.Println("Job posting accepted for JobId:", job.Jobid)
		return nil
	}

	if _, exists := response.Request.Header[exapi.HeaderJobAccepted]; exists {
		return nil

	} else {

		return errors.New("Error posting job")
	}

	// FIXME
	//r.Header.Values(http.StatusOK

	/*

		defer response.Body.Close()

		fmt.Println("response Status:", response.Status)
		fmt.Println("response Headers:", response.Header)
		body, _ := ioutil.ReadAll(response.Body)
		fmt.Println("response Body:", string(body))

		return nil
	*/
}

// This is the main function >:D
func main() {
	SECRET = exutils.InitializeSecretFromArgument()
	InitializeLogFile()
	JOBDB = exdb.InitializeJobdb()
	WORKERDB = exdb.InitializeWorkerdb()

	// Register handlers
	http.HandleFunc("/api/0/registration", handleWorkerRegistration)
	http.HandleFunc("/api/0/report", handleUpdateFromWorker)

	go http.ListenAndServe(WEBSERVER_PORT, nil)

	quit := make(chan bool)
	go runSchedulingLoop(quit)

	// Just testing
	//newWorkerId := uuid.New().String()
	//exdb.RegisterNewWorker(WORKERDB, newWorkerId, "1.1.1.1:69", exutils.P100_16GB_X1)

	KeepSchedulerRunningUntilExitSignalIsReceived()
	close(quit)
}
