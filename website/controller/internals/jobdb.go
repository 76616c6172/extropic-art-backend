package internals // Change this to internals later

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	// 3rd party packages
	_ "github.com/mattn/go-sqlite3"
)

// WIP manage jobdb database
// Initiates the jobs database using SQLite

var JOBDB *sql.DB // This pointer is shared within the module to do database operations

func JobdbInit() {
	var err error
	JOBDB, err = sql.Open("sqlite3", "../model/jobdb/jobs.db")
	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}
	//defer db.Close() // close db before returning from this function

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

// Add a new job to the database
// REturns the jobid of the new job
// Returns error if job already exists or could not be created
// Assumes JOBDB is initialized by JobdbInit()
func InsertNewJob(prompt string, job_params interface{}) (int, error) {

	// Convert job_params to a string
	job_params_unm, err := json.Marshal(job_params)
	if err != nil {
		return -1, err
	}
	job_params_str := string(job_params_unm)

	// Construct the statement
	/* for reference the jobs schema is:
	INSERT INTO "jobs" (jobid, prompt, status, job_params, iteration_status, iteration_max, time_created, time_last_updated, time_completed)
	VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
	*/
	/*
		// MAYBE I WE DON'T NEED THIS?
		jobStatement := `
		INSERT INTO "jobs" (jobid, prompt, status, job_params, iteration_status, iteration_max, time_created, time_last_updated, time_completed)
		VALUES (?,"
		`
		jobStatement += prompt + "\", \"" + job_params_str + "\", 0, 0, 0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, NULL);"
	*/

	// Prepare the stament
	stmnt, err := JOBDB.Prepare(` INSERT INTO "jobs" (prompt, status, job_params, iteration_status, iteration_max, time_created, time_last_updated, time_completed) values (?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		return -1, err
	}

	// for reference the jobs schema is:
	//INSERT INTO "jobs" (jobid, prompt, status, job_params, iteration_status, iteration_max, time_created, time_last_updated, time_completed)
	// Execute the statement
	result, err := stmnt.Exec(prompt, "queued", job_params_str, 0, 0, "CURRENT_TIMESTAMP", "CURRENT_TIMESTAMP", "")
	if err != nil {
		fmt.Println("3")
		return -1, err
	}

	numberOfNewJob, err := result.LastInsertId()
	if err != nil {
		return -1, err
	}

	return int(numberOfNewJob), nil
}

// Called by the main function so we can test the module
func EntryPointForTesting() {
	jobnumber, err := InsertNewJob("testing prompt", "")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(jobnumber)
}
