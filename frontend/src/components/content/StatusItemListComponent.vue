<template>
  <div class="mt-5 mb-4">
    <h2 class="text-start mt-5 mb-4">
      <strong> ImageList</strong>
    </h2>
    <ul class="list-group-flush" style="padding-left: 0 !important">
      <StatusItem
        v-for="(job, index) in storeGetJobs"
        :key="index"
        :prompt="job.prompt"
        :jobId="job.jobid"
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
  computed: {
    storeGetJobs() {
      return this.$store.getters.getJobs;
    },
  },
  async mounted() {
    this.$store.dispatch("fetchJobs");
  },
};
</script>
