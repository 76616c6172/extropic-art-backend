import axios from "axios";

const url = "https://exia.art/api/0";

const state = {
  jobs: [],
  selectedJob: [],
};
const getters = {
  getJobs: (state) => {
    return state.jobs;
  },
  getSelectedJob: (state) => {
    return state.selectedJob;
  },
};
const actions = {
  // Fetch Jobs
  async fetchJobs({ commit }) {
    try {
      return await axios.get(`${url}/status`).then((response) => {
        if (response.status == 200) {
          const payload = response.data.completed_jobs;
          commit("FETCH_JOBS", payload);
        }
      });
    } catch (error) {
      console.log(error);
    }
  },
  // Send Job
  async sendNewJob(newJobObj) {
    console.log(newJobObj);
    try {
      return await axios.post(`${url}/jobs`, newJobObj).then((response) => {
        if (response.status == 200) {
          console.log(response);
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
};
const mutations = {
  FETCH_JOBS(state, payload) {
    state.jobs = payload;
  },
  FETCH_SELECTED_JOB(state, payload) {
    state.selectedJob = payload;
    console.log(state.selectedJob);
  },
};

const apiModule = {
  state,
  mutations,
  actions,
  getters,
};

export default apiModule;
