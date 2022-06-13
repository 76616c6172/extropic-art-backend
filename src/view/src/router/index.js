import { createRouter, createWebHistory } from "vue-router";
import HomeComponent from "../components/content/home/HomeComponent.vue";

// create and config
const routes = [
  {
    path: "/",
    name: "Home",
    component: HomeComponent,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

// export router
export default router;
