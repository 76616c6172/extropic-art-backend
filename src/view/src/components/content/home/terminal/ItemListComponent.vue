<template>
  <div>
    <div class="row">
      <div class="col-9 mb-3 mb-lg-0 pt-3 pb-3">
        <form>
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
    <ul
      id="ul-list"
      class="list-group-flush"
      style="padding-left: 0 !important"
    >
      <Item
        v-for="(job, index) in getFilteredJobs"
        :key="index"
        :job="job"
        @click="onClickOpenFullImage(job)"
      />
    </ul>
  </div>
</template>

<script>
import { default as Item } from "./ItemComponent.vue";
import { mapGetters } from "vuex";

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
    handleScroll(event) {
      let ulElement = event.srcElement;
      let scrollTop = Math.ceil(ulElement.scrollTop);
      let containerHeight = ulElement.scrollHeight - ulElement.offsetHeight;

      if (scrollTop == containerHeight || scrollTop == containerHeight + 1) {
        this.$store.dispatch("fetchJobs", "add");
      }

      return;
    },
    onClickOpenFullImage(job) {
      let routeData = this.$router.resolve({
        name: "JobDetails",
        params: { jobId: job.jobid },
      });
      window.open(routeData.href, "_blank");
    },
  },
  computed: {
    getFilteredJobs() {
      let jobs = this.getJobs;
      return this.getFoundJobs(jobs);
    },
    ...mapGetters(["getJobs", "getIsInitialLoadStatus"]),
  },
  async mounted() {
    document
      .getElementById("ul-list")
      .addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    document
      .getElementById("ul-list")
      .addEventListener("scroll", this.handleScroll);
  },
  watch: {
    getIsInitialLoadStatus: {
      handler() {
        this.$store.dispatch("fetchJobs", "initial");
      },
    },
  },
};
</script>
<style scoped>
ul {
  max-height: 400px;
  overflow-y: auto;
  overflow-x: hidden;
}
input {
  border: none;
}
</style>
