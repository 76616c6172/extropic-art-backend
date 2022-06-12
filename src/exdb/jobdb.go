package exdb // Change this to internals later

import (
	"database/sql"
	"encoding/json"
	"log"
	"strconv"
	"time"

	// 3rd party packages
	_ "github.com/mattn/go-sqlite3"
)

//var JOBDB *sql.DB // This pointer is shared within the module to do database operations

// The parameters the diffusion model uses when running the job
type jobParam struct {
	placeholder int // TODO add more parameters here
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

// Initialize and connect to jobdb
func InitializeJobdb() *sql.DB {
	JOBDB, err := sql.Open("sqlite3", "../model/jobdb/jobs.db")
	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}

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
	defer stmnt.Close()

	_, err = stmnt.Exec()
	if err != nil {
		log.Fatal(err)
	}
	return JOBDB
}

// Adds a new job to the database
// Returns the jobid of the new job
// Returns error if job already exists or could not be created
// Assumes JOBDB is initialized first by calling InnitializeJobdb()
func InsertNewJob(db *sql.DB, prompt string, job_params interface{}) (int, error) {

	job_params_unm, err := json.Marshal(job_params) // Convert job_params to a string
	if err != nil {
		return -1, err
	}
	job_params_str := string(job_params_unm)

	stmnt, err := db.Prepare(` INSERT INTO "jobs" (prompt, status, job_params, iteration_status, iteration_max, time_created, time_last_updated, time_completed) values (?, ?, ?, ?, ?, ?, ?, ?);`) // Prepare the stament
	if err != nil {
		return -1, err
	}
	defer stmnt.Close()

	unixtime := strconv.Itoa(int(time.Now().Unix()))
	iteration_max := 240 // TODO: make this check if the user provided different values first

	// Execute the statement
	// for reference the jobs schema is:
	//INSERT INTO "jobs" (prompt, status, job_params, iteration_status, iteration_max, time_created, time_last_updated, time_completed)
	result, err := stmnt.Exec(prompt, "queued", job_params_str, 0, iteration_max, unixtime, unixtime, "")
	if err != nil {
		return -1, err
	}

	numberOfNewJob, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(numberOfNewJob), nil
}

/*
// Returns the newest job in the database
func GetLatestJob() (Job, error) {
	var j Job

	// rows, err := JOBDB.Query(`SELECT * FROM "jobs" ORDER BY jobid DESC LIMIT 1 ;`) // Query the database
	r, err := JOBDB.Query(`SELECT from sqlite_sequence where name='jobs';`)
	if err != nil {
		return j, err
	}
	defer r.Close()

	str := ""
	r.Scan(&str)
	fmt.Println(str)

	/*
		rows.Next()
		err = rows.Scan(&j)
		fmt.Println(j.Jobid, j.Prompt)
		if err != nil {
			return j, err
		}
	return j, err
}
*/

// Get job by jobid
// Returns the job with the given jobid
func GetJobByJobid(db *sql.DB, jobid int) (Job, error) {
	var j Job

	row, err := db.Query(`SELECT * FROM "jobs" WHERE jobid = ?;`, jobid) // Query the database
	if err != nil {
		return j, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&j.Jobid, &j.Prompt, &j.Status, &j.Job_params, &j.Iteration_status, &j.Iteration_max, &j.Time_created, &j.Time_last_updated, &j.Time_completed)
	if err != nil {
		return j, err
	}

	return j, err
}

// Dump entire jobdb into memory
// Returns a slice of all jobs
func GetAllJobs(db *sql.DB) ([]Job, error) {
	var jobs []Job

	rows, err := db.Query(`SELECT * FROM "jobs" ORDER BY jobid DESC;`) // Query the database
	if err != nil {
		return jobs, err
	}
	defer rows.Close()

	// Iterate over the rows and add them to the slice
	var j Job
	for rows.Next() {
		err = rows.Scan(&j.Jobid, &j.Prompt, &j.Status, &j.Job_params, &j.Iteration_status, &j.Iteration_max, &j.Time_created, &j.Time_last_updated, &j.Time_completed)
		if err != nil {
			return jobs, err
		}
		jobs = append(jobs, j)
	}

	return jobs, err
}

