<template>
  <div class="mt-5 mb-4">
    <div class="row">
      <h2 class="text-start mt-5 mb-4">
        <strong> ImageList</strong>
      </h2>
      <form class="col-4 mb-3 mb-lg-0">
        <input
          v-model="searchQuery"
          type="search"
          class="form-control"
          placeholder="Search..."
          aria-label="Search"
        />
      </form>
    </div>
    <ul class="list-group-flush" style="padding-left: 0 !important">
      <StatusItem
        v-for="(job, index) in getSearchedProducts"
        :key="index"
        :job="job"
      />
    </ul>
  </div>
</template>

<script>
import { default as StatusItem } from "./StatusItemComponent.vue";

export default {
  name: "StatusListComponent",
  components: {
    StatusItem,
  },
  data() {
    return {
      jobs: [],
      searchQuery: "",
    };
  },
  methods: {
    loadJobs(jobs) {
      if (this.jobs.length == 0) {
        jobs.forEach((job) => {
          job.prompt != "" ? this.jobs.push(job) : "";
        });
        this.jobs = this.jobs.sort((job) => job.job_status == "completed");
      }
    },
  },
  computed: {
    getJobsFromStore() {
      return this.$store.getters.getJobs;
    },
    getSearchedProducts() {
      return this.jobs.filter((job) => {
        return (
          job.prompt.toLowerCase().indexOf(this.searchQuery.toLowerCase()) != -1
        );
      });
    },
  },
  watch: {
    getJobsFromStore: {
      handler(jobs) {
        // console.log("here");
        if (jobs) {
          this.loadJobs(jobs);
        }
      },
      immediate: true,
    },
  },
  async mounted() {
    this.$store.dispatch("fetchJobs");
  },
};
</script>
