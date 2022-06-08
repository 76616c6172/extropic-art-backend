package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	"project-exia-monorepo/website/exapi"
	"project-exia-monorepo/website/exdb"
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
	case "POST":
		exapi.HandleJobsApiPost(JOBDB, w, r)
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

// Initializes log file for the controller
func initializeLogFile() {
	logFile, err := os.OpenFile(("./logs/exia.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	log.SetOutput(logFile)
}

// This is the main function :D
func main() {
	initializeLogFile()
	JOBDB = exdb.InitializeJobdb()

	// Handle requests for ressources in dist
	http.Handle("/", http.FileServer(http.Dir("../view/dist")))

	// Register API endpoints
	http.HandleFunc("/api/0/status", api_0_status)
	http.HandleFunc("/api/0/jobs", api_0_jobs)
	http.HandleFunc("/api/0/img", api_0_img)

	http.ListenAndServe(CONTROLLER_PORT, nil)
}
