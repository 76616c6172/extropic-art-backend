// Webhandlers and data structures for the API
package exapi

// DEPRECATED
// schema for a job posted by the old frontend
type newjob struct {
	Prompt string `json:"prompt"`
}

// schema for a job posted by the react frontend
type jobPostingFromApiVersion1 struct {
	ModelPipeline  int    `json:"model_pipeline"`
	Resolution     int    `json:"resolution"`
	Prompt         string `json:"prompt"`
	IsLockedSeed   bool   `json:"lock_seed"`
	Seed           string `json:"seed"`
	IsPrePrompt    bool   `json:"pre_prompt"`
	IsHighGuidance bool   `json:"high_guidance"`
	IsUpscale      bool   `json:"upscale"`
}

/*
	This is the response object of he /api/0/jobs endpoint

For reference here is the Schema the client expects from /api/0/jobs?jobid=1

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
	Seed             int    `json:"seed"`
	Owner            string `json:"owner"`
	IsHighGuidance   int    `json:"high_guidance"`
	IsUpscale        int    `json:"upscale"`
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
