import axios from "axios";

const url = "https://exia.art/api/0";

const state = {
  jobs: [],
  selectedJob: [],
  jobRange: {},
  jobsExist: true,
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
  // Fetch InitialJobs
  async fetchInitialJobs({ commit }) {
    state.jobRange.jobx = 1;
    state.jobRange.joby = 10;
    try {
      return await axios.get(`${url}/jobs?jobx=1&joby=10`).then((response) => {
        if (response.status == 200) {
          const payload = response.data;
          commit("FETCH_INITIAL_JOBS", payload);
        }
      });
    } catch (error) {
      console.log(error);
    }
  },
  // Fetch AdditionalJobs
  async fetchAdditionalJobs({ commit }) {
    if (state.jobsExist) {
      commit("INCREMENT_JOBRANGE", { amount: 10 });
      try {
        return await axios
          .get(
            `${url}/jobs?jobx=${state.jobRange.jobx}&joby=${state.jobRange.joby}`
          )
          .then((response) => {
            if (response.status == 200) {
              const payload = response.data;
              if (payload == null) {
                commit("DECREMENT_JOBRANGE", { amount: 10 });
              }
              commit("FETCH_ADDITIONAL_JOBS", payload);
            }
          });
      } catch (error) {
        console.log(error);
      }
    }
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
  INCREMENT_JOBRANGE(state, payload) {
    state.jobRange.jobx += payload.amount;
    state.jobRange.joby += payload.amount;
  },
  DECREMENT_JOBRANGE(state, payload) {
    state.jobRange.jobx -= payload.amount;
    state.jobRange.joby -= payload.amount;
  },
  FETCH_INITIAL_JOBS(state, payload) {
    state.jobs = payload;
  },
  FETCH_ADDITIONAL_JOBS(state, payload) {
    if (payload != null) {
      state.jobs.push(...payload);
    } else {
      state.jobsExist = false;
    }
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
