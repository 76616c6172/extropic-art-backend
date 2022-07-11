package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"net/http"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"project-exia-monorepo/website/exapi"
	"project-exia-monorepo/website/exdb"
	"project-exia-monorepo/website/exutils"
)

const IMAGE_PATH = "./images_out/TimeToDisco/progress.png"
const WORKER_PORT = ":8090"
const IS_COLAB_WORKER = false //set to false if not colab worker
//const SCHEDULER_IP = "http://exia.art:8091"
const SCHEDULER_IP = "http://127.0.0.1:8091"
const JOB_IS_DONE bool = true
const JOB_IS_NOT_DONE bool = false

var HAVE_JOB = false //set to true while the worker is busy
var WORKER_ID string
var SECRET string
var JOB_PROMPT string // prompt for the current job to be rendered

// Answers jobs posted to /api/0/worker
func handleNewJobPosting(w http.ResponseWriter, r *http.Request) {

	println("Receiving new posting")
	//println(r.Body)

	/*
		var jobRequest exdb.Job    // holds the request from the client
		var m exapi.WorkerResponse // Response for the scheduler
	*/

	if !HAVE_JOB {
		var jobRequest exdb.Job

		HAVE_JOB = true
		// Read the request
		jsonDecoder := json.NewDecoder(r.Body)
		err := jsonDecoder.Decode(&jobRequest)
		if err != nil {
			log.Println(err) // maybe handle this better
			return
		}

		println("received job:")
		println(jobRequest.Jobid, jobRequest.Prompt)

		HAVE_JOB = true
		JOB_PROMPT = jobRequest.Prompt

		// Set headers
		w.Header().Set(exapi.HeaderJobAccepted, "1")

		if IS_COLAB_WORKER {
			w.Header().Set(exapi.HeaderColabWorker, "1")
		}

		w.WriteHeader(http.StatusAccepted)

		defer r.Body.Close() // Close the response
		return

	} else {
		w.WriteHeader(http.StatusForbidden)
		// Send response to the scheduler
		//json.NewEncoder(w).Encode(m)
	}
}

// Runs the clip guided diffusion model
// FIXME: This function is ugly!
func runModel(prompt string) {

	// 0. TODO: Write the prompt and params to yaml config file

	// 1. Run the model
	//modelParameters := fmt.Sprintf("--text_prompts '{\"0\": [\"%s\"]}' --steps 240 --width_height '[1920, 1088]'", prompt)
	// --n_batches 0 --text_prompts '{"0": ["Space panorama of celestial space station, cosmic, photorealistic materials, beeple behance"]}' --steps 250 --width_height '[1920, 1088]'

	args := fmt.Sprintf("--n_batches 0 --text_prompts '{\"0\": [\"")
	args += prompt
	args += "\"]}' --steps 250 --width_height '[1920, 1088]'"

	modelSubProcess := exec.Command("./run_model", args)

	stdout, err := modelSubProcess.StdoutPipe()
	if err != nil {
		log.Println("testing: error connecting to stdout", err)
		return
	}
	defer stdout.Close()

	stderr, err := modelSubProcess.StderrPipe()
	if err != nil {
		log.Println("testing: error connecting to stderr", err)
		return
	}
	defer stderr.Close()

	// Run the model
	err = modelSubProcess.Start()
	if err != nil {
		log.Println("testing: error running cmd.Start()", err)
		return
	}

	// Read the output of the model line by line
	stdoutScanner := bufio.NewScanner(stdout)
	//stderrScanner := bufio.NewScanner(stderr)

	numberOfTimesInProgressPngWasCreated := -1
	var exiting bool
	for {

		if exiting {
			if modelSubProcess.ProcessState.Exited() {
				fmt.Println("Exiting the model gracefully")
				break
			}
		}

		isReceivingAnotherLineFromStdout := stdoutScanner.Scan()

		if isReceivingAnotherLineFromStdout {
			fmt.Println("1")
			currentLineFromStdout := stdoutScanner.Text()

			if strings.Contains(currentLineFromStdout, "<PIL.Image.Image") {
				// Another in progress.png was created by the model!
				numberOfTimesInProgressPngWasCreated++
				iterationStatus := numberOfTimesInProgressPngWasCreated * 50

				err = postJobdUpdateToScheduler(strconv.Itoa(iterationStatus), JOB_IS_NOT_DONE)
				if err != nil {
					println("error after posting update to scheduler: ", err)
				}

			} else if strings.Contains(currentLineFromStdout, "SUCCESS") {
				fmt.Println("SUCCESS, model run complete")
				err = postJobdUpdateToScheduler("250", JOB_IS_DONE)
				if err != nil {
					println("error sending final update to scheduler", err)
					log.Println("error sending final update to scheduler", err)
					break
				}
				exiting = true
				modelSubProcess.Wait()
			}

		}
	}
	HAVE_JOB = false
}

