package main

import (
	"extropic-art-backend/src/exapi"
	"fmt"
	"net/http"
)

// Answers calls to the endpoint /api/0/jobs
func api_0_jobs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	switch r.Method {
	case "GET":
		exapi.HandleJobsApiGet(JOB_DB, w, r)
	case "POST": // Handle new job postings
		// exapi.HandleJobsApiPost(JOB_DB, w, r) // deprecated!
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
	exapi.HandleStatusRequest(JOB_DB, w, r)
}

// Answers calls to the endpoint /api/1/queue
func api_1_queue(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")  // TESTING: allow CORS for testing purposes
	w.Header().Set("Access-Control-Allow-Headers", "*") //FIXME because I don't get it
	exapi.HandleQueueRequest(JOB_DB, w, r)
}

// Answers calls to the endpoint /api/1/status
func api_1_status_handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")
	exapi.Handle_status_api_endpoint_version_2(JOB_DB, w, r)
}

// Answers calls to the endpoint /api/1/jobs
func api_1_jobs(w http.ResponseWriter, r *http.Request) {
	fmt.Println("API 1 called!")

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "*")

	switch r.Method {
	case "GET":
		exapi.HandleJobsApiGet(JOB_DB, w, r)
	case "POST": // Handle new job postings
		exapi.HandleJobsApiPostVersion2(JOB_DB, w, r)
	}
}
