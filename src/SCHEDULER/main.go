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

var WORKERDB *sql.DB //pointer used to connect to the db, initialized in main
var SECRET string

// Initialize and connect to workerdb
// GPU_WORERS register themselves with the scheduler through this endpoint
// /api/0/registration
func handleWorkerRegistration(w http.ResponseWriter, r *http.Request) {
	error := WORKERDB.Ping()
	log.Println("Pinging workerdb error: ", error)

	log.Println("registering worker with ip: ", r.RemoteAddr)

	registrationSecret, err := r.Cookie("secret")
	if err != nil {
		log.Println("Error reading secret cookie: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if registrationSecret.Value == SECRET {
		newWorkerId := uuid.New().String()
		log.Println("New worker ID created:", newWorkerId)

		err := exdb.RegisterNewWorker(WORKERDB, newWorkerId, r.RemoteAddr, exutils.P100_16GB_X1)
		if err != nil {
			log.Println("Error registering worker:", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Send OK response
		w.WriteHeader(http.StatusAccepted)

		fmt.Println("Registration successful") // debug
		return
	}
	// Registration secret is wrong, send back bad response
	w.WriteHeader(http.StatusForbidden)
}

// GPU_WORKERS send images+metadata to this endpoint
// /api/0/report
func handleUpdateFromWorker(w http.ResponseWriter, r *http.Request) {
	println("Receiving update from worker:")
	println()
	println("1:", r.Method) // "POST"
	var maxBodySize int64 = 10 * 1024 * 1024

	iteration_status_from_header := r.Header.Values(exapi.HeaderJobIterationStatus)
	iteration_status := string(iteration_status_from_header[0])
	println(iteration_status)

	isJobDoneFromHeader := r.Header.Values(exapi.HeaderJobStatusComplete)
	isJobDone := string(isJobDoneFromHeader[0])
	println(isJobDone)

	err := r.ParseMultipartForm(maxBodySize)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	file, _, err := r.FormFile("image")
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tmpfile, err := os.Create("./" + "test.png")
	defer tmpfile.Close()
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(200)
	doSomeLogicAndUpdateModel(w, r)
}

// Does what it says
// Writes updated image
// Update jobdb
// Workerdb if Done
func doSomeLogicAndUpdateModel(w http.ResponseWriter, r *http.Request) {
	// r.RemoteAddr <- this is the ip:port of the worker

	// 1. Update jobdb and save the image correctly (+ make jpeg)
	workerFromDb, err := exdb.GetWorkerByIP(WORKERDB, r.RemoteAddr)
	if err != nil {
		log.Println("Error getting worker from DB", err)
		return
	}
	log.Println(workerFromDb.Worker_ip)
	log.Println(workerFromDb.Worker_current_job)
	// worker := readblabla..

	// exdb.UpdateJobStatus(worker.Iteration_status)

	// 2. Check if the worker is done
	//if worker_iteration.status ==

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
	exdb.InitializeJobdb()
	WORKERDB = exdb.InitializeWorkerdb()

	// Register handlers
	http.HandleFunc("/api/0/registration", handleWorkerRegistration)
	http.HandleFunc("/api/0/report", handleUpdateFromWorker)

	go http.ListenAndServe(WEBSERVER_PORT, nil)
	KeepSchedulerRunningUntilExitSignalIsReceived()
}
