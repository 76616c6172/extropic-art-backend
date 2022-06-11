package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"project-exia-monorepo/website/exdb"

	"github.com/google/uuid"
)

const WEBSERVER_PORT = ":8091" // Scheduler is listening on this port
const P100 = 1                 // Default worker type
const IP = "127.0.0.1"         // Local worker ip for testing

var WORKERDB *sql.DB //pointer used to connect to the db, initialized in main
var SECRET string

// Initialize and connect to workerdb
// GPU_WORERS register themselves with the scheduler through this endpoint
// /api/0/registration
func workerRegistrationHandler(w http.ResponseWriter, r *http.Request) {

	registrationSecret, err := r.Cookie("secret")
	if err != nil {
		log.Println("Error reading secret cookie: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if registrationSecret.Value == SECRET {
		newWorkerId := uuid.New().String()
		fmt.Println("New worker ID created:", newWorkerId)

		err := exdb.RegisterNewWorker(WORKERDB, newWorkerId, IP, P100)
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
func jobReportHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {

	}

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

func initializeSecretFromArgument() {
	if len(os.Args) < 2 || len(os.Args) > 2 { // Check arguments
		fmt.Println("Error: You must supply EXACTLY one argument (the GPU_WORER auth token) on startup.")
		os.Exit(1)
	}
	SECRET = strings.TrimSpace(os.Args[1])
}

// This is the main function >:D
func main() {
	initializeSecretFromArgument()
	InitializeLogFile()
	exdb.InitializeJobdb()
	WORKERDB = exdb.InitializeWorkerdb()

	// Register handlers
	http.HandleFunc("/api/0/registration", workerRegistrationHandler)
	http.HandleFunc("/api/0/report", jobReportHandler)

	go http.ListenAndServe(WEBSERVER_PORT, nil)
	KeepSchedulerRunningUntilExitSignalIsReceived()
}
