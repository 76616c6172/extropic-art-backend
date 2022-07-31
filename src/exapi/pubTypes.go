// Webhandlers and data structures for the API
package exapi

type newjob struct {
	Prompt string `json:"prompt"`
}

//const MAX_PROMPT_LENGTH = 600 // Reject a new job posted by the view if longer than this value
//const MAX_PROMPT_LENGTH = 322 // Reject a new job posted by the view if longer than this value
const MAX_PROMPT_LENGTH = 360 // Reject a new job posted by the view if longer than this value
var GPU_STATUS = "offline"

/* This is the response object of he /api/0/jobs endpoint
For reference here is the Schema the client expects from /api/0/jobs?jobid=1
{
  "jobid": "1",
  "prompt": "Space wool bla bla, bla bla..",
  "job_status": "qeued",
  "iteration_status": "125",
  "iteration_max": "240",
} */
type apiJob struct {
	Jobid            string `json:"jobid"`
	Prompt           string `json:"prompt"`
	Job_status       string `json:"job_status"`
	Iteration_status int    `json:"iteration_status"`
	Iteration_max    int    `json:"iteration_max"`
	Img_path         string `json:"img_path"`
}

type jobResponse struct { // This is the response object sent back to the VIEW after POSTING a new job
	Jobid      int    `json:"jobid"`
	Prompt     string `json:"prompt"`
	Job_status string `json:"job_status"`
}

type status struct { // Schema for the status object returned by the status endpoint
	Gpu string `json:"gpu"`
	//Completed_jobs []apiJob `json:"completed_jobs"` //no longer needed
	Newest_jobid          string   `json:"newest_jobid"`
	Jobs_completed        int      `json"jobs_completed"`
	Jobs_queued           int      `json"jobs_queued"`
	Newest_completed_jobs []string `json"newest_completed_jobs"`
}
