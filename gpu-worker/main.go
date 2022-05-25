package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const WORKER_PORT = ":8090"
const SCHEDULER_PORT = ":8091"
const SCHEDULER_IP = "127.0.0.1"                                 //localhost for testing
const SECRET = "kldsjfksdjfwefjeojfefjkksdjfdsfsd932849j92h2uhf" //TODO: Authenticate better

var IS_BUSY = false //set to true while the worker is busy

// This infotion is sent by the scheduler
// Then is used to run the model (disco.py python script)
type Job struct {
	Prompt     string `json:"prompt"`
	Job_params string `json:"job_params"` // TODO: should be a struct
	Secret     string `json:"secret"`     // TODO: Authenticate better
}

type JobRequest struct {
	Secret string `json:"secret"` // TODO: Authenticate better
}

// Answers jobs posted to /api/0/worker
// I think this happens asynchronously
func api_0_worker(w http.ResponseWriter, r *http.Request) {

	// define the struct for the response
	var jobRequest Job // holds the request from the client
	m := struct {
		Accepted bool `json:"accepted"`
	}{
		Accepted: true,
	}

	if !IS_BUSY {

		// Read the request
		jsonDecoder := json.NewDecoder(r.Body)
		err := jsonDecoder.Decode(&jobRequest)
		if err != nil {
			log.Println(err) // maybe handle this better
			return
		}

		//fmt.Println(jobRequest.Prompt) // print prompt for testing
		json.NewEncoder(w).Encode(m) // send back the json as a the response
		IS_BUSY = true
		runModel(jobRequest.Prompt)
	} else {
		m.Accepted = false
		json.NewEncoder(w).Encode(m) // send back the json as a the response
	}
}

// Runs the clip guided diffusion model
func runModel(prompt string) {

	// build the parameters to call the script with
	modelParameters := fmt.Sprintf("--text_prompts '{\"0\": [\"%s\"]}' --steps 240 --width_height '[1920, 1080]'", prompt)
	fmt.Println(modelParameters)

	// call the python script with the prompt as agrument
	cmd := exec.Command("./disco.py", modelParameters)
	cmd.Stdout = os.Stdout // debug print
	cmd.Stderr = os.Stderr //debug print

	err := cmd.Run() //wait for the command to finish
	if err != nil {
		log.Println(err) //TODO: notify the scheduler that something went wrong with this job
	}

	out, err := cmd.Output()
	if err != nil {
		log.Println(err) //TODO: notify the scheduler that something went wrong with this job
	}
	fmt.Println(out)

	fmt.Println("model run complete, setting worker to available") // DEBUG
	IS_BUSY = false
}

// Sends a job request to the scheduler
// returns error if it fails
// returns job if job was received
func sendJobRequest() (Job, error) {

	fmt.Println("Encoding jobrequest ") //debug

	var jr JobRequest
	jr.Secret = SECRET

	jsonReq, err := json.Marshal(jr)
	if err != nil {
		log.Println(err)
	}

	fmt.Println("Sending jobrequest") //debug

	schedulerAddress := "http://" + SCHEDULER_IP + SCHEDULER_PORT + "/api/0/scheduler"
	resp, err := http.Post(schedulerAddress, "application/json; charset=utf-8", bytes.NewBuffer(jsonReq))
	//resp, err := http.PostForm("http://"+SCHEDULER_IP+SCHEDULER_PORT+"/api/0/scheduler",)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	fmt.Println("Reading response") //debug

	bodyBytes, _ := ioutil.ReadAll(resp.Body)
	// Convert response body to string
	bodyString := string(bodyBytes)
	fmt.Println(bodyString)

	// Convert response body to Todo struct
	var jobReceivedFromScheduler Job
	json.Unmarshal(bodyBytes, &jobReceivedFromScheduler)

	return jobReceivedFromScheduler, err
}

func main() {

	logFile, err := os.OpenFile(("./logs/worker.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	http.HandleFunc("/api/0/worker", api_0_worker) //register handler for /api/0/worker
	go http.ListenAndServe(WORKER_PORT, nil)       //start the server

	// Do some kind of loop here to ask for jobs, and run them
	Job, err := sendJobRequest()
	if err != nil {
		log.Println("Error requesting job:", err)
	}
	fmt.Println(Job)
}
