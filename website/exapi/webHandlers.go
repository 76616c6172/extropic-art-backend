package exapi

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"

	"project-exia-monorepo/website/exdb"
)

const MAX_PROMPT_LENGTH = 600 // Reject a new job posted by the view if longer than this value

// Deals with requests to the status endpoint
func HandleStatusRequest(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" { // send back the response

		var responseObject = status{ // We append to this object as we go and send it back to the client
			Gpu:          "ready", // busy, offline
			Newest_jobid: "",
		}

		// TODO: dumping entire database into memory every time won't scale so get only the last 10!
		allJobs, err := exdb.GetAllJobs()
		if err != nil {
			log.Println(err)
			//TODO: et response code to 500 to indicate server error
			json.NewEncoder(w).Encode(responseObject) // send back the json as a the response
		}

		/*
			var respJob apiJob
			for i := 0; i < 10; i++ {
				if i < len(allJobs) { // sanity check if we have a fresh database with less than 10 entries
					respJob.Jobid = allJobs[i].Jobid
					respJob.Prompt = allJobs[i].Prompt
					respJob.Job_status = allJobs[i].Status
					respJob.Iteration_max = allJobs[i].Iteration_max
					respJob.Iteration_status = allJobs[i].Iteration_status
					respJob.Img_path = "https://exia.art/api/0/img?jobid=" + respJob.Jobid
					responseObject.Completed_jobs = append(responseObject.Completed_jobs, respJob)
				}
			}
		*/
		responseObject.Newest_jobid = allJobs[0].Jobid

		json.NewEncoder(w).Encode(responseObject) // send back the json as a the response
	}

}

// Obtains the coorect image for a given job and sends it back
func HandleImgRequests(w http.ResponseWriter, r *http.Request) {

	input := fmt.Sprintln(r.URL)
	inputstring := strings.TrimLeft(input, "/api/0/img?jobid=")
	inputstring2 := strings.TrimSpace(inputstring)

	//img, err := os.Open("./model/images/" + imgRequest.Jobid + ".png") // for now just get this image for testing
	img, err := os.Open("../model/images/" + inputstring2 + ".png") // temporary
	// TODO: Actually lookfor the image in SQLite database
	// img, err := os.Open("./model/images/" + imgRequest.Jobid + ".png") // for now just get this image for testing
	if err != nil {
		img, err = os.Open("../model/images/placeholder.png") // send placeholder image
		log.Println(err)                                      // perhaps handle this nicer
	}
	defer img.Close()

	w.Header().Set("Content-Type", "image/png") // <-- set the content-type header
	_, err = io.Copy(w, img)                    // send the image
	if err != nil {
		log.Println(err)
	}
}

// Sends back metadata for one or multiple jobs
// takes /api/0/jobs=?jobid=1 OR /api/0/jobx=3&joby=4
func HandleJobsApiGet(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		str_a := strings.TrimLeft(r.URL.String(), "/api/0/")

		// CHeck if we have a jobid
		if strings.Contains(str_a, "jobs?jobid") {

			// 1. Determine the jobid from the request
			input := fmt.Sprintln(r.URL)
			input2 := strings.TrimLeft(input, "/api/0/jobs?jobid=")
			input3 := strings.TrimSpace(input2)

			// 2. Sanitize the input
			sanitized_input, err := strconv.Atoi(input3)
			if err != nil {
				log.Println(err)
				return
			}

			// 2. Get the row from the database
			var realJob exdb.Job
			realJob, err = exdb.GetJobByJobid(sanitized_input)
			if err != nil {
				log.Println(err)
				return
			}

			// 3. Build the response object
			var responseJob apiJob
			responseJob.Jobid = realJob.Jobid
			responseJob.Prompt = realJob.Prompt
			responseJob.Job_status = realJob.Status
			responseJob.Iteration_status = realJob.Iteration_status
			responseJob.Iteration_max = realJob.Iteration_max
			responseJob.Img_path = "https://exia.art/api/0/img?jobid=" + responseJob.Jobid

			// 4. Send back the response
			json.NewEncoder(w).Encode(responseJob)
			return
		} else { // assume the client wants multiple jobs jobx=1&joby=2
			b_str := strings.TrimLeft(str_a, "jobs?jobx=")
			c_str := strings.Split(b_str, "&")
			x_str := strings.TrimRight(c_str[0], "&")
			y_str := strings.TrimLeft(c_str[1], "joby=")

			// 1. Sanitize the input
			x, err := strconv.Atoi(x_str)
			if err != nil {
				log.Println(err)
				return
			}
			y, err := strconv.Atoi(y_str)
			if err != nil {
				log.Println(err)
				return
			}
			var realJobs []exdb.Job

			// 2. Query the database
			realJobs, err = exdb.GetJobsByXY(x, y)
			if err != nil {
				log.Println(err)
				return
			}

			// 3. Build the resposne
			var responseJobs []apiJob
			var responseJob apiJob

			for _, v := range realJobs {
				responseJob.Jobid = v.Jobid
				responseJob.Prompt = v.Prompt
				responseJob.Job_status = v.Status
				responseJob.Iteration_status = v.Iteration_status
				responseJob.Iteration_max = v.Iteration_max
				responseJob.Img_path = "https://exia.art/api/0/img?jobid=" + responseJob.Jobid
				responseJobs = append(responseJobs, responseJob)
			}

			// 4. Send back the response
			json.NewEncoder(w).Encode(responseJobs)
			return
		}
	}

}

// Deals with POST requests made to the jobs endpoint (POST new jobs)
// Sends the job to jobs.db
func HandleJobsApiPost(w http.ResponseWriter, r *http.Request) {

	jsonDecoder := json.NewDecoder(r.Body)
	var jobRequest newjob
	err := jsonDecoder.Decode(&jobRequest)
	if err != nil {
		log.Println(err) // maybe handle this better
		return
	}

	// 1. Santizie and check the input
	if len(jobRequest.Prompt) > MAX_PROMPT_LENGTH {
		log.Println("JobsApiPost: Prompt is of length < 1")
		return
	}
	if len(jobRequest.Prompt) < 1 { // send back an error response
		var j jobResponse
		j.Jobid = -1                 // Placeholder
		j.Prompt = jobRequest.Prompt // TODO: Validate and sanitize user input first
		j.Job_status = "Rejected"    //pending? rejected?

		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(j) // send back the json as a the response
		return
	}

	// 2. Create the job in the database
	newJobid, err := exdb.InsertNewJob(jobRequest.Prompt, "")
	if err != nil {
		log.Println("JobsApiPost: err")
		return
	}

	// 3. Send back the jobid of the newly created job to the client

	var j jobResponse
	j.Jobid = newJobid           // Placeholder
	j.Prompt = jobRequest.Prompt // TODO: Validate and sanitize user input first
	j.Job_status = "accepted"    //pending? rejected?

	json.NewEncoder(w).Encode(j) // send back the json as a the response
}