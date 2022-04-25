package internals

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const GPU_STATUS string = "offline" // can be offline, online, or busy

type job struct {
	Prompt string `json:"prompt"`
	Jobid  string `json:"jobid"`
}

type imgRequest struct {
	Jobid string `json:"jobid"`
}

// Checks if a GPU is available and returns: online, offline, or busy
func Get_gpu_status() string {
	return GPU_STATUS
}

// Obtains the coorect image for a given job and sends it back
func HandleImgRequests(w http.ResponseWriter, r *http.Request) {

	jsonDecoder := json.NewDecoder(r.Body)
	var imgRequest imgRequest
	err := jsonDecoder.Decode(&imgRequest)
	if err != nil {
		log.Println(err) // maybe handle this better
		return
	}

	fmt.Println(imgRequest.Jobid)

	// To do, actually lookfor the image in a DB
	img, err := os.Open("./model/images/" + imgRequest.Jobid + ".png") // for now just get this image for testing
	if err != nil {
		log.Fatal(err) // perhaps handle this nicer
	}
	defer img.Close()
	w.Header().Set("Content-Type", "image/png") // <-- set the content-type header
	io.Copy(w, img)
}

// Obtains and sends back the current status like info and current jobs as json
func ReturnInfoAboutAllJobs(w http.ResponseWriter) {
	type status struct {
		Current_job    string   `json:"current_job"`
		Completed_jobs []string `json:"completed_jobs"`
		//Description string `json:"Description"`
	}

	var all_response = status{
		Completed_jobs: []string{"1650746663"},
	}

	json.NewEncoder(w).Encode(all_response) // send back the json as a the response
}

func HandleNewJobRequest(w http.ResponseWriter, r *http.Request) {
	jsonDecoder := json.NewDecoder(r.Body)
	var jobRequest job
	err := jsonDecoder.Decode(&jobRequest)
	if err != nil {
		log.Println(err) // maybe handle this better
		return
	}

	type message struct {
		Jobid      string `json:"jobid"`
		Job_status string `json:"job_status"`
		//Description string `json:"Description"`
	}

	var responsePost message
	responsePost.Jobid = jobRequest.Jobid
	responsePost.Job_status = "accepted" //pending? rejected?

	json.NewEncoder(w).Encode(responsePost) // send back the json as a the response

	fmt.Println(jobRequest.Prompt) // debug
	fmt.Println(jobRequest.Jobid)  // debug

}
