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

// The parameters the diffusion model uses when running the job
type jobParam struct {
	placeholder int // TODO add more parameters here
}

// Initialize and connect to jobdb
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
func InitializeJobdb() *sql.DB {
	JOBDB, err := sql.Open("sqlite3", "../model/jobdb/jobs.db")
	if err != nil {
		log.Fatal("error opening jobdb", err) // TODO: Maybe handle this better
	}

	// Create table if it doesn't exist
	stmnt, err := JOBDB.Prepare(`
	CREATE TABLE IF NOT EXISTS "jobs" (
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
	);`)
	if err != nil {
		log.Fatal("error initializing jobdb", err)
	}
	defer stmnt.Close()

	_, err = stmnt.Exec()
	if err != nil {
		log.Fatal("error executing jobdb initialize statememtn", err)
	}
	return JOBDB
}

// Adds a new job to the database
// Returns the jobid of the new job
// Returns error if job already exists or could not be created
// Assumes JOBDB is initialized first by calling InnitializeJobdb()
func InsertNewJob(db *sql.DB, model string, lock_seed int, seed int, guidance int, upscale int, prompt string, owner string, job_params interface{}) (int, error) {

	job_params_unm, err := json.Marshal(job_params)
	if err != nil {
		log.Println("Error marshalling job_params", err)
		return -1, err
	}
	job_params_str := string(job_params_unm)

	stmnt, err := db.Prepare(`
	INSERT INTO "jobs" (
	model_pipeline,
	lock_seed,
	seed,
	guidance,
	upscale,
	prompt,
	status,
	job_params,
	owner,
	iteration_status,
	iteration_max,
	time_created,
	time_last_updated,
	time_completed
	) values (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ? ,? ,?);`)
	if err != nil {
		log.Println("error preparing job insert statement", err)
		return -1, err
	}
	defer stmnt.Close()

	time := strconv.Itoa(int(time.Now().Nanosecond()))
	iteration_max := 250

	result, err := stmnt.Exec(model, lock_seed, seed, guidance, upscale, prompt, "queued", job_params_str, owner, 0, iteration_max, time, time, "")
	if err != nil {
		log.Println("error executing job insert statement", err)
		return -1, err
	}

	numberOfNewJob, err := result.LastInsertId()
	if err != nil {
		log.Println("error getting result from LastInsertId", err)
		return -1, err
	}

	return int(numberOfNewJob), nil
}

