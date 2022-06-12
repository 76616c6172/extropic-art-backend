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
func handleWorkerReportUpdateToJob(w http.ResponseWriter, r *http.Request) {
	println("Receiving update from worker:")
	println()
	println("1:", r.Method) // "POST"
	var maxBodySize int64 = 10 * 1024 * 1024

	workerIp := r.Host
	println(workerIp)

	iteration_status_from_header := r.Header.Values("Iteration-Status")
	iteration_status := string(iteration_status_from_header[0])
	println(iteration_status)

	err := r.ParseMultipartForm(maxBodySize)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	//Access the photo key - First Approach
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

	doSomeLogicAndUpdateWorkerDb(w, r)

}

// Does what it says
func doSomeLogicAndUpdateWorkerDb(w http.ResponseWriter, r *http.Request) {
	// r.RemoteAddr <- this is the ip:port of the worker

	//exdb.GetWorkerByIP()

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
	http.HandleFunc("/api/0/report", handleWorkerReportUpdateToJob)

	go http.ListenAndServe(WEBSERVER_PORT, nil)
	KeepSchedulerRunningUntilExitSignalIsReceived()
}
