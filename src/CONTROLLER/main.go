package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"extropic-art-backend/src/exapi"
	"extropic-art-backend/src/exdb"
)

const CONTROLLER_PORT = ":8080"

var JOBDB *sql.DB

// Answers calls to the endpoint /api/0/jobs
func api_0_jobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	switch r.Method {
	case "GET":
		exapi.HandleJobsApiGet(JOBDB, w, r)
	case "POST": // Handle new job postings
		// exapi.HandleJobsApiPost(JOBDB, w, r) // deprecated!
	}
}

// Answers calls to the endpoint /api/0/img
func api_0_img(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	exapi.HandleImgRequests(w, r)
}

// Answers calls to the endpoint /api/0/status
func api_0_status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	exapi.HandleStatusRequest(JOBDB, w, r)
}

// Answers calls to the endpoint /api/1/queue
func api_1_queue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")  // TESTING: allow CORS for testing purposes
	w.Header().Set("Access-Control-Allow-Headers", "*") //FIXME because I don't get it
	exapi.HandleQueueRequest(JOBDB, w, r)
}

// Answers calls to the endpoint /api/1/status
func api_1_status_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	exapi.Handle_status_api_endpoint_version_2(JOBDB, w, r)
}

// Answers calls to the endpoint /api/1/jobs
func api_1_jobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API 1 called!")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	switch r.Method {
	case "GET":
		exapi.HandleJobsApiGet(JOBDB, w, r)
	case "POST": // Handle new job postings
		exapi.HandleJobsApiPostVersion2(JOBDB, w, r)
	}
}

// Initializes log file for the controller
func initializeLogFile() {
	logFile, err := os.OpenFile(("./logs/controller.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	log.SetOutput(logFile)
}

// This is the main function :D
func main() {
	initializeLogFile()
	JOBDB = exdb.InitializeJobdb()

	http.Handle("/", http.FileServer(http.Dir("../view/build"))) // Handles requests for ressources in dist

	// Register API endpoints
	http.HandleFunc("/api/0/status", api_0_status)
	http.HandleFunc("/api/0/jobs", api_0_jobs)
	http.HandleFunc("/api/0/img", api_0_img)

	// New API endpoints used by the react frontend
	http.HandleFunc("/api/1/queue", api_1_queue)
	http.HandleFunc("/api/1/status", api_1_status_handler)
	http.HandleFunc("/api/1/jobs", api_1_jobs)

	http.ListenAndServe(CONTROLLER_PORT, nil)
}
