package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
)

const WEBSERVER_PORT = ":8090"

var IS_BUSY = false //set to true while the worker is busy

// This infotion is sent by the scheduler
// Then is used to run the model (disco.py python script)
type job struct {
	Prompt     string `json:"prompt"`
	Job_params string `json:"job_params"` // TODO: should be a struct
}

// Answers jobs posted to /api/0/worker
// I think this happens asynchronously
func api_0_worker(w http.ResponseWriter, r *http.Request) {

	// define the struct for the response
	var jobRequest job // holds the request from the client
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

func main() {
	fmt.Println("Yo")

	// Select the logfile
	logFile, err := os.OpenFile(("./logs/worker.log"), os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("main: error opening logfile")
	}
	defer logFile.Close()
	log.SetOutput(logFile)

	http.HandleFunc("/api/0/worker", api_0_worker) //register handler for /api/0/worker
	http.ListenAndServe(WEBSERVER_PORT, nil)       //start the server
}
