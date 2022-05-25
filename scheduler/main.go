package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

// Schema of a Job for the gpu-worker
type Job struct {
	Prompt     string `json:"prompt"`
	Job_params string `json:"job_params"` // TODO: should be a struct
	Secret     string `json:"secret"`     // TODO: Authenticate better
}

const WEBSERVER_PORT = ":8091" // Scheduler is listening on this port
const SECRET = "kldsjfksdjfwefjeojfefjkksdjfdsfsd932849j92h2uhf"

type jobRequest struct {
	Secret string `json:"secret"` // TODO: Authenticate better
}

// Scheduler listens on /api/0/scheduler for jobrequests
// Sends back jobs to gpu-worker on request and writes to db
func api_0_scheduler(w http.ResponseWriter, r *http.Request) {

	var jobResponse Job //  TODO Get this Job from the db instead
	jobResponse.Prompt = "Panorama photo of an intricate wizzards tower on mount everest, dark elf gothic arthitecture, trending on arstation, 3d photorealistic materials,"
	jobResponse.Job_params = ""
	jobResponse.Secret = SECRET

	// define the struct for the response
	var jobRequest jobRequest              // holds the request from the client
	jsonDecoder := json.NewDecoder(r.Body) // Read the request
	err := jsonDecoder.Decode(&jobRequest) // Store the request in jobRequest
	if err != nil {
		log.Println(err) // maybe handle this better
		return
	}

	if jobRequest.Secret == SECRET {

		json.NewEncoder(w).Encode(jobResponse) // send back the response
	}

	fmt.Println("Sent Job: ", jobResponse.Prompt)
}

func main() {
	fmt.Println("Yo")

	logFile, err := os.OpenFile(("./logs/scheduler.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	http.HandleFunc("/api/0/scheduler", api_0_scheduler) //register handler for /api/0/scheduler
	go http.ListenAndServe(WEBSERVER_PORT, nil)          //start the server

	fmt.Println("Scheduler is running..") // keep the scheduler running until CTRL-C
	channel := make(chan os.Signal, 1)
	signal.Notify(channel, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-channel
	fmt.Println("scheduler closed gracefully")

	// Run some kind of loop here to check if jobs have been processing for a long time
	// if so set them back to queued.
}
