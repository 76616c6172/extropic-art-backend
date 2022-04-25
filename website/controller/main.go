package main

import (
	"dd-web-app/controller/internals"
	"encoding/json"

	"net/http"
	"text/template"
)

const WEBSERVER_PORT = ":80" // this is where we serve for testing purposes

var view *template.Template

func init() {
	view = template.Must(template.ParseGlob("view/*.gohtml"))
}

// Handler for all requests to www.the_url/
func index_handler(w http.ResponseWriter, r *http.Request) {
	view.ExecuteTemplate(w, "index.gohtml", nil)
}

/* // Maybe we don't need a gpu endpoint
// Answers calls to the endpoint /api/0/gpu/0
// Returns information about the state of the GPUs
func api_0_gpu(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, internals.Get_gpu_status())
	fmt.Println(r.Method) //debug
}
*/

// Answers calls to the endpoint /api/0/jobs
// POST and GET
func api_0_jobs(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET": // Return info about existing jobs
		internals.ReturnInfoAboutAllJobs(w)

	case "POST": // take in new jobs
		internals.HandleNewJobRequest(w, r)
	}
}

// Answers calls to the endpoint /api/0/img
// GET
// TODO: requires a jobid and sends back the latest image for that job id.
func api_0_img(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		//jobid := "752b5cd9013dcf3f6ebf577f99fa76adf4f32459"
		internals.HandleImgRequests(w, r) // TODO: Return the correct image based on the request (request with jobid)
	}

}

// Answers calls to the endpoint /api/0/all
// This answers with a json containing all information the view needs when it first loads
func api_0_status(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // send back the response

		type status struct {
			Gpu0           string   `json:"gpu0"`
			Current_job    string   `json:"current_job"`
			Completed_jobs []string `json:"completed_jobs"`
			//Description string `json:"Description"`
		}

		var all_response = status{
			Gpu0:           "ready", // busy, offline
			Completed_jobs: []string{"752b5cd9013dcf3f6ebf577f99fa76adf4f32459"},
		}

		json.NewEncoder(w).Encode(all_response) // send back the json as a the response
	}
}

// Starts the webserver on port WEBSERVER_PORT (8080 for testing)
func main() {
	http.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./model/pub")))) //serves requests to www.url/assets/

	http.HandleFunc("/", index_handler) // handles requests to url

	http.HandleFunc("/api/0/status", api_0_status) // registers api_0_status() as the handler for "/api/0/status"
	http.HandleFunc("/api/0/jobs", api_0_jobs)
	http.HandleFunc("/api/0/img", api_0_img)
	//http.HandleFunc("/api/0/gpu", api_0_gpu)     // maybe we don't need it!

	http.ListenAndServe(WEBSERVER_PORT, nil) // Run the server
}

/* According to Tod MclEod, this is how to handle requests
type Handler interface {
        ServeHTTP(ResponseWriter, *Request)
}
*/
