package internals

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
)

const GPU_STATUS string = "offline" // can be offline, online, or busy

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
		testJob.Jobid = "752byy5cd9013dcf3f6ebf577f99fa76adf4f32459"
		testJob.Prompt = "3d render of celestial space nebula, cosmic, space station, unreal engine 3, photorealistic materials, trending on Artstation"

		var responseObject = status{
			Gpu: "ready", // busy, offline
		}

		// Just testing
		responseObject.Completed_jobs = append(responseObject.Completed_jobs, testJob)
		testJob.Jobid = "847d5090d689508acf1a6c29695e0d05ad4b60ba"
		testJob.Prompt = "Space panorama of moon-shaped burning wool, large as the moon, races towards  the blue planet earth, nasa earth, trending on artstation"
		responseObject.Completed_jobs = append(responseObject.Completed_jobs, testJob)

		json.NewEncoder(w).Encode(responseObject) // send back the json as a the response
	}

}

// Obtains the coorect image for a given job and sends it back
func HandleImgRequests(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // TESTING: allow CORS for testing purposes

	jsonDecoder := json.NewDecoder(r.Body)
	var imgRequest imgRequest
	err := jsonDecoder.Decode(&imgRequest)
	if err != nil {
		log.Println(err) // maybe handle this better
		return
	}

	// TODO: Actually lookfor the image in SQLite database
	img, err := os.Open("./model/images/" + imgRequest.Jobid + ".png") // for now just get this image for testing
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Access-Control-Allow-Origin", "*") // TESTING: allow CORS for testing purposes
	w.Header().Set("Content-Type", "image/png")        // <-- set the content-type header
	_, err = io.Copy(w, img)                           // send the image
	if err != nil {
		log.Println(err)
	}
}

// Deals with POST requests made to the jobs endpoint (POST new jobs)
func HandleJobsApiPost(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*") // TESTING: allow CORS for testing purposes

	jsonDecoder := json.NewDecoder(r.Body)
	var jobRequest job
	err := jsonDecoder.Decode(&jobRequest)
	if err != nil {
		log.Println(err) // maybe handle this better
		return
	}

	type newJob struct {
		Jobid      string `json:"jobid"`
		Prompt     string `json:"prompt"`
		Job_status string `json:"job_status"`
	}

	var j newJob
	j.Jobid = "552byy5cd9013dcf3f6ebf577f99fa76asd4f32459" // Placeholder
	j.Prompt = jobRequest.Prompt                           // TODO: Validate and sanitize user input first
	j.Job_status = "accepted"                              //pending? rejected?

	w.Header().Set("Access-Control-Allow-Origin", "*") // TESTING: allow CORS for testing purposes
	json.NewEncoder(w).Encode(j)                       // send back the json as a the response
}
