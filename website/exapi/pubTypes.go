// Contains all data types sent between the VIEW and the CONTROLLER
package exapi

type newjob struct {
	Prompt string `json:"prompt"`
}

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

// This is the response object sent back to the VIEW after POSTING a new job
type jobResponse struct {
	Jobid      int    `json:"jobid"`
	Prompt     string `json:"prompt"`
	Job_status string `json:"job_status"`
}

// Schema for the status object returned by the status endpoint
type status struct {
	Gpu            string   `json:"gpu"`
	Completed_jobs []apiJob `json:"completed_jobs"`
	//Description string `json:"Description"`
}
