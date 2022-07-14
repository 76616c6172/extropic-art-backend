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

//const SCHEDULER_IP = "http://127.0.0.1:8091"
const SCHEDULER_IP = "http://exia.art:8091"
const JOB_IS_DONE bool = true
const JOB_IS_NOT_DONE bool = false

var HAVE_JOB = false //set to true while the worker is busy
var WORKER_ID string
var SECRET string
var JOB_PROMPT string         // prompt for the current job to be rendered
var TUNNEL_URL string         // send to scheduler if we're using a tunnel (provided as 2nd startup argument)
var WORKER_HAS_TUNNEL = false // is set to false if not colab worker

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

		// Read the request
		jsonDecoder := json.NewDecoder(r.Body)
		err := jsonDecoder.Decode(&jobRequest)
		if err != nil {
			log.Println(err) // maybe handle this better
			return
		}

		println("received job:")
		println(jobRequest.Jobid, jobRequest.Prompt)

		// Set headers
		w.Header().Set(exapi.HeaderJobAccepted, "1")

		w.WriteHeader(http.StatusAccepted)

		defer r.Body.Close() // Close the response

		JOB_PROMPT = jobRequest.Prompt
		HAVE_JOB = true
		return

	} else {
		w.WriteHeader(http.StatusForbidden)
		// Send response to the scheduler
		//json.NewEncoder(w).Encode(m)
	}
}

func writeYamlConfiguration(prompt string) {
	filePath := "./configs/TimeToDisco.yaml"

	FileContent := fmt.Sprintln("RN101: false\nRN50: true\nRN50x16: false\nRN50x4: false\nRN50x64: false\nViTB16: true\nViTB32: true\nViTL14: false\nViTL14_336: false\nangle: 0:(0)\nanimation_mode: None\nbatch_name: TimeToDisco\ncheck_model_SHA: false\nclamp_grad: true\nclamp_max: 0.05\nclip_denoised: false\nclip_guidance_scale: 5000\nconsole_preview: false\nconsole_preview_width: 80\ncuda_device: cuda:0\ncut_ic_pow: 1\ncut_icgray_p: '[0.2]*400+[0]*600'\ncut_innercut: '[4]*400+[12]*600'\ncut_overview: '[12]*400+[4]*600'\ncutn_batches: 4\ncutout_debug: false\ndb: null\ndiffusion_model: 512x512_diffusion_uncond_finetune_008100\ndiffusion_sampling_mode: ddim\ndisplay_rate: 50\neta: 0.8\nextract_nth_frame: 2\nfar_plane: 10000\nfov: 40\nframes_scale: 1500\nframes_skip_steps: 60%\nfuzzy_prompt: false\nimage_prompts: {}\nimages_out: images_out\ninit_image: null\ninit_images: init_images\ninit_scale: 1000\nintermediate_saves: 0\nintermediates_in_subfolder: true\ninterp_spline: Linear\nkey_frames: true\nmax_frames: 10000\nmidas_depth_model: dpt_large\nmidas_weight: 0.3\nmodels: models2\nmodifiers: {}\nmultipliers: {}\nn_batches: 1\nnear_plane: 200\npadding_mode: border\nper_job_kills: false\nperlin_init: false\nperlin_mode: mixed\nrand_mag: 0.05\nrandomize_class: true\nrange_scale: 150\nresume_from_frame: latest\nresume_run: false\nretain_overwritten_frames: false\nrotation_3d_x: '0: (0)'\nrotation_3d_y: '0: (0)'\nrotation_3d_z: '0: (0)'\nrun_to_resume: latest\nsampling_mode: bicubic\nsat_scale: 0\nsave_metadata: false\nset_seed: random_seed\nsimple_nvidia_smi_display: true\nskip_augs: false\nskip_steps: 10\nskip_video_for_run_all: false\nsteps: 250\nsymmetry_loss: false\nsymmetry_loss_scale: 1500\nsymmetry_switch: 40")

	FileContent += fmt.Sprintf("text_prompts:\n 0:\n - %s\n", prompt)

	FileContent += fmt.Sprintln("translation_x: '0: (0)'\ntranslation_y: '0: (0)'\ntranslation_z: '0: (10.0)'\ntransport: '1'\nturbo_mode: false\nturbo_preroll: 10\nturbo_steps: 3\ntv_scale: 0\ntwilio_account_sid: null\ntwilio_auth_token: null\ntwilio_from: null\ntwilio_to: null\nuseCPU: false\nuse_checkpoint: true\nuse_secondary_model: true\nv_symmetry_loss: false\nv_symmetry_loss_scale: 1500\nv_symmetry_switch: 40\nvideo_init_path: training.mp4\nvideo_init_seed_continuity: true\nvr_eye_angle: 0.5\nvr_ipd: 5.0\nvr_mode: false\nwidth_height:\n- 1920\n- 1088\nzoom: '\"0: (1), 10: (1.05)'")

	contentBytes := []byte(FileContent)

	err := os.WriteFile(filePath, contentBytes, 0666)
	if err != nil {
		println("Error writing model yaml config", err)
		log.Println("Error writing model yaml config", err)
	}
}

// Runs the clip guided diffusion model
// FIXME: This function is ugly!
func runModel(prompt string) {

	// 1. Write Model Configuration in the apropriate file
	writeYamlConfiguration(prompt)

	// 2. Set up running the model as a subprocess while capturing the output
	modelSubProcess := exec.Command("./run_model")

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

	// 3. Run the model
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
			currentLineFromStdout := stdoutScanner.Text()

			if strings.Contains(currentLineFromStdout, "<PIL.Image.Image") {
				// Another in progress.png was created by the model!
				numberOfTimesInProgressPngWasCreated++
				iterationStatus := numberOfTimesInProgressPngWasCreated * 50

				err = postJobdUpdateToScheduler(strconv.Itoa(iterationStatus), JOB_IS_NOT_DONE)
				if err != nil {
					println("error after posting update to scheduler: ", err)
				}

			} else if strings.Contains(currentLineFromStdout, "finished by user") {
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

	if WORKER_HAS_TUNNEL {
		req.AddCookie(&http.Cookie{Name: exapi.CookieWorkerTunnel, Value: TUNNEL_URL})
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

// Takes the second startupt argument as URL to be used for the tunnel
func setTunnelUrl() (string, bool) {

	if len(os.Args) > 2 {
		return os.Args[2], true
	}
	return "", false
}

// This is the main function :D
func main() {

	initializeLogFile()
	SECRET = exutils.InitializeSecretFromArgument()
	TUNNEL_URL, WORKER_HAS_TUNNEL = setTunnelUrl()

	registerWorkerWithScheduler()

	http.HandleFunc("/api/0/worker", handleNewJobPosting) // Listen for new jobs on this endpoint

	go http.ListenAndServe(WORKER_PORT, nil)
	fmt.Println("Worker is running, waiting for job posts..") // Debug

	// Run worker loop
	runWorkerLoop()
}
