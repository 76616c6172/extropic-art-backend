import axios from "axios";

const url = "https://exia.art/api/0";

const state = {
  jobs: [],
};
const getters = {
  getJobs: (state) => {
    return state.jobs;
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
  // Fetch Image
  // params: {
  //   jobid: "752b5cd9013dcf3f6ebf577f99fa76adf4f32459",
  // },
  // headers: {
  //   "Content-Type": "application/json",
  // },
  async fetchImage() {
    try {
      return await axios
        .post(
          `${url}/img`,
          {
            headers: {
              "Content-Type": "application/json",
            },
          },
          {
            data: {
              jobid: "752b5cd9013dcf3f6ebf577f99fa76adf4f32459",
            },
          }
        )
        .then((response) => {
          if (response.status == 200) {
            console.log(
              Buffer.from(response.data, "binary").toString("base64")
            );
            console.log(response);
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
};

const apiModule = {
  state,
  mutations,
  actions,
  getters,
};

export default apiModule;
