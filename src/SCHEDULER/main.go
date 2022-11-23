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
	"os/exec"
	"os/signal"
	"strconv"
	"strings"
	"syscall"
	"time"

	"extropic-art-backend/src/exapi"
	"extropic-art-backend/src/exdb"
	"extropic-art-backend/src/exutils"

	"github.com/google/uuid"
)

const PORT_TO_LISTEN_ON = ":8091"

var WORKER_DB *sql.DB
var JOB_COMPLETE = true
var JOBDB *sql.DB
var SECRET string
var PANIC bool
var WORKER_IP_TO_TUNNEL_URL = map[string]string{} // FIXME: scheduler should be stateless

// handles /api/0/registration endpoint
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

		workerIpForDB := strings.Split(r.RemoteAddr, ":")
		err := exdb.RegisterNewWorker(WORKER_DB, newWorkerId, workerIpForDB[0], exutils.P100_16GB_X1)
		if err != nil {
			log.Println("Error registering worker:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Check if the worker is using a tunnel (TODO: this should be saved in workerdb)
		workerTunnel, err := r.Cookie(exapi.CookieWorkerTunnel)
		if err != nil {
			println("No worker tunnel cookie", err)
		}
		if len(workerTunnel.Value) > 0 {
			println("Received worker tunnel:", workerTunnel.Value)
			WORKER_IP_TO_TUNNEL_URL[workerIpForDB[0]] = workerTunnel.Value
		}

		println("Registered new worker with: ", newWorkerId)
		w.WriteHeader(http.StatusAccepted)
		return
	}

	// If registration secret is wrong, send back bad response
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

	println("+Looking up worker by IP: ", workerIp)
	// Identify the worker based on ip and get info from the databases
	// TODO: Identify worker based on workerid instead!
	worker, err = exdb.GetWorkerByIP(WORKER_DB, workerIp)
	if err != nil {
		log.Println("Error getting worker from DB", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	println("+Got worker from db: ", worker.Worker_ip, "job", worker.Worker_current_job)

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

	// Create JPG "thumbnail"
	exec.Command("../model/make_jpgs_from_name", jobString).Run()

	// Extract metadata out of the request using custom headers defined in exapi/headers.go
	jobIsDone, iterStatus := getIterationStatusAndJobStatusFromHeaders(w, r)

	// Update the jobdb
	// TODO: Add a safety measure where jobs already marked completed can't be updated by the worker
	if jobIsDone {
		exdb.UpdateJobById(JOBDB, jobString, "completed", iterStatus)
		exdb.UpdateWorkerByJobid(WORKER_DB, jobString, true) // set worker to no longer busy
	} else {
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
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
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

	var PROMPT, SEED, WIDTH, HEIGHT, STEPS, SCALE string

	for {
		select {

		case <-quit:
			println("scheduler exiting")
			return

		default:

			time.Sleep(5 * time.Second)

			// 1. Get the oldest queued job from the jobdb
			println("checking jobdb for oldest queued job")
			queuedJob, err := exdb.GetOldestQueuedJob(JOBDB)
			if err != nil {
				println("..waiting for queued job")
				//log.Println("error getting queued job from db")
				continue
			}
			println("Got queued job:", queuedJob.Jobid, queuedJob.Prompt)

			// 0. TODO: select the pipelines
			modelCmd := "./run_worker.py"

			// 2. Run serverless GPU worker

			PROMPT = queuedJob.Prompt

			// This could be much cleaner
			WIDTH = "512"
			HEIGHT = "512"
			if queuedJob.Job_params == "2" {
				WIDTH = "512"
				HEIGHT = "768"
			} else if queuedJob.Job_params == "3" {
				WIDTH = "768"
				HEIGHT = "512"
			}

			SEED = strconv.Itoa(queuedJob.Seed)
			STEPS = "75"
			SCALE = "7"
			if queuedJob.Guidance == 1 {
				STEPS = "145"
				SCALE = "6"
			}

			// Ugly hack to communicate the pre-prompt option without changing the jobdb schema..
			// fix this.
			if strings.Contains(queuedJob.Owner, "with_pre_prompt") {
				PROMPT = "mdjrny v4 style " + queuedJob.Prompt
			} else {
				PROMPT = queuedJob.Prompt
			}

			exdb.UpdateJobById(JOBDB, queuedJob.Jobid, "processing", "1")

			cmd := exec.Command(modelCmd, PROMPT, SEED, WIDTH, HEIGHT, STEPS, SCALE, queuedJob.Jobid)
			out, err := cmd.Output()
			if err != nil {
				log.Println(err)
			}
			fmt.Println(string(out))

			// Create JPG "thumbnail", TODO: Pull this into the gi runtime and just use the standard library for this
			exec.Command("../model/make_jpgs_from_name", queuedJob.Jobid).Run()

			// Update the jobdb, TODO: Add safety measure where jobs already marked completed can't be updated by the worker
			exdb.UpdateJobById(JOBDB, queuedJob.Jobid, "completed", "250")
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

	// Default case if worker is reachable by ip
	workerUrl := "http://" + worker.Worker_ip + ":8090/api/0/worker"

	// Check if this worker is using a tunnel
	if url, exists := WORKER_IP_TO_TUNNEL_URL[worker.Worker_ip]; exists {
		workerUrl = "https://" + url + "/api/0/worker"
	}

	log.Println("Posting job:\"", job.Prompt, "\"to worker: ", worker.Worker_id)

	client := &http.Client{}
	client.Timeout = time.Second * 10
	client.Transport = http.DefaultTransport

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

	response, _ := client.Do(request)
	if err != nil {
		log.Println("Error posting job to scheduler", err)
		return err
	}
	defer response.Body.Close()

	// Only return everything okay if correct status code
	if _, exists := response.Header[exapi.HeaderJobAccepted]; !exists {
		println("job posting rejected by worker")
		return errors.New("job posting rejected by worker")

	}

	return nil
}

// This is the main function >:D
func main() {
	SECRET = exutils.InitializeSecretFromArgument()
	InitializeLogFile()
	JOBDB = exdb.InitializeJobdb()
	WORKER_DB = exdb.InitializeWorkerdb()

	// Register handlers
	http.HandleFunc("/api/0/registration", handleWorkerRegistration)
	http.HandleFunc("/api/0/report", handleUpdateFromWorker)

	go http.ListenAndServe(PORT_TO_LISTEN_ON, nil)

	quit := make(chan bool)
	go runSchedulingLoop(quit)

	KeepSchedulerRunningUntilExitSignalIsReceived()
	close(quit)
}
