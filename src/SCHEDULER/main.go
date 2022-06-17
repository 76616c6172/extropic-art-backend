package main

import (
	"database/sql"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"project-exia-monorepo/website/exapi"
	"project-exia-monorepo/website/exdb"
	"project-exia-monorepo/website/exutils"

	"github.com/google/uuid"
)

const WEBSERVER_PORT = ":8091" // Scheduler is listening on this port
const JOB_COMPLETE = true

var WORKERDB *sql.DB //pointer used to connect to the db, initialized in main
var JOBDB *sql.DB
var SECRET string

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

		err := exdb.RegisterNewWorker(WORKERDB, newWorkerId, r.RemoteAddr, exutils.P100_16GB_X1)
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

	//w.WriteHeader(http.StatusAccepted) // THIS SENDS A RESPONSE RIGHt?
	//return

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

	// Identify the worker based on ip and get info from the databases
	worker, err := exdb.GetWorkerByIP(WORKERDB, r.RemoteAddr)
	if err != nil {
		log.Println("Error getting worker from DB", err)
		w.WriteHeader(http.StatusBadRequest)
		return

	}
	println(worker.Worker_current_job)
	job, err := exdb.GetJobByJobid(JOBDB, worker.Worker_current_job)
	if err != nil {
		log.Println("Error getting job from DB", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// 3. Run logic based on the information we received
	// FIXME: This may cause unexpected behavior when trying to override the file
	filepath := exutils.PNG_PATH + job.Jobid + ".png"
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
	if jobIsDone {
		exdb.UpdateJobById(JOBDB, job.Jobid, "completed", iterStatus)
		exdb.UpdateWorkerByJobid(WORKERDB, job.Jobid, JOB_COMPLETE)
	} else {
		exdb.UpdateJobById(JOBDB, job.Jobid, "processing", iterStatus)

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
		return true, iteration_status
	}
	return true, iteration_status
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

	
	// Just testing
	//newWorkerId := uuid.New().String()
	//exdb.RegisterNewWorker(WORKERDB, newWorkerId, "1.1.1.1:69", exutils.P100_16GB_X1)

	KeepSchedulerRunningUntilExitSignalIsReceived()
}
