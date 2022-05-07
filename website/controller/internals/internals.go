package internals

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

const GPU_STATUS string = "offline" // can be offline, online, or busy

type newjob struct {
	Prompt string `json:"prompt"`
}

// Schema for request sent to job endpoint by client
type job struct {
	Jobid  string `json:"jobid"`
	Prompt string `json:"prompt"`
}

// Schema for request sent to img endpoint by client
type imgRequest struct {
	Jobid string `json:"jobid"`
}

// Schema for the status object returned by the status endpoint
type status struct {
	Gpu            string `json:"gpu"`
	Current_job    string `json:"current_job"`
	Completed_jobs []job  `json:"completed_jobs"`
	//Description string `json:"Description"`
}

// Deals with requests to the status endpoint
func HandleStatusRequest(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // TESTING: allow CORS for testing purposes

	if r.Method == "GET" { // send back the response

		var testJob job
		testJob.Jobid = "752b5cd9013dcf3f6ebf577f99fa76adf4f32459"
		testJob.Prompt = "3d render of celestial space nebula, cosmic, space station, unreal engine 3, photorealistic materials, trending on Artstation"

		var responseObject = status{
			Gpu: "ready", // busy, offline
		}

		// Just testing
		responseObject.Completed_jobs = append(responseObject.Completed_jobs, testJob)
		testJob.Jobid = "47d5090d689508acf1a6c29695e0d05ad4b60ba"
		testJob.Prompt = "Space panorama of moon-shaped burning wool, large as the moon, races towards  the blue planet earth, nasa earth, trending on artstation"
		responseObject.Completed_jobs = append(responseObject.Completed_jobs, testJob)

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
	img, err := os.Open("./model/images/" + inputstring2 + ".png") // temporary
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

	type newJob struct {
		Jobid      string `json:"jobid"`
		Prompt     string `json:"prompt"`
		Job_status string `json:"job_status"`
	}

	var j newJob
	// TODO: do this for real not just hardcode
	if inputstring2 == "752b5cd9013dcf3f6ebf577f99fa76adf4f32459" {

		j.Jobid = inputstring2
		j.Prompt = "3d render of celestial space nebula, cosmic, space station, unreal engine 3, photorealistic materials, trending on Artstation"
		j.Job_status = "completed" //pending? rejected?

		json.NewEncoder(w).Encode(j) // send back the json as a the response
	}
	if inputstring2 == "47d5090d689508acf1a6c29695e0d05ad4b60ba" {
		j.Jobid = inputstring2
		j.Prompt = "Space panorama of moon-shaped burning wool, large as the moon, races towards  the blue planet earth, nasa earth, trending on artstation"
		j.Job_status = "completed"   //pending? rejected?
		json.NewEncoder(w).Encode(j) // send back the json as a the response
	}
}

// Deals with POST requests made to the jobs endpoint (POST new jobs)
func HandleJobsApiPost(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Access-Control-Allow-Origin", "https://www.exia.art") // TESTING: allow CORS for testing purposes
	//w.Header().Set("Access-Control-Allow-Origin", "*")
	//w.Header().Set("Access-Control-Allow-Origin", r.Header.Get("Origin"))
	//w.Header().Set("Access-Control-Allow-Headers", r.Header.Get("Access-Control-Request-Headers"))
	w.Header().Set("Access-Control-Allow-Headers", "*")

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
	j.Jobid = "PLACEHOLDER_552byy5ebf577f99fa76asd4f32459" // Placeholder
	j.Prompt = jobRequest.Prompt                           // TODO: Validate and sanitize user input first
	j.Job_status = "accepted"                              //pending? rejected?

	json.NewEncoder(w).Encode(j) // send back the json as a the response
}
