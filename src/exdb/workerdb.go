package exdb

import (
	// 3rd party packages
	"database/sql"
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

	unixtime := int(time.Now().Unix())
	result, err := stmnt.Exec(workerId, ipAddress, 0, 0, 0, unixtime, "", workerType) // Execute the statement
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

	row, err := db.Query(`SELECT * FROM "workers" WHERE worker_ip = ?;`, workerIp) // Query the database
	if err != nil {
		log.Println("Error getting worker by IP", err)
		return wk, err
	}
	defer row.Close()

	row.Next()
	err = row.Scan(&wk.Worker_id, &wk.Worker_ip, &wk.Worker_Busy, &wk.Worker_current_job, &wk.Worker_last_health_check, &wk.Worker_time_created, &wk.Worker_secret, &wk.Worker_type)
	if err != nil {
		log.Println("Error reading worker row", err)
		return wk, err
	}

	return wk, err
}
