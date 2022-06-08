package exdb

import (
	// 3rd party packages
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

// Initialize and connect to workerdb
func InitializeWorkerdb() *sql.DB {
	var err error
	db, err := sql.Open("sqlite3", "../model/workerdb/workers.db")
	if err != nil {
		log.Fatal(err) // TODO: Maybe handle this better
	}

	// Create table if it doesn't exist
	stmnt, err := db.Prepare(`
	CREATE TABLE IF NOT EXISTS "workers" (
	"worker_id" TEXT UNIQUE,
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

	// testing
	err = db.Ping()
	println("WE GOT HERE: WE CAN PING WORKERDB FINE FROM INITWORKERDB")
	println(err)

	return db
}

// Add a new worker to the db
func RegisterNewWorker(db *sql.DB, workerId string, ipAddress string, workerType int) error {

	println("WE GOT HERE: 0")

	//THEORY: Since we imported sqlite3 in a weird way, I am guessing that GO
	// Garbage Collects the pointers at some point
	err := db.Ping()
	println("WE GOT HERE: 1")
	println(err)

	/*
		stmnt, err := JOBDB.Prepare(`INSERT INTO "jobs" (
			prompt,
		  status,
			job_params,
			 iteration_status,
			 iteration_max,
			 time_created,
			 time_last_updated,
			 time_completed) values (?, ?, ?, ?, ?, ?, ?, ?);`)
	*/

	println("WE GOT HERE: 2")

	stmnt, err := db.Prepare(`INSERT INTO "workers" (
			worker_id,
			worker_ip,
			worker_busy,
			worker_current_job,
			worker_last_health_check,
			worker_time_created,
			worker_secret,
			worker_type) values (?, ?, ?, ?, ?, ?, ?, ?);`)
	if err != nil {
		println("ERROR PREPARING STATEMENT", err)
		return err
	}

	println("WE GOT HERE: 3")

	unixtime := int(time.Now().Unix())
	result, err := stmnt.Exec(workerId, ipAddress, 0, 0, 0, unixtime, "", workerType) // Execute the statement
	if err != nil {
		return err
	}
	fmt.Println(result.LastInsertId())

	println("WE GOT HERE: 4")

	return err
}
