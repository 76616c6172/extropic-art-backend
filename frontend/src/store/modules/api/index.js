import axios from "axios";

const url = "https://exia.art/api/0";

const state = {
  jobs: [],
  selectedImgUrl: "",
  selectedPrompt: "",
};
const getters = {
  getJobs: (state) => {
    return state.jobs;
  },
  getSelectedPrompt: (state) => {
    return state.selectedPrompt;
  },
  getSelectedImageUrl: (state) => {
    return state.selectedImgUrl;
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
  async sendNewJob({ commit }, newJobObj) {
    console.log(newJobObj);
    try {
      return await axios
        .post(
          `${url}/jobs`,
          {
            headers: {
              "Content-Type": "text/json",
            },
          },
          newJobObj
        )
        .then((response) => {
          if (response.status == 200) {
            console.log(response);
            console.log("200");
            console.log(commit);
          }
        });
    } catch (error) {
      console.log(error);
    }
  },
  // Fetch Selected Image
  fetchSelectedImage({ commit }, jobId) {
    const selectedImgUrl = `${url}/img?${jobId}`;
    commit("FETCH_SELECTED_IMAGE", selectedImgUrl);
  },
  // Fetch Selected Job
  fetchSelectedPrompt({ commit }, jobId) {
    const selectedPrompt = this.getters.getJobs.filter(
      (job) => job.jobid === jobId
    )[0].prompt;
    commit("FETCH_SELECTED_PROMPT", selectedPrompt);
  },
};
const mutations = {
  FETCH_JOBS(state, payload) {
    state.jobs = payload;
  },
  FETCH_SELECTED_IMAGE(state, payload) {
    state.selectedImgUrl = payload;
  },
  FETCH_SELECTED_PROMPT(state, payload) {
    state.selectedPrompt = payload;
  },
};

const apiModule = {
  state,
  mutations,
  actions,
  getters,
};

export default apiModule;
