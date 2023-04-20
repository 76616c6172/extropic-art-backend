package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"
	"time"

	"extropic-art-backend/src/exapi"
	"extropic-art-backend/src/exdb"
)

const PORT_TO_SERVE_ON = ":8080"

var JOB_DB *sql.DB

// initializeLogFile as controller.log
func initializeLogFile() {
	logFile, err := os.OpenFile(("./controller.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	log.SetOutput(logFile)
}

// go resetDailyUsageLimit() resets the usage period every 24h
func resetDailyUsageLimit() {
	for {
		<-time.After(24 * time.Hour)
		exapi.Mutex.Lock()
		exapi.FREE_USES_REMAINING = exapi.MAXIMUM_DAILY_USES
		exapi.Mutex.Unlock()
	}
}

// This is the main function :D
func main() {
	initializeLogFile()
	JOB_DB = exdb.InitializeJobdb()

	// Website ressources are served from filepath ../view/build
	http.Handle("/", http.FileServer(http.Dir("../view/build")))

	// Old API endpoints
	http.HandleFunc("/api/0/status", api_0_status)
	http.HandleFunc("/api/0/jobs", api_0_jobs)
	http.HandleFunc("/api/0/img", api_0_img)

	// New API endpoints
	http.HandleFunc("/api/1/queue", api_1_queue)
	http.HandleFunc("/api/1/status", api_1_status_handler)
	http.HandleFunc("/api/1/jobs", api_1_jobs)

	go resetDailyUsageLimit()

	http.ListenAndServe(PORT_TO_SERVE_ON, nil)
}
