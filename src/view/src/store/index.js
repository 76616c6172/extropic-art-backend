import { createStore } from "vuex";

import apiModule from "./modules/api";

export default createStore({
  modules: {
    api: apiModule,
  },
});
