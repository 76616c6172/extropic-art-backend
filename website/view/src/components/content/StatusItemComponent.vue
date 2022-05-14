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
        <div
          :class="getJobBorderClass(job.job_status)"
          class="badge border text-secondary"
        >
          {{ job.job_status }}
        </div>
        <!-- {{ job.iteration_status }}/{{ job.iteration_max }} -->
        <div class="progress mt-1">
          <div
            :style="`width: ${getProgressbarPercent(
              job.iteration_status,
              job.iteration_max
            )}%;`"
            class="progress-bar progress-bar-animated"
            role="progressbar"
            :aria-valuenow="
              getProgressbarPercent(job.iteration_status, job.iteration_max)
            "
            aria-valuemin="0"
            aria-valuemax="100"
          >
            {{
              `${getProgressbarPercent(
                job.iteration_status,
                job.iteration_max
              )}%`
            }}
          </div>
        </div>
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
    getProgressbarPercent(iterationStatus, iterationMax) {
      return (iterationStatus / iterationMax) * 100;
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
li.list-group-item {
  background: transparent;
  padding: 20px 0px 15px 0px;
  cursor: pointer;
  transition: 0.1s background ease-out;
  --color-highlight: 204, 204, 204;
}
li.list-group-item:hover,
li.list-group-item:active,
li.list-group-item:focus {
  background: #dedede;
  background-color: rgba(var(--color-highlight), 0.1);
}
.badge {
  width: 100%;
}
p {
  font-family: "Rubik-Light";
  white-space: nowrap;
  width: 100%; /* IE6 needs any width */
  overflow: hidden; /* "overflow" value must be different from  visible"*/
  -o-text-overflow: ellipsis; /* Opera < 11*/
  text-overflow: ellipsis; /* IE, Safari (WebKit), Opera >= 11, FF > 6 */
}
.progress {
  background: transparent;
}
.progress-bar {
  color: gray;
  background-color: #eee;
}
</style>
