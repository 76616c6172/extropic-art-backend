package main

import (
	"log"
	"net/http"
	"os"

	"project-exia-monorepo/website/exapi"
	"project-exia-monorepo/website/exdb"
)

const CONTROLLER_PORT = ":8080"

// Answers calls to the endpoint /api/0/jobs
func api_0_jobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	switch r.Method {
	case "GET":
		exapi.HandleJobsApiGet(w, r)
	case "POST":
		exapi.HandleJobsApiPost(w, r)
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
	exapi.HandleStatusRequest(w, r)
}

// Initializes log file for the controller
func initializeLogFile() {

	logFile, err := os.OpenFile(("./logs/exia.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

}

// This is the main function :D
func main() {

	initializeLogFile()
	exdb.InitializeJobdb()

	// Handle requests for everything in ../view/dist - dist is accessible to the public without AUTH
	http.Handle("/", http.FileServer(http.Dir("../view/dist")))

	// Register API endpoints
	http.HandleFunc("/api/0/status", api_0_status)
	http.HandleFunc("/api/0/jobs", api_0_jobs)
	http.HandleFunc("/api/0/img", api_0_img)

	http.ListenAndServe(CONTROLLER_PORT, nil)
}
