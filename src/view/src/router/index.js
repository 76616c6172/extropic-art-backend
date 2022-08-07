import { createRouter, createWebHistory } from "vue-router";
import HomeComponent from "../components/content/home/HomeWrapperComponent.vue";
import FullSizeImageComponent from "../components/content/fullsizeImg/FullSizeImageWrapperCompoment.vue";

// create and config
const routes = [
  {
    path: "/",
    name: "Home",
    component: HomeComponent,
  },
  {
    path: "/jobdetails/:jobId",
    name: "JobDetails",
    component: FullSizeImageComponent,
    props: true,
  },
];

const router = createRouter({
  history: createWebHistory(process.env.BASE_URL),
  routes,
});

export default router;
