package exapi

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"extropic-art-backend/src/exdb"
)

// Used to decide to send back JPG or PNG
const (
	JPG = iota
	PNG
)

const MAXIMUM_NUMBER_OF_JOBS_IN_QUEUE = 25
const MAXIMUM_DAILY_USES = 100

var FREE_USES_REMAINING = MAXIMUM_DAILY_USES

// Sends back the status response
func HandleStatusRequest(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var statusResponse status

		statusResponse.Gpu = GPU_STATUS
		statusResponse.Jobs_completed = exdb.GetNumberOfJobsThatHaveStatus(db, "completed")
		statusResponse.Jobs_queued = exdb.GetNumberOfJobsThatHaveStatus(db, "queued")
		// statusResponse.Newest_completed_jobs = exdb.GetNewestCoupleJobsThatHaveStatus(db, "completed", 3) //dprecated
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
		status.Newest_completed_job = exdb.GetNewestCompletedJob(db, "completed").Jobid
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
		responseJob.Seed = realJob.Seed
		responseJob.Prompt = realJob.Prompt
		responseJob.Job_status = realJob.Status
		responseJob.Iteration_status = realJob.Iteration_status
		responseJob.Iteration_max = realJob.Iteration_max
		responseJob.Img_path = "https://exia.art/api/0/img?jobid=" + responseJob.Jobid
		responseJob.Model_id = realJob.Model_pipeline

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

var Mutex = &sync.Mutex{}

// Deals with requests made to the api endpoint /api/1/jobs
func HandleJobsApiPostVersion2(db *sql.DB, w http.ResponseWriter, r *http.Request) {
	Mutex.Lock()
	defer Mutex.Unlock()

	if !GPU_IS_ONLINE || FREE_USES_REMAINING <= 0 {
		w.WriteHeader(http.StatusForbidden)
		return
	}

	// 1. Check if the job post is authorized
	var newJobPosting jobPostingFromApiVersion1
	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&newJobPosting)
	if err != nil {
		log.Println("error decoding request", err)
		w.Write([]byte("500 - Something bad happened!"))
		return
	}

	//fmt.Println("[Job posted]: ", newJobPosting.Prompt, "\n[User info hash]: ", sha1.Sum([]byte(r.UserAgent())))

	// 1. Santizie the input

	var jobResponse jobResponse
	newJobPosting.Prompt = SanitizeUserPrompt(newJobPosting.Prompt)
	if !InputIsValid(newJobPosting.Prompt) {
		log.Println("HandleJobsApiPost: Error decoding request", err)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(jobResponse) // send back the json as a the response
		return
	}

	if len(newJobPosting.Prompt) > MAX_PROMPT_LENGTH {
		log.Println("HandleJobsApiPost: Error large prompt length")
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(jobResponse) // send back the json as a the response
		return
	}

	if len(newJobPosting.Prompt) < 1 || len(newJobPosting.Prompt) > MAX_PROMPT_LENGTH {
		jobResponse.Jobid = -1
		jobResponse.Prompt = newJobPosting.Prompt
		jobResponse.Job_status = "Rejected"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(jobResponse)
		log.Println("HandleJobsApiPost: Prompt length out of bounds")
		return
	}

	// 1.5 Check that we're not exceeding the queued job limit
	numberOfQUeuedJobs := exdb.GetNumberOfQueuedJobs(db)
	if numberOfQUeuedJobs > MAXIMUM_NUMBER_OF_JOBS_IN_QUEUE {
		jobResponse.Jobid = -1
		jobResponse.Prompt = newJobPosting.Prompt
		jobResponse.Job_status = "Rejected"
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		json.NewEncoder(w).Encode(jobResponse)
		log.Println("HandleJobsApiPost: Error exceeded maximum jobs in queue")
	}

	newJobPosting.Seed = strings.TrimSpace(newJobPosting.Seed)

	// 2. Create the job in the database
	var lockSeed int
	var seedNumber int
	if newJobPosting.IsLockedSeed {
		lockSeed = 1
		if seedNumber, err = strconv.Atoi(newJobPosting.Seed); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("500 - Something bad happened!"))

			// Check to make sure seed is within acceptable bounds
			if seedNumber < 0 || seedNumber > 4294967295 {
				seedNumber = createRandomSeed()
			}
		}

	} else {
		seedNumber = createRandomSeed()
	}

	var guidance int
	if newJobPosting.IsHighGuidance {
		guidance = 1
	}
	var upscale int
	if newJobPosting.IsUpscale {
		upscale = 1
	}

	// dirty hack to get preprompt metadata across
	var User string
	if newJobPosting.IsPrePrompt {
		User = "anonymous_with_pre_prompt"
	} else {
		User = "anonymous_no_preprompt"
	}
	if newJobPosting.ModelPipeline != 1 { // hack to only do pre prompting for midjourny v4 finetune
		newJobPosting.IsPrePrompt = false
		User = "anonymous_no_pre_prompt"
	}

	jobResponse.Jobid, err = exdb.InsertNewJob(db, strconv.Itoa(newJobPosting.ModelPipeline), lockSeed, seedNumber, guidance, upscale, newJobPosting.Prompt, User, newJobPosting.Resolution)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("500 - Something bad happened!"))
		log.Println("HandleJobsApiPost: Error inserting new job", err)
		return
	}

	FREE_USES_REMAINING--
	w.WriteHeader(http.StatusAccepted)
	json.NewEncoder(w).Encode(jobResponse)
}

// createRandomSeed emits a between 0 and 4294967295
func createRandomSeed() int {
	return rand.Intn(4294967295)
}

// Takes a raw list of jobs from the DB and builds a lighter list of job objects meant for the view
// Without the metadata the view doesn need.
func buildApiJobListForTheView(jobs []exdb.Job) []apiJob {
	var newJobList []apiJob

	for _, j := range jobs {
		newJobObject := apiJob{
			Jobid:            j.Jobid,
			Prompt:           j.Prompt,
			Seed:             j.Seed,
			IsHighGuidance:   j.Guidance,
			IsUpscale:        j.Upscale,
			Job_status:       j.Status,
			Iteration_status: j.Iteration_status,
			Iteration_max:    j.Iteration_max,
			Img_path:         "https://exia.art/api/0/img?type=thumbnail?jobid=" + j.Jobid,
		}
		newJobList = append(newJobList, newJobObject)
	}

	return newJobList
}

type queueResponse struct {
	CurrentJobQueue   []apiJob `json:"queue"`
	FreeUsesRemaining int      `json:"freeUsesRemaining"`
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
		response := queueResponse{
			CurrentJobQueue:   listOfJobObjectsForTheView,
			FreeUsesRemaining: FREE_USES_REMAINING,
		}
		json.NewEncoder(w).Encode(response) // send back the json as a the response
	}
}