// Sends image + metadata to the scheduler
func postJobdUpdateToScheduler(iteration_status string, jobIsDone bool) error {

	println("Posting job update to scheduler..")

	client := &http.Client{
		Timeout: time.Second * 10,
	}

	// New multipart writer.
	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	fw, err := writer.CreateFormFile("image", "test.png")
	if err != nil {
		log.Println(err)
		return err
	}

	file, err := os.Open("images_out/TimeToDisco/progress.png")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = io.Copy(fw, file)
	if err != nil {
		log.Println(err)
		return err
	}
	writer.Close()
	req, err := http.NewRequest("POST", SCHEDULER_IP+"/api/0/report", bytes.NewReader(body.Bytes()))
	if err != nil {
		log.Println(err)
		return err
	}

	if jobIsDone {
		req.Header.Add(exapi.HeaderJobStatusComplete, "1")
	} else {
		req.Header.Add(exapi.HeaderJobStatusComplete, "0")
	}

	req.Header.Add(exapi.HeaderJobIterationStatus, iteration_status)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	rsp, _ := client.Do(req)
	println(rsp.StatusCode)

	/*
		if rsp.ContentLength > 0 {

			if rsp.StatusCode != http.StatusOK {
				log.Printf("Request failed with response code: %d\n", rsp.StatusCode)
				println("Request failed with response code: ", rsp.StatusCode)
				return err
			}
			println("successfully posted update to scheduler")
		}
	*/
	return nil
}

// Send an authenticated webrequest to the scheduler, registering the worker
func registerWorkerWithScheduler() {
	fmt.Println("Registering with scheduler") // debug

	// Prepare the web request
	req, err := http.NewRequest("POST", SCHEDULER_IP+"/api/0/registration", nil)
	if err != nil {
		log.Println("Error registering worker:", err)
		return
	}
	req.AddCookie(&http.Cookie{Name: "secret", Value: SECRET})

	if IS_COLAB_WORKER {
		req.Header.Add(exapi.HeaderColabWorker, "1")
	}

	// Send the web request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("Error sending registration: ", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		log.Println("Error registering worker - ", resp.StatusCode)
		return
	}

	fmt.Println("Worker successully registered with scheduler")
}

// Initializes log file for the GPU_WORKER
func initializeLogFile() {
	logFile, err := os.OpenFile(("./logs/worker.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	log.SetOutput(logFile)
}

// If HAVE_JOB is true, run the model against the JOB_PROMPT
// if HAVE_JOB is false, keep checking every 5 seconds if we have a job
func runWorkerLoop() {
	for {

		if HAVE_JOB {
			println("Running job:", JOB_PROMPT)
			runModel(JOB_PROMPT)
			println("Completed Job", JOB_PROMPT)
		}

		HAVE_JOB = false
		time.Sleep(5 * time.Second)
	}
}

// This is the main function :D
func main() {
	initializeLogFile()
	SECRET = exutils.InitializeSecretFromArgument()

	registerWorkerWithScheduler()
	http.HandleFunc("/api/0/worker", handleNewJobPosting) // Listen for new jobs on this endpoint
	//postJobdUpdateToScheduler("250", JOB_IS_DONE)

	fmt.Println("Worker is running, waiting for job posts..") // Debug
	go http.ListenAndServe(WORKER_PORT, nil)

	// Run worker loop
	runWorkerLoop()
}
