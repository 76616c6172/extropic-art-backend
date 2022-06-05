package exdb

import (
	// 3rd party packages
	"database/sql"
	"fmt"
	"log"
	"time"

	//_ "github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

var WORKERDB *sql.DB // This pointer is shared within the module to do database operations

// Initializes and connects to workers.db
func InitializeWorkerdb() {

	var err error
	WORKERDB, err = sql.Open("sqlite3", "../model/workerdb/workers.db")
	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}
	//defer db.Close() // If we wanted to close it

	// worker_uid
	// worker_ip "127.0.0.1:9080"
	// worker_busy true/false
	// worker_gpu_type "v100"
	// worker_current_job "jobid"

	// Create table if it doesn't exist
	stmnt, err := WORKERDB.Prepare(`
CREATE TABLE IF NOT EXISTS "workers" (
 "worker_id" TEXT,
 "worker_ip" TEXT,
 "worker_busy" INTEGER,
 "worker_current_job", INTEGER,
 "worker_last_health_check" INTEGER,
 "worker_time_created" INTEGER,
 "worker_secret" TEXT,
 "worker_type" INT
	);`)

	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}
	_, err = stmnt.Exec()
	if err != nil {
		log.Fatal(err)
	}
}

// Add a new worker to the db
func RegisterNewWorker(workerId string) error {
	var err error

	stmnt, err := WORKERDB.Prepare(` INSERT INTO workers (worker_id, worker_ip, worker_busy, worker_current_job, worker_last_health_check, worker_time_created, worker_secret, worker_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	unixtime := int(time.Now().Unix())
	result, err := stmnt.Exec(workerId, "", 0, 0, 0, unixtime, "", 0) // Execute the statement
	if err != nil {
		return err
	}
	fmt.Println(result.LastInsertId())

	return err
}
