<template>
  <div>
    <ul style="display: block; margin: 0; padding: 0">
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
