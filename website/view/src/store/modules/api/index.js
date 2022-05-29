import axios from "axios";

const url = "https://exia.art/api/0";

const state = {
  jobs: [],
  selectedJob: [],
  jobRange: {},
  newestJobID: "",
  isOldestJobID: false,
};
const getters = {
  getJobs: (state) => {
    return state.jobs;
  },
  getSelectedJob: (state) => {
    return state.selectedJob;
  },
  getJobRange: (state) => {
    return state.jobRange;
  },
  getJobsExist: (state) => {
    return state.jobsExist;
  },
};
const actions = {
  // Get Newest Job
  async fetchNewestJobID({ commit }, scrollEvent) {
    try {
      return await axios.get(`${url}/status`).then((response) => {
        if (response.status == 200) {
          const newestJobId = Number(response.data.newest_jobid);
          switch (scrollEvent) {
            case "initial":
              commit("SET_JOBRANGE", {
                jobx: newestJobId - 9,
                joby: newestJobId,
              });
              break;
            case "add":
              commit("SET_JOBRANGE", {
                jobx:
                  state.jobRange.jobx > 1
                    ? state.jobRange.jobx - 10 > 0
                      ? state.jobRange.jobx - 10
                      : (state.jobRange.jobx = 1)
                    : (state.jobRange.jobx = 1),
                joby:
                  state.jobRange.joby > 1
                    ? state.jobRange.joby - 10
                    : (state.jobRange.joby = 1),
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
    dispatch("fetchNewestJobID", scrollEvent).then(() => {
      if (!state.isOldestJobID) {
        try {
          return axios
            .get(
              `${url}/jobs?jobx=${state.jobRange.jobx}&joby=${state.jobRange.joby}`
            )
            .then((response) => {
              if (response.status == 200) {
                const payload = response.data.sort((job) => job.id).reverse();
                commit("FETCH_JOBS", payload);
                state.jobRange.jobx == 1 ? (state.isOldestJobID = true) : "";
              }
            });
        } catch (error) {
          console.log(error);
        }
      }
    });
  },
  // Send Job
  async sendNewJob({ commit }, newJobObj) {
    try {
      return await axios.post(`${url}/jobs`, newJobObj).then((response) => {
        if (response.status == 200) {
          const payload = response.data;
          commit("SEND_NEW_JOB", payload);
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
  SET_JOBRANGE(state, payload) {
    state.jobRange.jobx = payload.jobx;
    state.jobRange.joby = payload.joby;
  },
  FETCH_JOBS(state, payload) {
    state.jobs.push(...payload);
  },
  FETCH_SELECTED_JOB(state, payload) {
    state.selectedJob = payload;
  },
  SEND_NEW_JOB(state, payload) {
    state.jobs.push(payload);
  },
};

const apiModule = {
  state,
  mutations,
  actions,
  getters,
};

export default apiModule;
