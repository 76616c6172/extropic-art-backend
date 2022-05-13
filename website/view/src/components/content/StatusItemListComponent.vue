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
        v-for="(job, index) in getFilteredJobs"
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
      searchQuery: "",
    };
  },
  methods: {
    getFoundJobs(jobs) {
      return jobs.filter((job) => {
        return (
          job.prompt.toLowerCase().indexOf(this.searchQuery.toLowerCase()) != -1
        );
      });
    },
  },
  computed: {
    getJobs() {
      return this.$store.getters.getJobs;
    },
    getFilteredJobs() {
      let jobs = this.getJobs;
      jobs.sort((job) => job.job_status == "completed");
      return this.getFoundJobs(jobs);
    },
  },
  watch: {
    getFilteredJobs: {
      handler(jobs) {
        if (jobs) {
          jobs.forEach((job) => {
            if (job.job_status == "accepted") {
              setTimeout(() => {
                this.$store.dispatch("fetchJobs");
              }, 1500);
            }
          });
        }
      },
    },
  },
  async mounted() {
    this.$store.dispatch("fetchJobs");
  },
};
</script>
