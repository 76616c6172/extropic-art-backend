package exapi

import (
	"database/sql"
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

// Sends back the status response
func HandleStatusRequest(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var statusResponse status

		statusResponse.Gpu = GPU_STATUS
		statusResponse.Jobs_completed = exdb.GetNumberOfJobsThatHaveStatus(db, "completed")
		statusResponse.Jobs_queued = exdb.GetNumberOfJobsThatHaveStatus(db, "queued")
		statusResponse.Newest_completed_jobs = exdb.GetNewestCoupleJobsThatHaveStatus(db, "completed", 3)
		statusResponse.Newest_jobid = exdb.GetLatestJobid(db)

		json.NewEncoder(w).Encode(statusResponse) // send back the response
	}

}

// Obtains the coorect image for a given job and sends it back
func HandleImgRequests(w http.ResponseWriter, r *http.Request) {

	input := fmt.Sprintln(r.URL)
	inputstring := strings.TrimLeft(input, "/api/0/img?jobid=")
	inputstring2 := strings.TrimSpace(inputstring)

	img, err := os.Open("../model/images/" + inputstring2 + ".png")
	if err != nil {
		img, err = os.Open("../model/images/placeholder.png") // send placeholder image
		log.Println(err)
	}
	defer img.Close()

	w.Header().Set("Content-Type", "image/png") // <-- set the content-type header

	_, err = io.Copy(w, img) // send the image
	if err != nil {
		log.Println(err)
	}
}

// Finishes parsing the query param from the web request, then looks up the requested job
// in the database by jobid and sends back the job to the client.
func sendBackOneJob(db *sql.DB, w http.ResponseWriter, r *http.Request, str_a string) {

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
		realJob, err = exdb.GetJobByJobid(db, sanitized_input)
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
	}
}

// Sends back list of jobs to the client
func sendBackMultipleJobs(db *sql.DB, w http.ResponseWriter, r *http.Request, str_a string) {

	// 0. Finish parsing the query parameters
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
	realJobs, err = exdb.GetjobsBetweenJobidXandJobidY(db, x, y)
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

// Sends back metadata for one or multiple jobs
// takes /api/0/jobs=?jobid=1 OR /api/0/jobx=3&joby=4
func HandleJobsApiGet(db *sql.DB, w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {

		str_a := strings.TrimLeft(r.URL.String(), "/api/0/")
		if strings.Contains(str_a, "jobs?jobid") { // assume the client wants only one job back
			sendBackOneJob(db, w, r, str_a)
			return
		}

		sendBackMultipleJobs(db, w, r, str_a) // assume the client wants multiple jobs jobx=1&joby=2
	}
}

// Deals with POST requests made to the jobs endpoint (POST new jobs)
// Adds job to the database with the "queued" status
func HandleJobsApiPost(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	var j jobResponse

	jsonDecoder := json.NewDecoder(r.Body)
	var jobRequest newjob
	err := jsonDecoder.Decode(&jobRequest)
	if err != nil {
		log.Println("HandleJobsApiPost: Error decoding request", err)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(j) // send back the json as a the response
		return
	}

	// 1. Santizie and check the input
	if len(jobRequest.Prompt) > MAX_PROMPT_LENGTH {
		log.Println("HandleJobsApiPost: Error large prompt length")
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(j) // send back the json as a the response
		return
	}

	if len(jobRequest.Prompt) < 1 {
		j.Jobid = -1
		j.Prompt = jobRequest.Prompt
		j.Job_status = "Rejected"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(j)
		log.Println("HandleJobsApiPost: Error small prompt length")
		return
	}

	// 2. Create the job in the database
	newJobid, err := exdb.InsertNewJob(db, jobRequest.Prompt, "")
	if err != nil {
		log.Println("HandleJobsApiPost: Error inserting new job", err)
		return
	}

	// 3. Send back the jobid of the newly created job to the client
	j.Jobid = newJobid
	j.Prompt = jobRequest.Prompt
	j.Job_status = "accepted"

	json.NewEncoder(w).Encode(j) // send back the json as a the response
}
