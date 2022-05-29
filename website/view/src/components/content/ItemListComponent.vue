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
      <div class="col-3">
        <button
          @click="refreshJobList"
          type="button"
          class="btn btn-transparent text-secondary pt-2 pb-2"
          id="refreshButton"
        >
          <i class="fa fa-refresh" aria-hidden="true"></i>
        </button>
      </div>
    </div>
    <ul
      id="ul-list"
      class="list-group-flush"
      style="padding-left: 0 !important"
    >
      <Item v-for="(job, index) in getFilteredJobs" :key="index" :job="job" />
    </ul>
    <!-- <div>{{ this.$store }}</div> -->
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
      let scrollTop = ulElement.scrollTop;
      let containerHeight = ulElement.scrollHeight - ulElement.offsetHeight;

      if (scrollTop == containerHeight) {
        this.$store.dispatch("fetchJobs", "add");
      }
      return;
    },
    refreshJobList() {
      this.$store.dispatch("fetchJobs", "initial");
    },
  },
  computed: {
    getFilteredJobs() {
      let jobs = this.getJobs;
      return this.getFoundJobs(jobs);
    },
    ...mapGetters(["getJobs"]),
  },
  watch: {
    getFilteredJobs: {
      handler(jobs) {
        if (jobs) {
          jobs.forEach((job) => {
            if (job.job_status == "accepted") {
              setTimeout(() => {
                this.$store.dispatch("fetchJobs", "initial");
              }, 1500);
            }
          });
        }
      },
    },
  },
  async mounted() {
    this.$store.dispatch("fetchJobs", "initial");
    document
      .getElementById("ul-list")
      .addEventListener("scroll", this.handleScroll);
  },
  unmounted() {
    document
      .getElementById("ul-list")
      .addEventListener("scroll", this.handleScroll);
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
.btn#refreshButton {
  padding: 10px !important;
}
</style>
