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

// Used to decide to send back JPG or PNG
const (
	JPG = iota
	PNG
)

const MAXIMUM_NUMBER_OF_JOBS_IN_QUEUE = 50

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

// Answers requests to version 2 of the status api
func Handle_status_api_endpoint_version_2(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		type status_response struct {
			Newest_completed_job string `json:"newest_completed_job"`
		}
		var status status_response
		//int Newest_completed_job: 11,
		status.Newest_completed_job = exdb.GetNewestCoupleJobsThatHaveStatus(db, "completed", 1)[0]
		json.NewEncoder(w).Encode(status) // send back the response
	}
}

// Sends back PNG to the client
// expects requests like: https://exia.art/api/0/img?type=full?jobid=10
func SendBackPNGorJPG(w http.ResponseWriter, r *http.Request, imageType int) {

	if imageType == PNG {
		jobidFromQuery := strings.TrimPrefix(r.URL.RawQuery, "type=full?jobid=")

		img, err := os.Open("../model/images/pngs/" + jobidFromQuery + ".png")
		if err != nil {
			img, err = os.Open("../model/images/pngs/placeholder.png")
			log.Println(err)
		}
		defer img.Close()

		_, err = io.Copy(w, img) // send the image
		if err != nil {
			log.Println(err)
		}
		return
	}

	if imageType == JPG {
		jobidFromQuery := strings.TrimPrefix(r.URL.RawQuery, "type=thumbnail?jobid=")

		img, err := os.Open("../model/images/jpgs/" + jobidFromQuery + ".jpg")
		if err != nil {
			img, err = os.Open("../model/images/jpgs/placeholder.jpg")
			log.Println(err)
		}
		defer img.Close()

		_, err = io.Copy(w, img) // send the image
		if err != nil {
			log.Println(err)

		}
	}
}

// Obtains the coorect image for a given job and sends it back
func HandleImgRequests(w http.ResponseWriter, r *http.Request) {

	if strings.Contains(r.URL.RawQuery, "type=full") { // client wants a PNG
		SendBackPNGorJPG(w, r, PNG)
	} else if strings.Contains(r.URL.RawQuery, "type=thumbnail") { // client wants JPG
		SendBackPNGorJPG(w, r, JPG)
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

// Returns sanitized user input that won't yaml
func SanitizeUserPrompt(p string) string {
	p = strings.Replace(p, " :", "", -1)
	p = strings.Replace(p, ": ", "", -1)

	p = strings.Replace(p, " -", "", -1)
	p = strings.Replace(p, "- ", "", -1)

	//p = strings.Replace(p, "\"", "", -1)
	p = strings.Replace(p, "'", "", -1)
	p = strings.Replace(p, "`", "", -1)
	p = strings.Replace(p, "\n", "", -1)
	p = strings.Replace(p, "\t", "", -1)
	p = strings.Replace(p, "\r", "", -1)
	p = strings.Replace(p, "(", "", -1)
	p = strings.Replace(p, ")", "", -1)
	p = strings.Replace(p, "[", "", -1)
	p = strings.Replace(p, "]", "", -1)
	p = strings.Replace(p, "—", "", -1)

	return p
}

// Returns false if input is not valid and would be unsafe/yaml breaking
// TODO: write some good rules here
func InputIsValid(input string) bool {

	// If prompt contains no weights that is fine too
	if !strings.Contains(input, ":") && !strings.Contains(input, "|") {
		return true

	}

	// Check if every prompt has a weight
	// Maybe we don't need this!
	collonAmount := strings.Count(input, ":")
	barAmount := strings.Count(input, "|")
	if collonAmount != barAmount+1 {
		return false
	}

	// Split on |
	s := strings.Split(input, "|")
	var weightsTotal float64 = 0

	for _, ss := range s {
		// Split on :-
		s3 := strings.Split(ss, ":")
		strNumber := s3[len(s3)-1]

		// Convert number to float
		strNumNoSpace := strings.TrimSpace(strNumber)
		if number, err := strconv.ParseFloat(strNumNoSpace, 64); err == nil {
			//fmt.Println(number)
			weightsTotal += number
		} else {
			//println(err)
			return false
		}
	}

	//println("d: weights total=", weightsTotal)
	if weightsTotal == 1 {
		return true
	}

	return false
}

// Deals with POST requests made to the jobs endpoint (POST new jobs)
// Checks if the posted job is valid and then
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

	// 1. Santizie the input
	jobRequest.Prompt = SanitizeUserPrompt(jobRequest.Prompt)
	if !InputIsValid(jobRequest.Prompt) {
		log.Println("HandleJobsApiPost: Error decoding request", err)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(j) // send back the json as a the response
		return
	}

	if len(jobRequest.Prompt) > MAX_PROMPT_LENGTH {
		log.Println("HandleJobsApiPost: Error large prompt length")
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(j) // send back the json as a the response
		return
	}

	if len(jobRequest.Prompt) < 1 || len(jobRequest.Prompt) > MAX_PROMPT_LENGTH {
		j.Jobid = -1
		j.Prompt = jobRequest.Prompt
		j.Job_status = "Rejected"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(j)
		log.Println("HandleJobsApiPost: Prompt length out of bounds")
		return
	}

	// 1.5 Check that we're not exceeding the queued job limit
	numberOfQUeuedJobs := exdb.GetNumberOfQueuedJobs(db)
	if numberOfQUeuedJobs > MAXIMUM_NUMBER_OF_JOBS_IN_QUEUE {
		j.Jobid = -1
		j.Prompt = jobRequest.Prompt
		j.Job_status = "Rejected"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(j)
		log.Println("HandleJobsApiPost: Error exceeded maximum jobs in queue")
	}
	// HERE

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

// Takes a raw list of jobs from the DB and builds a lighter list of job objects meant for the view
// Without the metadata the view doesn need.
func buildApiJobListForTheView(jobs []exdb.Job) []apiJob {
	var newJobList []apiJob

	for _, j := range jobs {
		newJobObject := apiJob{
			Jobid:            j.Jobid,
			Prompt:           j.Prompt,
			Job_status:       j.Status,
			Iteration_status: j.Iteration_status,
			Iteration_max:    j.Iteration_max,
			Img_path:         "https://exia.art/api/0/img?type=thumbnail?jobid=" + j.Jobid,
		}
		newJobList = append(newJobList, newJobObject)
	}

	fmt.Println("DB Input:", jobs)
	println()
	fmt.Println("List", newJobList)

	return newJobList
}

// Send back the prompts that are in the current queue
func HandleQueueRequest(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" { // send back the response

		jobsInQueue, err := exdb.GetAllJobsInQueue(db)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			println(err)
			return
		}

		listOfJobObjectsForTheView := buildApiJobListForTheView(jobsInQueue)
		json.NewEncoder(w).Encode(listOfJobObjectsForTheView) // send back the json as a the response
	}
}
