import axios from "axios";

const url = "https://exia.art/api/0";

const state = {
  jobs: [],
  selectedJob: [],
  jobStatus: {
    jobRange: {},
    jobsCompleted: "",
    jobsQueued: "",
    newestJobId: "",
    newestCompletedJobs: [],
  },
  isOldestJobID: false,
  isInitialLoad: false,
};
const getters = {
  getIsInitialLoadStatus: (state) => {
    return state.isInitialLoad;
  },
  getJobs: (state) => {
    return state.jobs;
  },
  getSelectedJob: (state) => {
    return state.selectedJob;
  },
  getJobStatus: (state) => {
    return state.jobStatus;
  },
  getJobsExist: (state) => {
    return state.jobsExist;
  },
};
const actions = {
  // Get Newest Job
  async fetchJobStatus({ commit }, scrollEvent) {
    try {
      return await axios.get(`${url}/status`).then((response) => {
        if (response.status == 200) {
          // console.log(response.data);
          const data = response.data;
          const newestJobId = Number(data.newest_jobid);
          switch (scrollEvent) {
            case "initial":
              commit("SET_JOBSTATUS", {
                jobRange: {
                  jobx: newestJobId - 9 >= 1 ? newestJobId - 9 : 1,
                  joby: newestJobId,
                },
                jobsCompleted: data.Jobs_completed,
                jobsQueued: data.Jobs_queued,
                newestJobId: newestJobId,
                newestCompletedJobs: [...data.Newest_completed_jobs],
              });
              break;
            case "add":
              commit("SET_JOBSTATUS", {
                jobRange: {
                  jobx:
                    state.jobStatus.jobRange.jobx > 1
                      ? state.jobStatus.jobRange.jobx - 10 > 0
                        ? state.jobStatus.jobRange.jobx - 10
                        : (state.jobStatus.jobRange.jobx = 1)
                      : (state.jobStatus.jobRange.jobx = 1),
                  joby:
                    state.jobStatus.jobRange.joby > 1
                      ? state.jobStatus.jobRange.joby - 10
                      : (state.jobStatus.jobRange.joby = 1),
                },
                jobsCompleted: data.Jobs_completed,
                jobsQueued: data.Jobs_queued,
                newestJobId: newestJobId,
                newestCompletedJobs: [...data.Newest_completed_jobs],
              });
              break;
            default:
              break;
          }
        }
      });
    } catch (error) {
      console.log(error);
    }
  },
  // Fetch InitialJobs
  async fetchJobs({ commit, dispatch }, scrollEvent) {
    switch (scrollEvent) {
      case "initial":
        if (!state.isOldestJobID) {
          try {
            return axios
              .get(
                `${url}/jobs?jobx=${state.jobStatus.jobRange.jobx}&joby=${state.jobStatus.jobRange.joby}`
              )
              .then((response) => {
                if (response.status == 200) {
                  const payload = response.data.sort((job) => job.id).reverse();
                  commit("FETCH_JOBS", payload);
                  state.jobStatus.jobRange.jobx == 1
                    ? (state.isOldestJobID = true)
                    : "";
                }
              });
          } catch (error) {
            console.log(error);
          }
        }
        break;
      default:
        dispatch("fetchJobStatus", scrollEvent).then(() => {
          if (!state.isOldestJobID) {
            try {
              return axios
                .get(
                  `${url}/jobs?jobx=${state.jobStatus.jobRange.jobx}&joby=${state.jobStatus.jobRange.joby}`
                )
                .then((response) => {
                  if (response.status == 200) {
                    const payload = response.data
                      .sort((job) => job.id)
                      .reverse();
                    commit("FETCH_JOBS", payload);
                    state.jobStatus.jobRange.jobx == 1
                      ? (state.isOldestJobID = true)
                      : "";
                  }
                });
            } catch (error) {
              console.log(error);
            }
          }
        });
        break;
    }
  },
  // Send Job
  async sendNewJob({ commit }, newJobObj) {
    console.log(newJobObj);
    try {
      return await axios.post(`${url}/jobs`, newJobObj).then((response) => {
        if (response.status == 200) {
          console.log(response.data);
          const newJobID = response.data.jobid;
          try {
            return axios
              .get(`${url}/jobs?jobx=${newJobID}&joby=${newJobID}`)
              .then((response) => {
                if (response.status == 200) {
                  const payload = response.data[0];
                  commit("SEND_NEW_JOB", payload);
                }
              });
          } catch (error) {
            console.log(error);
          }
        }
      });
    } catch (error) {
      console.log(error);
    }
  },
  getSelectedJob({ commit }, jobId) {
    const payload = this.getters.getJobs.filter(
      (job) => job.jobid === jobId
    )[0];
    payload.img = `${url}/img?${jobId}`;
    commit("FETCH_SELECTED_JOB", payload);
  },
  async getSelectedImg() {
    let imgURL = state.selectedJob.img_path;
    try {
      return await axios
        .get(`${imgURL}`, { responseType: "blob" })
        .then((response) => {
          if (response.status == 200) {
            return new Promise((resolve) => {
              const payload = response.data;
              resolve(payload);
            });
          }
        });
    } catch (error) {
      console.log(error);
    }
  },
};
const mutations = {
  SET_JOBSTATUS(state, payload) {
    state.jobStatus = payload;
    state.isInitialLoad == false ? (state.isInitialLoad = true) : "";
  },
  FETCH_JOBS(state, payload) {
    state.jobs.push(...payload);
  },
  FETCH_SELECTED_JOB(state, payload) {
    state.selectedJob = payload;
  },
  SEND_NEW_JOB(state, payload) {
    state.jobs.unshift(payload);
  },
};

const apiModule = {
  state,
  mutations,
  actions,
  getters,
};

export default apiModule;
