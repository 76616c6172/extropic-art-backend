package internals // Change this to internals later

import (
	"database/sql"
	"log"

	// 3rd party packages
	_ "github.com/mattn/go-sqlite3"
)

// WIP manage jobdb database
func Testing() {

	db, err := sql.Open("sqlite3", "../model/jobdb/jobs.db")
	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}
	defer db.Close() // close db before returning from this function

	// Create table if it doesn't exist
	stmnt, err := db.Prepare(`
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
