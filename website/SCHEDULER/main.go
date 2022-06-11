package main

import (
	"database/sql"
	"fmt"
	"io"
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
func handleWorkerRegistration(w http.ResponseWriter, r *http.Request) {

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
func handleWorkerReportUpdateToJob(w http.ResponseWriter, r *http.Request) {
	println("Receiving update from worker:")
	println()
	println("1:", r.Method) // "POST"
	var maxBodySize int64 = 10 * 1024 * 1024

	err := r.ParseMultipartForm(maxBodySize)
	if err != nil {
		println("A: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	//Access the photo key - First Approach
	file, _, err := r.FormFile("image")
	if err != nil {
		println("B: ", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	tmpfile, err := os.Create("./" + "test.png")
	defer tmpfile.Close()
	if err != nil {
		println("C: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	_, err = io.Copy(tmpfile, file)
	if err != nil {
		println("D: ", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(200)
	return

	/* NO
	filename := "test.png"
	buf := new(bytes.Buffer)
	writer := zip.NewWriter(buf)
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	f, err := writer.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	_, err = f.Write([]byte(data))
	if err != nil {
		log.Fatal(err)
	}
	err = writer.Close()
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile("test.png", buf.Bytes(), 0666)
	println(err)
	//os.O_RDWR|os.O_CREATE|os.O_APPEND,

	w.Header().Set("Content-Type", "application/zip")
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s.zip\"", filename))
	w.Write(buf.Bytes())
	*/

	/*
		println()
		println("2:", r.ContentLength) // 2360284
		println()
		println("3:", r.Header.Values("name"))
		println()
		println("4:", r.Form.Get("name"))
		println()
		println("5:", r.Body)
		println()
		println("6:", r.URL)
		println()
	*/

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
	http.HandleFunc("/api/0/registration", handleWorkerRegistration)
	http.HandleFunc("/api/0/report", handleWorkerReportUpdateToJob)

	go http.ListenAndServe(WEBSERVER_PORT, nil)
	KeepSchedulerRunningUntilExitSignalIsReceived()
}
