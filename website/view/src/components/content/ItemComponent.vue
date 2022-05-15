<template>
  <li
    :disabled="job.job_status != 'completed'"
    :class="[job.job_status != 'completed' ? 'disabled' : '']"
    class="list-group-item list-group-item-action"
  >
    <div class="row">
      <div class="col-lg-10 col-md-10">
        <p @click="onClickSetSelected($event)" class="text-start">
          {{ job.prompt }}
        </p>
      </div>
      <div class="col-lg-2 col-md-2">
        <div :class="getJobBorderClass" class="badge border text-secondary">
          {{ job.job_status }}
        </div>
        <div class="progress mt-1">
          <div
            :style="`width: ${getProgressbarPercent}%;`"
            class="progress-bar progress-bar-animated"
            role="progressbar"
            :aria-valuenow="getProgressbarPercent"
            aria-valuemin="0"
            aria-valuemax="100"
          >
            {{ `${getProgressbarPercent}%` }}
          </div>
        </div>
      </div>
    </div>
  </li>
</template>

<script>
export default {
  name: "ItemComponent",
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
    onClickSetSelected(e) {
      let selectedLi = e.target.parentElement.parentElement.parentElement;
      let ulChildren =
        document.getElementsByClassName("list-group-flush")[0].children;
      let liActiveClass = "item-group-active";

      for (let i = 0; i < ulChildren.length; i++) {
        ulChildren[i].classList.remove(liActiveClass);
      }

      selectedLi.classList.add(liActiveClass);

      if (this.job.job_status != "completed") {
        e.preventDefault();
      }
      this.$store.dispatch("getSelectedJob", this.job.jobid);
    },
  },
  computed: {
    getSelectedJob() {
      return this.$store.getters.getSelectedJob;
    },
    getJobBorderClass() {
      let returnJobStatus;
      switch (this.job.job_status) {
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
    getProgressbarPercent() {
      return (this.job.iteration_status / this.job.iteration_max) * 100;
    },
  },
};
</script>

<style scoped>
li.list-group-item {
  background: transparent;
  cursor: pointer;
  transition: 0.1s background ease-out;
  --color-highlight: 204, 204, 204;
}
li p {
  padding: 2px 0px 20px 0px;
  margin: 0;
}
li.list-group-item:hover,
li.list-group-item:active,
li.list-group-item:focus,
li.list-group-item.item-group-active {
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