// Queries jobdb and returns a slice of all jobs in the set of [a,...,b]
// returns an error if the query fails
func GetjobsBetweenJobidXandJobidY(db *sql.DB, a int, b int) ([]Job, error) {
	if a == 0 { // FIXME: Handle the edge case, idk why 0 is not allowed in the query
		a = 1
	}

	rows, err := db.Query(`SELECT * FROM "jobs" WHERE jobid BETWEEN ? AND ?;`, a, b) // Query the database
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Unmarshal the results rnto a slice of Jobs
	var jobs []Job
	var tempJob Job
	for rows.Next() {

		err = rows.Scan(&tempJob.Jobid, &tempJob.Prompt, &tempJob.Status, &tempJob.Job_params, &tempJob.Iteration_status, &tempJob.Iteration_max, &tempJob.Time_created, &tempJob.Time_last_updated, &tempJob.Time_completed)
		if err != nil {
			return nil, err
		}
		jobs = append(jobs, tempJob)

	}

	return jobs, nil
}

// Returns oldest queued jon from the database
// Specifically, returns the job with the lowest jobid that has status "queued"
func GetOldestQueuedJob(db *sql.DB) (Job, error) {
	var j Job

	row, err := db.Query(`SELECT * FROM "jobs" WHERE status = "queued" ORDER BY jobid ASC LIMIT 1;`) // Query the database
	if err != nil {
		return j, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&j.Jobid, &j.Prompt, &j.Status, &j.Job_params, &j.Iteration_status, &j.Iteration_max, &j.Time_created, &j.Time_last_updated, &j.Time_completed)
	if err != nil {
		return j, err
	}

	return j, err
}

// Returns number of jobs from the db that have the status
func GetNumberOfJobsThatHaveStatus(db *sql.DB, status string) int {
	var jobs []Job
	var j Job

	row, err := db.Query(`SELECT * FROM "jobs" WHERE status = ?;`, status) // Query the database
	if err != nil {
		log.Println("Error in GetNumberOfCompletedJobs", err)
		return -1
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&j.Jobid, &j.Prompt, &j.Status, &j.Job_params, &j.Iteration_status, &j.Iteration_max, &j.Time_created, &j.Time_last_updated, &j.Time_completed)
		if err != nil {
			log.Println("Error in GetNumberOfCompletedJobs", err)
			return -1
		}
		jobs = append(jobs, j)
	}

	return len(jobs)
}

func GetLatestJobid(db *sql.DB) string {
	var j Job

	row, err := db.Query(`SELECT * FROM "jobs" ORDER BY jobid DESC LIMIT 1;`)
	if err != nil {
		log.Println("Error in GetNumberOfCompletedJobs", err)
		return "error getting last job"
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&j.Jobid, &j.Prompt, &j.Status, &j.Job_params, &j.Iteration_status, &j.Iteration_max, &j.Time_created, &j.Time_last_updated, &j.Time_completed)
	if err != nil {
		log.Println("Error in GetNumberOfCompletedJobs", err)
		return "error scanning last job"
	}

	return j.Jobid
}

// Returns the last couple jobs from the databasee that have the status you ask
// status = the status you are filtering for
// numberOfJobs = the amount of jobs the function will return
func GetNewestCoupleJobsThatHaveStatus(db *sql.DB, status string, numberOfJobs int) []string {
	var jobs []Job
	var j Job

	row, err := db.Query(`SELECT * FROM "jobs" WHERE status = ? ORDER BY jobid DESC LIMIT ?;`, status, numberOfJobs) // Query the database
	if err != nil {
		log.Println("Error in GetNewestFiveCompletedJobs", err)
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(&j.Jobid, &j.Prompt, &j.Status, &j.Job_params, &j.Iteration_status, &j.Iteration_max, &j.Time_created, &j.Time_last_updated, &j.Time_completed)
		if err != nil {
			log.Println("Error in GetNewestFiveCompletedJobs", err)
		}
		jobs = append(jobs, j)
	}

	var answer []string
	for _, job := range jobs {
		answer = append(answer, job.Jobid)
	}

	return answer
}

/ * to be continued */
JOBDB, jobid, newJobStatus, iterStatus)
func UpdateJobById(
