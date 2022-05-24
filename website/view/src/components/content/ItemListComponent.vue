<template>
  <div>
    <div class="row">
      <div class="pt-3 pb-3">
        <form class="col-4 mb-3 mb-lg-0">
          <input
            v-model="searchQuery"
            type="search"
            class="form-control bg-transparent text-white"
            placeholder="Search..."
            aria-label="Search"
          />
        </form>
      </div>
    </div>
    <ul class="list-group-flush" style="padding-left: 0 !important">
      <Item v-for="(job, index) in getFilteredJobs" :key="index" :job="job" />
    </ul>
  </div>
</template>

<script>
import { default as Item } from "./ItemComponent.vue";

export default {
  name: "StatusListComponent",
  components: {
    Item,
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
<style scoped>
input {
  border: none;
}
</style>
