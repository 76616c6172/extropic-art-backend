package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"extropic-art-backend/src/exapi"
	"extropic-art-backend/src/exdb"
	"extropic-art-backend/src/exutils"
)

const MAX_RETRIES = 3
const INFERENCE_CMD = "./run_inference.py"

var MAX_PARALELL_EXECUTION = 8
var CURRENT_WORKERS = 0
var MJSD_INSTANCE_1_IS_IN_USE = false
var MJSD_INSTANCE_2_IS_IN_USE = false
var WORKER_DB *sql.DB
var JOB_COMPLETE = true
var JOBDB *sql.DB
var SECRET string
var PANIC bool

// Keeps the scheduler running until CTRL-C or exit signal is received.
func KeepSchedulerRunningUntilExitSignalIsReceived() {
	fmt.Println("Scheduler is running..") // debug
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-channel
	fmt.Println("scheduler closed gracefully")
}

// Initalizes/creates the log file as needed.
func InitializeLogFile() {
	logFile, err := os.OpenFile(("./logs/scheduler.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	log.SetOutput(logFile)
}

func runJob(j exdb.Job) error {
	var PROMPT, SEED, WIDTH, HEIGHT, STEPS, SCALE string
	var timesRetried int

	CURRENT_WORKERS++

	WIDTH = "512"
	HEIGHT = "512"
	if j.Job_params == "2" {
		WIDTH = "512"
		HEIGHT = "768"
	} else if j.Job_params == "3" {
		WIDTH = "768"
		HEIGHT = "512"
	} else if j.Job_params == "4" {
		WIDTH = "1024"
		HEIGHT = "512"
	} else if j.Job_params == "5" {
		WIDTH = "512"
		HEIGHT = "1024"
	} else if j.Job_params == exapi.RES_768_BY_768 {
		WIDTH = "768"
		HEIGHT = "768"
	} else if j.Job_params == exapi.RES_768_BY_1024 {
		WIDTH = "768"
		HEIGHT = "1024"
	} else if j.Job_params == exapi.RES_1024_BY_768 {
		WIDTH = "1024"
		HEIGHT = "768"
	}
	SEED = strconv.Itoa(j.Seed)
	SCALE = "7"
	PROMPT = j.Prompt

	// FIXME: steps and scale should be collumns in jobdb and directly exposed to the user
	MODEL := j.Model_pipeline
	if MODEL == "1" {
		STEPS = "75"
	} else if MODEL == "2" {
		STEPS = "30"
	}

	exdb.UpdateJobById(JOBDB, j.Jobid, "processing", "1")

	cmd := exec.Command(INFERENCE_CMD, MODEL, PROMPT, SEED, WIDTH, HEIGHT, STEPS, SCALE, j.Jobid)
	out, err := cmd.Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println(string(out))

	CURRENT_WORKERS--

	// Create JPG "thumbnail", TODO: Pull this into the gi runtime and just use the standard library for this
	if err = exec.Command("../model/make_jpgs_from_name", j.Jobid).Run(); err != nil {
		if timesRetried < MAX_RETRIES {
			timesRetried++
			exdb.UpdateJobById(JOBDB, j.Jobid, "j", "0")
			return err
		} else {
			timesRetried = 0
			exdb.UpdateJobById(JOBDB, j.Jobid, "failed", "0")
		}
	}

	// Update the jobdb, TODO: Add safety measure where jobs already marked completed can't be updated by the worker
	exdb.UpdateJobById(JOBDB, j.Jobid, "completed", "250")
	return nil
}

// runSchedulingLoop looks for "queued" jobs and schedules them to serverless workers
func runSchedulingLoop(quit chan bool) {
	for {
		select {
		case <-quit:
			println("scheduler exiting")
			return

		default:
			time.Sleep(1 * time.Second)
			println("checking jobdb for oldest queued job")
			queuedJob, err := exdb.GetOldestQueuedJob(JOBDB)
			if err != nil {
				println("..waiting for queued job")
				continue
			}
			println("Got queued job:", queuedJob.Jobid, queuedJob.Prompt)
			if CURRENT_WORKERS < MAX_PARALELL_EXECUTION {
				go runJob(queuedJob)
			}
			println("..waiting for free worker instance")
			continue
		}
	}
}

// This is the main function >:D
func main() {
	InitializeLogFile()
	SECRET = exutils.InitializeSecretFromArgument()
	JOBDB = exdb.InitializeJobdb()
	WORKER_DB = exdb.InitializeWorkerdb()

	quit := make(chan bool)
	go runSchedulingLoop(quit)

	KeepSchedulerRunningUntilExitSignalIsReceived()
	close(quit)
}
