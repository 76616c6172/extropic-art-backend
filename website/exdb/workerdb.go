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

var WORKERDB *sql.DB // This pointer is used to do interact with the database through its methods

// Initializes and connects to workers.db
func InitializeWorkerdb() {

	WORKERDB, err := sql.Open("sqlite3", "../model/workerdb/workers.db")
	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}

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
		log.Fatal("InitializeWorkerdb: Error preparing SQLite statement", err)
	}

	_, err = stmnt.Exec()
	if err != nil {
		log.Fatal("InitializeWorkerdb: Error executing SQLite statment", err)
	}
}

// Add a new worker to the db
func RegisterNewWorker(workerId string, ipAddress string, workerType int) error {

	stmnt, err := WORKERDB.Prepare(`INSERT INTO workers (worker_id, worker_ip, worker_busy, worker_current_job, worker_last_health_check, worker_time_created, worker_secret, worker_type) VALUES (?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return err
	}

	unixtime := int(time.Now().Unix())
	result, err := stmnt.Exec(workerId, ipAddress, 0, 0, 0, unixtime, "", workerType) // Execute the statement
	if err != nil {
		return err
	}
	fmt.Println(result.LastInsertId())

	return err
}