// Get job by jobid
// Returns the job with the given jobid
func GetJobByJobid(db *sql.DB, jobid int) (Job, error) {
	var j Job

	row, err := db.Query(`SELECT * FROM "jobs" WHERE jobid = ?;`, jobid) // Query the database
	if err != nil {
		log.Println("error querying jobdb while getting job by id", err)
		return j, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&j.Jobid, &j.Model_pipeline, &j.Lock_seed, &j.Seed, &j.Guidance, &j.Upscale, &j.Prompt, &j.Status, &j.Job_params, &j.Owner, &j.Iteration_status, &j.Iteration_max, &j.Time_created, &j.Time_last_updated, &j.Time_completed)
	if err != nil {
		log.Println("error scaning jobdb row while getting job by id", err)
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
		log.Println("error querying jobdb while getting all jobs", err)
		return jobs, err
	}
	defer rows.Close()

	// Iterate over the rows and add them to the slice
	var j Job

	for rows.Next() {
		err = rows.Scan(&j.Jobid, &j.Model_pipeline, &j.Lock_seed, &j.Seed, &j.Guidance, &j.Upscale, &j.Prompt, &j.Status, &j.Job_params, &j.Owner, &j.Iteration_status, &j.Iteration_max, &j.Time_created, &j.Time_last_updated, &j.Time_completed)
		if err != nil {
			log.Println("error scanning rows from jobdb while getting all jobs", err)
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
		log.Println("error querying jobdb while getting jobs between X/Y", err)
		return nil, err
	}
	defer rows.Close()

	// Unmarshal the results into a slice of Jobs
	var j Job
	var jobs []Job
	for rows.Next() {
		err = rows.Scan(
			&j.Jobid,
			&j.Model_pipeline,
			&j.Lock_seed,
			&j.Seed,
			&j.Guidance,
			&j.Upscale,
			&j.Prompt,
			&j.Status,
			&j.Job_params,
			&j.Owner,
			&j.Iteration_status,
			&j.Iteration_max,
			&j.Time_created,
			&j.Time_last_updated,
			&j.Time_completed,
		)

		if err != nil {
			log.Println("error scanning rows while getting jobs between x/y", err)
			return nil, err
		}
		jobs = append(jobs, j)

	}

	return jobs, nil
}

// Returns oldest queued jon from the database
// Specifically, returns the job with the lowest jobid that has status "queued"
func GetOldestQueuedJob(db *sql.DB) (Job, error) {
	var j Job

	rows, err := db.Query(`SELECT * FROM "jobs" WHERE status = "queued" ORDER BY jobid ASC LIMIT 1;`) // Query the database
	if err != nil {
		log.Println("error getting row for oldest queued job", err)
		return j, err
	}
	defer rows.Close()

	rows.Next()
	err = rows.Scan(
		&j.Jobid,
		&j.Model_pipeline,
		&j.Lock_seed,
		&j.Seed,
		&j.Guidance,
		&j.Upscale,
		&j.Prompt,
		&j.Status,
		&j.Job_params,
		&j.Owner,
		&j.Iteration_status,
		&j.Iteration_max,
		&j.Time_created,
		&j.Time_last_updated,
		&j.Time_completed,
	)
	if err != nil {
		log.Println("error scanning rows while getting oldest queued job", err)
		return j, err
	}

	return j, err
}

// Returns number of jobs from the <db> that have the <status>
func GetNumberOfJobsThatHaveStatus(db *sql.DB, status string) int {
	var jobs []Job
	var j Job

	row, err := db.Query(`SELECT * FROM "jobs" WHERE status = ?;`, status)
	if err != nil {
		log.Println("Error in GetNumberOfCompletedJobs", err)
		return -1
	}
	defer row.Close()

	for row.Next() {
		err = row.Scan(
			&j.Jobid,
			&j.Model_pipeline,
			&j.Lock_seed,
			&j.Seed,
			&j.Guidance,
			&j.Upscale,
			&j.Prompt,
			&j.Status,
			&j.Job_params,
			&j.Owner,
			&j.Iteration_status,
			&j.Iteration_max,
			&j.Time_created,
			&j.Time_last_updated,
			&j.Time_completed,
		)
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
	err = row.Scan(
		&j.Jobid,
		&j.Model_pipeline,
		&j.Lock_seed,
		&j.Seed,
		&j.Guidance,
		&j.Upscale,
		&j.Prompt,
		&j.Status,
		&j.Job_params,
		&j.Owner,
		&j.Iteration_status,
		&j.Iteration_max,
		&j.Time_created,
		&j.Time_last_updated,
		&j.Time_completed,
	)
	if err != nil {
		log.Println("Error in GetNumberOfCompletedJobs", err)
		return "error scanning last job"
	}

	return j.Jobid
}

// returns only the newest job
func GetNewestCompletedJob(db *sql.DB, status string) Job {
	var newestCompletedJob Job

	rows, err := db.Query(`SELECT * FROM "jobs" WHERE status = ? ORDER BY jobid DESC LIMIT 1;`, status)
	if err != nil {
		log.Println("Error in GetNewestCompletedJob", err)
		return newestCompletedJob
	}
	defer rows.Close()
	println("B")

	newestCompletedJob, err = scan_single_row_of_jobs(rows, newestCompletedJob)
	if err != nil {
		log.Println("Error in GetNewestCompletedJob scan operation", err)
		return newestCompletedJob
	}
	println("C")
	return newestCompletedJob
}

func scan_single_row_of_jobs(r *sql.Rows, j Job) (Job, error) {
	r.Next()
	err := r.Scan(
		&j.Jobid,
		&j.Model_pipeline,
		&j.Lock_seed,
		&j.Seed,
		&j.Guidance,
		&j.Upscale,
		&j.Prompt,
		&j.Status,
		&j.Job_params,
		&j.Owner,
		&j.Iteration_status,
		&j.Iteration_max,
		&j.Time_created,
		&j.Time_last_updated,
		&j.Time_completed,
	)
	if err != nil {
		log.Println("error while scanning single row of jobs", err)
		return j, err
	}

	return j, nil
}

func scan_all_job_rows_into_a_job_slice(rows *sql.Rows, jobs []Job) ([]Job, error) {
	var j Job
	for rows.Next() {
		err := rows.Scan(
			&j.Jobid,
			&j.Model_pipeline,
			&j.Lock_seed,
			&j.Seed,
			&j.Guidance,
			&j.Upscale,
			&j.Prompt,
			&j.Status,
			&j.Job_params,
			&j.Owner,
			&j.Iteration_status,
			&j.Iteration_max,
			&j.Time_created,
			&j.Time_last_updated,
			&j.Time_completed,
		)
		if err != nil {
			log.Println("error while scanning job rows into a job slice", err)
			return jobs, err
		}

		jobs = append(jobs, j)
	}

	return jobs, nil
}

// Returns the last couple jobs from the databasee that have the status you ask
// status = the status you are filtering for
// numberOfJobs = the amount of jobs the function will return
func GetNewestCoupleJobsThatHaveStatus(db *sql.DB, status string, numberOfJobs int) []string {
	row, err := db.Query(`SELECT * FROM "jobs" WHERE status = ? ORDER BY jobid DESC LIMIT ?;`, status, numberOfJobs) // Query the database
	if err != nil {
		log.Println("Error in GetNewestFiveCompletedJobs", err)
	}
	defer row.Close()

	var jobs []Job
	jobs, err = scan_all_job_rows_into_a_job_slice(row, jobs)
	if err != nil {
		log.Println("error in get newest couple jobs that have status", err)
	}

	var answer []string
	for _, job := range jobs {
		answer = append(answer, job.Jobid)
	}

	return answer
}

/* to be continued */
// Updates the job in jobdb with the new status (completed or prrocessing) as well
// as the correct iteration status ( n/250)
func UpdateJobById(db *sql.DB, jobid string, newStatus string, iterStatus string) {

	iterStatusNumber, err := strconv.Atoi(iterStatus)
	if err != nil {
		log.Println("Error converting iterStatus string to int", err)
	}

	jobidNumber, err := strconv.Atoi(jobid)
	if err != nil {
		log.Println("Error converting jobid string to int", err)
	}

	println("UPDATING JOBDB WITH:", newStatus, iterStatus, "for jobid = ", jobid)

	// Iteration status is an int
	stmnt, err := db.Prepare(`
		UPDATE jobs
		SET
			status = ?,
			iteration_status = ?
		WHERE
    	jobid = ?;`)
	if err != nil {
		log.Println("error preparing statement", err)
	}
	defer stmnt.Close()

	result, err := stmnt.Exec(newStatus, iterStatusNumber, jobidNumber)
	if err != nil {
		log.Println("error Executing statement", err)
	}
	println(result)
}

// Returns the number of queued jobs
func GetNumberOfQueuedJobs(db *sql.DB) int {
	var s string

	rows, err := db.Query("select COUNT(*) from jobs where status = \"queued\";")
	if err != nil {
		println("Error in GetNumberOfQueued jobs", err)
		return 999
	}
	defer rows.Close()

	rows.Next()
	rows.Scan(&s)

	num, err := strconv.Atoi(s)
	if err != nil {
		println("Error converting in GetNumberOfQueued jobs", err)
		return 999
	}

	return num
}

// Returns slice with all jobs that match status <status>
func GetAllJobsInQueue(db *sql.DB) ([]Job, error) {
	var jobs []Job

	rows, err := db.Query(`SELECT * FROM "jobs" where
	status = "queued" or status = "processing" ORDER BY jobid DESC;`) // Query the database
	if err != nil {
		log.Println("error getting all jobs in queue", err)
		return jobs, err
	}
	defer rows.Close()

	jobs, err = scan_all_job_rows_into_a_job_slice(rows, jobs)
	if err != nil {
		log.Println("error after trying to scan al jobs rows into new job slice", err)
		return jobs, err
	}

	return jobs, err
}
