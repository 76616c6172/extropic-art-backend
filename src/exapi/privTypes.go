package exapi

// Job for internal use by gpu-worker + scheduler
//type Job struct {
//	Jobid            string `json:"jobid"`
//	Prompt           string `json:"prompt"`
//	Job_status       string `json:"job_status"`
//	Job_params       string `json:"job_params"`
//	Iteration_status int    `json:"iteration_status"`
//	Iteration_max    int    `json:"iteration_max"`
//}

type Job struct {
	Jobid             string `json:"jobid"`
	Prompt            string `json:"prompt"`
	Status            string `json:"status"`
	Job_params        string `json:"job_params"` //changeme job params shoyld be a struct/object
	Iteration_status  int    `json:"iteration_status"`
	Iteration_max     int    `json:"iteration_max"`
	Time_created      string `json:"time_created"`
	Time_last_updated string `json:"time_last_updated"`
	Time_completed    string `json:"time_completed"`
	Secret            string `json:"secret"`
}

// Form sent by gpu-workers to register with the scheduler
type WorkerRegistrationForm struct {
	Secret string `json:"secret"` // TODO: Authenticate better
}

// Sent back to the scheduler after receiving a job posting
type WorkerResponse struct {
	Job_accepted bool `json:"job_accepted"`
}
