package main

import (
	"encoding/json"
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

var SECRET string

// Sends back empty response with status code 500
func sendErrorResponseToWorker(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("500 - Error registering worker"))
	e := struct {
		Registration string `json:"registration"`
	}{}
	json.NewEncoder(w).Encode(e) // send back the json as a the response
}

// GPU_WORERS register themselves with the scheduler through this endpoint
// /api/0/registration
func api_0_registration(w http.ResponseWriter, r *http.Request) {

	registrationSecret, err := r.Cookie("secret")
	if err != nil {
		log.Println("Error reading secret cookie: ", err)
		sendErrorResponseToWorker(w, r)
		return
	}

	if registrationSecret.Value == SECRET {
		newWorkerId := uuid.New().String()
		fmt.Println("New worker ID created:", newWorkerId)

		err := exdb.RegisterNewWorker(newWorkerId, IP, P100)
		if err != nil {
			log.Println("Error registering worker:", err)
			sendErrorResponseToWorker(w, r)
			return
		}

		fmt.Println("Registration successful") // debug
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

func main() {

	// Check arguments
	if len(os.Args) < 2 || len(os.Args) > 2 {
		fmt.Println("Error: You must supply EXACTLY one argument (the GPU_WORER auth token) on startup.")
		os.Exit(1)
	}
	SECRET = strings.TrimSpace(os.Args[1])

	InitializeLogFile()
	exdb.InitializeWorkerdb()
	exdb.InitializeJobdb()

	http.HandleFunc("/api/0/registration", api_0_registration)

	go http.ListenAndServe(WEBSERVER_PORT, nil)

	KeepSchedulerRunningUntilExitSignalIsReceived()
}
