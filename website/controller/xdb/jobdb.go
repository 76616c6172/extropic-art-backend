package xdb // Change this to internals later

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"strconv"
	"time"

	// 3rd party packages
	_ "github.com/mattn/go-sqlite3"
)

var JOBDB *sql.DB // This pointer is shared within the module to do database operations

// The parameters the diffusion model uses when running the job
type jobParam struct {
	placeholder int

	// TODO add more parameters here
}

// the job schema based on the data from the jobdb
/* the schema for reference
"jobid" INTEGER PRIMARY KEY AUTOINCREMENT,
"prompt" TEXT,
"status" TEXT,
"job_params" JSON,
"iteration_status" INTEGER,
"iteration_max" INTEGER,
"time_created" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
"time_last_updated" TIMESTAMP,
"time_completed" TIMESTAMP
*/
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
}

func JobdbInit() {
	var err error
	JOBDB, err = sql.Open("sqlite3", "../model/jobdb/jobs.db")
	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}
	//defer db.Close() // If we wanted to close it

	// Create table if it doesn't exist
	stmnt, err := JOBDB.Prepare(`
CREATE TABLE IF NOT EXISTS "jobs" (
 "jobid" INTEGER PRIMARY KEY AUTOINCREMENT,
 "prompt" TEXT,
 "status" TEXT,
 "job_params" JSON,
 "iteration_status" INTEGER,
 "iteration_max" INTEGER,
 "time_created" TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 "time_last_updated" TIMESTAMP,
 "time_completed" TIMESTAMP
	);`)
	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}
	_, err = stmnt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

// Adds a new job to the database
// Returns the jobid of the new job
// Returns error if job already exists or could not be created
// Assumes JOBDB is initialized first by calling JobdbInit()
func InsertNewJob(prompt string, job_params interface{}) (int, error) {

	job_params_unm, err := json.Marshal(job_params) // Convert job_params to a string
	if err != nil {
		return -1, err
	}
	job_params_str := string(job_params_unm)

	stmnt, err := JOBDB.Prepare(` INSERT INTO "jobs" (prompt, status, job_params, iteration_status, iteration_max, time_created, time_last_updated, time_completed) values (?, ?, ?, ?, ?, ?, ?, ?);`) // Prepare the stament
	if err != nil {
		return -1, err
	}

	unixtime := strconv.Itoa(int(time.Now().Unix()))

	// Execute the statement
	// for reference the jobs schema is:
	//INSERT INTO "jobs" (jobid, prompt, status, job_params, iteration_status, iteration_max, time_created, time_last_updated, time_completed)
	result, err := stmnt.Exec(prompt, "queued", job_params_str, 0, 0, unixtime, unixtime, "")
	if err != nil {
		return -1, err
	}

	numberOfNewJob, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(numberOfNewJob), nil
}

// Queries jobdb and returns a slice of all jobs in the set of [a,...,b]
// returns an error if the query fails
func GetJobsBetweenTwoJobids(a int, b int) ([]Job, error) {
	if a == 0 { // FIXME: Handle the edge case, idk why 0 is not allowed in the query
		a = 1
	}

	rows, err := JOBDB.Query(`SELECT * FROM "jobs" WHERE jobid BETWEEN ? AND ?;`, a, b) // Query the database
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Unmarshal the results rnto a slice of Jobs
	var jobs []Job
	var tempJob Job
	var i int
	maxNumberOfResults := b - a + 1
	for { // iterate through the rows until we get to the end
		rows.Next()

		err = rows.Scan(&tempJob.Jobid, &tempJob.Prompt, &tempJob.Status, &tempJob.Job_params, &tempJob.Iteration_status, &tempJob.Iteration_max, &tempJob.Time_created, &tempJob.Time_last_updated, &tempJob.Time_completed)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, tempJob)

		i++
		if i > maxNumberOfResults-1 { // stop scanning when we have reached the end of the results
			break
		}
	}

	return jobs, nil
}

// Called by the main function so we can test the module
func EntryPointForTesting() {

	InsertNewJob("test", "")

	lastJobs, err := GetJobsBetweenTwoJobids(21, 32)
	if err != nil {
		log.Fatal(err)
	}
	for i, v := range lastJobs {
		fmt.Println(i, v.Prompt, v.Jobid)
	}

}
