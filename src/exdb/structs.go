package exdb

// Struct that matches the schema for jobs in jobs.db
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

// Struct that matches the schema for workers in workers.db
type Worker struct {
	Worker_id                string `json:"worker_id"`
	Worker_ip                string `json:"worker_ip"`
	Worker_Busy              int    `json:"Worker_Busy"`
	Worker_current_job       int    `json:"worker_current_job"`
	Worker_last_health_check int    `json:"worker_last_health_check"`
	Worker_time_created      int    `json:"worker_time_created"`
	Worker_secret            string `json:"worker_secret"`
	Worker_type              int    `json:"secret"`
}
