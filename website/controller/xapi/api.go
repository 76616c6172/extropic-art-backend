package xapi

import (
	"encoding/json"
	"exia/controller/xdb"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

const GPU_STATUS string = "offline" // can be offline, online, or busy

type newjob struct {
	Prompt string `json:"prompt"`
}

// This is the response object of he /api/0/jobs endpoint
// For reference here is the Schema for request sent to job endpoint by client
/*
{
  "jobid": "1",
  "prompt": "Space wool bla bla, bla bla..",
  "job_status": "qeued",
  "iteration_status": "125",
  "iteration_max": "240",
}
*/
type apiJob struct {
	Jobid            string `json:"jobid"`
	Prompt           string `json:"prompt"`
	Job_status       string `json:"job_status"`
	Iteration_status int    `json:"iteration_status"`
	Iteration_max    int    `json:"iteration_max"`
}

// Schema for the status object returned by the status endpoint
type status struct {
	Gpu            string   `json:"gpu"`
	Completed_jobs []apiJob `json:"completed_jobs"`
	//Description string `json:"Description"`
}

// Deals with requests to the status endpoint
func HandleStatusRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // TESTING: allow CORS for testing purposes

	if r.Method == "GET" { // send back the response

		var responseObject = status{ // We append to this object as we go and send it back to the client
			Gpu: "ready", // busy, offline
		}

		// TODO: dumping entire database into memory every time won't scale so get only the last 10!
		allJobs, err := xdb.GetAllJobs()
		if err != nil {
			log.Println(err)
			//TODO: et response code to 500 to indicate server error
			json.NewEncoder(w).Encode(responseObject) // send back the json as a the response
		}

		var respJob apiJob
		for i := 0; i < 10; i++ {
			if i < len(allJobs) { // sanity check if we have a fresh database with less than 10 entries
				respJob.Jobid = allJobs[i].Jobid
				respJob.Prompt = allJobs[i].Prompt
				respJob.Job_status = allJobs[i].Status
				respJob.Iteration_max = allJobs[i].Iteration_max
				respJob.Iteration_status = allJobs[i].Iteration_status
				responseObject.Completed_jobs = append(responseObject.Completed_jobs, respJob)
			}
		}

		json.NewEncoder(w).Encode(responseObject) // send back the json as a the response
	}

}

// Obtains the coorect image for a given job and sends it back
func HandleImgRequests(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Access-Control-Allow-Origin", "*") // TESTING: allow CORS for testing purposes

	input := fmt.Sprintln(r.URL)
	inputstring := strings.TrimLeft(input, "/api/0/img?jobid=")
	inputstring2 := strings.TrimSpace(inputstring)

	/* Don/t care bout the body for now
	jsonDecoder := json.NewDecoder(r.Body)
	var imgRequest imgRequest
	err := jsonDecoder.Decode(&imgRequest)
	if err != nil {
		log.Println(err) // maybe handle this better
		return
	}
	*/

	//img, err := os.Open("./model/images/" + imgRequest.Jobid + ".png") // for now just get this image for testing
	img, err := os.Open("../model/images/" + inputstring2 + ".png") // temporary
	// TODO: Actually lookfor the image in SQLite database
	// img, err := os.Open("./model/images/" + imgRequest.Jobid + ".png") // for now just get this image for testing
	if err != nil {
		fmt.Println(err) // perhaps handle this nicer
		return
	}
	defer img.Close()

	w.Header().Set("Content-Type", "image/png") // <-- set the content-type header
	_, err = io.Copy(w, img)                    // send the image
	if err != nil {
		log.Println(err)
	}

}

// takes /api/0/jobs=?jobid="yourjodidhere"
// sends back json object with info about the job
func HandleJobsApiGet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // TESTING: allow CORS for testing purposes

	// Get the jobid
	input := fmt.Sprintln(r.URL)
	inputstring := strings.TrimLeft(input, "/api/0/jobs?jobid=")
	inputstring2 := strings.TrimSpace(inputstring)

	sanitized_input, err := strconv.Atoi(inputstring2)
	if err != nil {
		log.Println(err)
		return
	}

	var jobResponse apiJob
	var j xdb.Job
	j, err = xdb.GetJobByJobid(sanitized_input)
	if err != nil {
		log.Println(err)
		return
	}

	// Build the response
	jobResponse.Jobid = j.Jobid
	jobResponse.Prompt = j.Prompt
	jobResponse.Iteration_status = j.Iteration_status
	jobResponse.Iteration_max = j.Iteration_max

	json.NewEncoder(w).Encode(j) // send back the json as a the response
}

// Deals with POST requests made to the jobs endpoint (POST new jobs)
func HandleJobsApiPost(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "https://www.exia.art") // TESTING: allow CORS for testing purposes
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	//w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))

	jsonDecoder := json.NewDecoder(r.Body)
	var jobRequest newjob
	err := jsonDecoder.Decode(&jobRequest)
	if err != nil {
		log.Println(err) // maybe handle this better
		return
	}

	type jobResponse struct {
		Jobid      string `json:"jobid"`
		Prompt     string `json:"prompt"`
		Job_status string `json:"job_status"`
	}

	var j jobResponse
	j.Jobid = ""                 // Placeholder
	j.Prompt = jobRequest.Prompt // TODO: Validate and sanitize user input first
	j.Job_status = "accepted"    //pending? rejected?

	json.NewEncoder(w).Encode(j) // send back the json as a the response
}
