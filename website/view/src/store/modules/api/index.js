import axios from "axios";

const url = "https://exia.art/api/0";

const state = {
  jobs: [],
  selectedJob: [],
  jobRange: {
    jobx: 1,
    joby: 10,
  },
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
};
const actions = {
  // Set JobRange
  setJobRange(_, scrollDirection) {
    switch (scrollDirection) {
      case "down":
        state.jobRange.jobx += 1;
        state.jobRange.joby += 10;
        break;
      case "up":
        if (state.jobRange.jobx == 1) {
          state.jobRange.jobx = 1;
          state.jobRange.joby = 10;
        } else {
          state.jobRange.jobx -= 1;
          state.jobRange.joby -= 10;
        }
        break;
    }
  },
  // Fetch Jobs
  async fetchJobs({ commit }) {
    try {
      return await axios
        .get(
          `${url}/jobs?jobx=${state.jobRange.jobx}&joby=${state.jobRange.joby}`
        )
        .then((response) => {
          if (response.status == 200) {
            const payload = response.data;
            commit("FETCH_JOBS", payload);
          }
        });
    } catch (error) {
      console.log(error);
    }
  },
  // async fetchJobs({ commit }) {
  //   try {
  //     return await axios.get(`${url}/status`).then((response) => {
  //       if (response.status == 200) {
  //         const payload = response.data.completed_jobs;
  //         commit("FETCH_JOBS", payload);
  //       }
  //     });
  //   } catch (error) {
  //     console.log(error);
  //   }
  // },
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
  FETCH_JOBS(state, payload) {
    state.jobs = payload;
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
