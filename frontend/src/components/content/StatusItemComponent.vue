<template>
  <li
    @click="onClickSetSelected(job.jobid)"
    class="list-group-item list-group-item-action"
  >
    <div class="row">
      <div class="col-lg-10 col-md-10">
        <p class="text-start">{{ job.prompt }}</p>
      </div>
      <div class="col-lg-2 col-md-2">
        <span
          :class="getJobBorderClass(job.job_status)"
          class="badge border text-secondary"
          >{{ job.job_status }}
        </span>
        <span class="badge border text-secondary"
          >{{ job.iteration_status }}/{{ job.iteration_max }}
        </span>
      </div>
    </div>
  </li>
  <ConfirmDialogue ref="confirmDialogue" />
</template>

<script>
import ConfirmDialogue from "../modal/Confirmdialogue.vue";
export default {
  name: "StatusItemComponent",
  components: {
    ConfirmDialogue,
  },
  props: {
    job: {
      type: Object,
      required: true,
    },
    iteration_max: {
      type: String,
      required: true,
    },
    iteration_status: {
      type: String,
      required: true,
    },
    job_status: {
      type: String,
      required: true,
    },
    job_id: {
      type: String,
      required: true,
    },
    prompt: {
      type: String,
      required: true,
    },
  },
  methods: {
    onClickSetSelected(jobId) {
      this.$store.dispatch("getSelectedJob", jobId);
      this.buildModalDialogue();
    },
    async buildModalDialogue() {
      console.log("triggered");
      const ok = await this.$refs.confirmDialogue.show({
        title: this.getSelectedJob.prompt,
        image: this.getSelectedJob.img,
        message: "",
        okButton: "Löschen",
      });
      if (ok) {
        // this.$store.dispatch("deleteTodo", "todoObj._id");
      }
    },
    getJobBorderClass(jobStatus) {
      let returnJobStatus;
      switch (jobStatus) {
        case "completed":
          returnJobStatus = "border-success";
          break;
        case "processing":
          returnJobStatus = "border-info";
          break;
        case "queued":
          returnJobStatus = "border-warning";
          break;
        default:
          break;
      }
      return returnJobStatus;
    },
  },
  computed: {
    getSelectedJob() {
      return this.$store.getters.getSelectedJob;
    },
  },
};
</script>

<style scoped>
.list-group-item {
  background: transparent;
  padding: 20px 5px 15px 5px;
}
.badge {
  width: 100%;
}
.badge.iterationStatus {
}
</style>
