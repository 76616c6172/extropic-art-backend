package exdb

/*
	// SCHEMA FROM THE JOB DB
	"jobid" INTEGER PRIMARY KEY AUTOINCREMENT,
	"model_pipeline" TEXT,
	"lock_seed" INTEGER,
	"seed" INTEGER,
	"guidance" INTEGER,
	"upscale" INTEGER,
	"prompt" TEXT,
	"status" TEXT,
	"job_params" JSON,
	"owner" TEXT,
	"iteration_status" INTEGER,
	"iteration_max" INTEGER,
	"time_created" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	"time_last_updated" TIMESTAMP,
	"time_completed" TIMESTAMP
*/

// Struct that matches the schema for jobs in jobs.db
type Job struct {
	Jobid             string `json:"jobid"`
	Model_pipeline    string `json:"model_pipeline"`
	Lock_seed         int    `json:"lock_seed"`
	Seed              int    `json:"seed"`
	Guidance          int    `json:"guidance"`
	Upscale           int    `json:"upscale"`
	Prompt            string `json:"prompt"`
	Status            string `json:"status"`
	Job_params        string `json:"job_params"` //changeme job params shoyld be a struct/object
	Owner             string `json:"owner"`
	Iteration_status  int    `json:"iteration_status"`
	Iteration_max     int    `json:"iteration_max"`
	Time_created      string `json:"time_created"`
	Time_last_updated string `json:"time_last_updated"`
	Time_completed    string `json:"time_completed"`
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
