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
  // Set JobRange
  setJobRange({ commit }) {
    const jobRange = {
      jobx: (state.jobRange.jobx += 10),
      joby: (state.jobRange.joby += 10),
    };
    commit("SET_JOB_RANGE", jobRange);
  },
  // Fetch InitialJobs
  async fetchInitialJobs({ commit }) {
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
    try {
      return await axios
        .get(
          `${url}/jobs?jobx=${state.jobRange.jobx}&joby=${state.jobRange.joby}`
        )
        .then((response) => {
          if (response.status == 200) {
            const payload = response.data;
            commit("FETCH_ADDITIONAL_JOBS", payload);
          }
        });
    } catch (error) {
      console.log(error);
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
  FETCH_INITIAL_JOBS(state, payload) {
    state.jobRange.jobx = 1;
    state.jobRange.joby = 10;
    state.jobs = payload;
  },
  FETCH_ADDITIONAL_JOBS(state, payload) {
    if (payload == null) {
      state.jobsExist = false;
      state.jobRange.jobx -= 10;
      state.jobRange.joby -= 10;
    } else {
      state.jobs.push(...payload);
    }
  },
  FETCH_SELECTED_JOB(state, payload) {
    state.selectedJob = payload;
  },
  SEND_NEW_JOB(state, payload) {
    state.jobs.push(payload);
  },
  SET_JOB_RANGE(state, payload) {
    state.jobRange.jobx = payload.jobx;
    state.jobRange.joby = payload.joby;
  },
};

const apiModule = {
  state,
  mutations,
  actions,
  getters,
};

export default apiModule;
