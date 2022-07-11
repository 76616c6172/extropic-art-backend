package exdb

import (
	// 3rd party packages
	"database/sql"
	"log"
	"strconv"
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
	"worker_id" TEXT PRIMARY KEY,
	"worker_ip" TEXT,
	"worker_busy" INTEGER,
	"worker_current_job" INTEGER,
	"worker_last_health_check" INTEGER,
	"worker_time_created" INTEGER,
	"worker_secret" TEXT,
	"worker_type" INTEGER
	) WITHOUT ROWID;`)
	if err != nil {
		log.Fatal("InitializeWorkerdb: Error preparing SQLite statement", err)
	}
	defer stmnt.Close()

	_, err = stmnt.Exec()
	if err != nil {
		log.Fatal("InitializeWorkerdb: Error executing SQLite statment", err)
	}

	return db
}

// Add a new worker to the db
func RegisterNewWorker(db *sql.DB, workerId string, ipAddress string, workerType int) error {

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
		log.Println("ERROR PREPARING STATEMENT", err)
		return err
	}
	defer stmnt.Close()

	unixtime := int(time.Now().Unix())
	curJobAssignedToWorker := -1
	result, err := stmnt.Exec(workerId, ipAddress, NOT_BUSY, curJobAssignedToWorker, unixtime, unixtime, "none", workerType) // Execute the statement
	if err != nil {
		log.Println("ERROR EXECUTING STATEMENT", err)
		return err
	}

	insertID, err := result.LastInsertId()
	if err != nil {
		log.Println("ERROR READING WORKERDB RESULT", err)
		return err
	}

	log.Println("successfully registered new worker", insertID)

	return err
}

func GetWorkerByIP(db *sql.DB, workerIp string) (Worker, error) {
	var wk Worker

	println(workerIp)
	row, err := db.Query(`SELECT * FROM "workers" WHERE worker_ip = ?;`, workerIp) // Query the database
	if err != nil {
		log.Println("Error getting worker by IP", err)
		return wk, err
	}
	defer row.Close()

	rowExists := row.Next()
	if rowExists {
		err = row.Scan(&wk.Worker_id, &wk.Worker_ip, &wk.Worker_Busy, &wk.Worker_current_job, &wk.Worker_last_health_check, &wk.Worker_time_created, &wk.Worker_secret, &wk.Worker_type)
		if err != nil {
			log.Println("Error reading worker row", err)
			return wk, err
		}
	}

	return wk, err
}

// Updates the worker in the database that matches the id with provided parameters
func UpdateWorkerByWorkerId(db *sql.DB, currentJob string, workerId string, isBusy int) {

	currentJobNumber, err := strconv.Atoi(currentJob)
	if err != nil {
		log.Println("error converting currentJob to int", err)
		return
	}

	stmnt, err := db.Prepare(`
		UPDATE workers
		SET
			worker_busy = ?,
			worker_current_job = ?
		WHERE
    	 worker_id = ?;`)
	if err != nil {
		log.Println("error preparing statement", err)
		return
	}
	defer stmnt.Close()

	// if job is done override the job assigned to the worker in the db with -1
	result, err := stmnt.Exec(isBusy, currentJobNumber, workerId)
	if err != nil {
		log.Println("error executing statement", err)
	}

	println(result)
}

// Updates a worker in the db by the job it was assigned to
func UpdateWorkerByJobid(db *sql.DB, currentJob string, jobCompleted bool) {
	var isBusy int

	currentJobNumber, err := strconv.Atoi(currentJob)
	if err != nil {
		log.Println("error converting currentJob to int", err)
		return
	}

	stmnt, err := db.Prepare(`
		UPDATE workers
		SET
			worker_busy = ?,
			worker_current_job = ?
		WHERE
    	 worker_current_job = ?;`)
	if err != nil {
		log.Println("error preparing statement", err)
		return
	}
	defer stmnt.Close()

	// if job is done override the job assigned to the worker in the db with -1
	if jobCompleted {
		isBusy = 0
		result, err := stmnt.Exec(isBusy, -1, currentJobNumber)
		if err != nil {
			log.Println("error executing statement", err)
		}
		println(result)
		return
	}
	isBusy = 1
	result, err := stmnt.Exec(isBusy, currentJobNumber, currentJobNumber)
	if err != nil {
		log.Println("error executing statement", err)
	}

	println(result)
}

// Fetches and returns a free GPU-worker from the workerdb
func GetFreeWorker(db *sql.DB) (Worker, error) {
	var w Worker

	row, err := db.Query(`SELECT * FROM "workers" WHERE worker_busy = 0 ORDER BY worker_id ASC LIMIT 1;`) // Query the database
	if err != nil {
		return w, err
	}
	defer row.Close()
	/*
	   "worker_id" TEXT PRIMARY KEY,
	   "worker_ip" TEXT,
	   "worker_busy" INTEGER,
	   "worker_current_job" INTEGER,
	   "worker_last_health_check" INTEGER,
	   "worker_time_created" INTEGER,
	   "worker_secret" TEXT,
	   "worker_type" INTEGER
	*/

	row.Next()
	err = row.Scan(&w.Worker_id, &w.Worker_ip, &w.Worker_Busy, &w.Worker_current_job, &w.Worker_last_health_check, &w.Worker_time_created, &w.Worker_secret, &w.Worker_type)
	if err != nil {
		return w, err
	}

	return w, err
}
