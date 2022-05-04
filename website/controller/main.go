package main

import (
	"exia/controller/internals"

	"net/http"
	"text/template"
)

const WEBSERVER_PORT = ":8080"

var view *template.Template

// Know which templates to use
func init() {
	view = template.Must(template.ParseGlob("view/*.gohtml"))
}

// Handler for all requests to www.the_url/
func index_handler(w http.ResponseWriter, r *http.Request) {
	view.ExecuteTemplate(w, "index.gohtml", nil)
}

// Answers calls to the endpoint /api/0/jobs
func api_0_jobs(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "POST": // Return info about existing jobs (I would like to make this GET but the frontend requires POST)
		internals.HandleJobsApiPost(w, r)
	case "PUT": // take in new jobs
		// TODO: deal with accepting new jobs
	}
}

// Answers calls to the endpoint /api/0/img
// TODO: requires a jobid and sends back the latest image for that job id.
func api_0_img(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		internals.HandleImgRequests(w, r) // TODO: Return the correct image based on the request (request with jobid)
	}

}

// Answers calls to the endpoint /api/0/all
// This answers with a json containing all information the view needs when it first loads
func api_0_status(w http.ResponseWriter, r *http.Request) {
	internals.HandleStatusRequest(w, r)
}

// This is the main function :D
func main() {

	// Handle requests for web assets
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./model/pub")))) //serves requests to www.url/assets/
	http.HandleFunc("/", index_handler)                                                            // handles requests to /

	// Handle API endpoints
	http.HandleFunc("/api/0/status", api_0_status) // registers api_0_status() as the handler for "/api/0/status"
	http.HandleFunc("/api/0/jobs", api_0_jobs)
	http.HandleFunc("/api/0/img", api_0_img)

	// Run the webserver
	http.ListenAndServe(WEBSERVER_PORT, nil)
}
