<template>
  <div>
    <div class="row">
      <div class="col-4">
        <!-- Total Jobs -->
        <h2 v-if="!isPageLoaded" class="display-3 text-start">
          {{ getText("counterTotal") }}
        </h2>
        <h2 v-else class="display-3 text-start">
          {{ jobStatusObj.newestJobId }}
        </h2>
        <p v-if="jobStatusObj.newestJobId" class="text-start">Images total</p>
      </div>
      <div class="col-4">
        <!-- Jobs Queued -->
        <h2 v-if="!isPageLoaded" class="display-3 text-start">
          {{ getText("counterQueued") }}
        </h2>
        <h2 v-else class="display-2 text-start">
          {{ jobStatusObj.jobsQueued }}
        </h2>
        <p v-if="jobStatusObj.jobsQueued" class="text-start">Images queued</p>
      </div>
      <div class="col-4">
        <!-- Jobs Completed -->
        <h2 v-if="!isPageLoaded" class="display-3 text-start">
          {{ getText("counterCompleted") }}
        </h2>
        <h2 v-else class="display-3 text-start">
          {{ jobStatusObj.jobsCompleted }}
        </h2>
        <p v-if="jobStatusObj.jobsCompleted" class="text-start">
          Images completed
        </p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "StatsComponent",
  data() {
    return {
      counterObj: {
        counterTotal: 0,
        counterQueued: 0,
        counterCompleted: 0,
      },
      jobStatusObj: {},
      jobStatusObjProperty: "",
    };
  },
  props: {
    jobStatus: {
      type: Object,
      required: true,
      jobRange: {
        type: Object,
        required: true,
      },
      jobsCompleted: {
        type: String,
        required: true,
      },
      jobsQueued: {
        type: String,
        required: true,
      },
      newestJobId: {
        type: String,
        required: true,
      },
      newestCompletedJobs: {
        type: Array,
        required: true,
      },
    },
  },
  methods: {
    getText(type) {
      return this.counterObj[type];
    },
    delay(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
    async setText(type) {
      if (type !== this.jobStatus[this.jobStatusObjProperty]) {
        switch (type) {
          case "counterTotal":
            this.jobStatusObjProperty = "newestJobId";
            break;
          case "counterQueued":
            this.jobStatusObjProperty = "jobsQueued";
            break;
          case "counterCompleted":
            this.jobStatusObjProperty = "jobsCompleted";
            break;
          default:
            break;
        }
      }
      if (this.counterObj[type] < this.jobStatus[this.jobStatusObjProperty]) {
        await this.delay(150);
        this.counterObj[type]++;
        this.setText(type);
      } else {
        sessionStorage.setItem(this.$options.name, "statsComponent");
      }
    },
  },
  computed: {
    isPageLoaded() {
      return sessionStorage.getItem(this.$options.name) == "statsComponent";
    },
  },
  watch: {
    jobStatus: {
      handler() {
        this.jobStatusObj = this.jobStatus;
        this.setText("counterTotal");
        this.setText("counterQueued");
        this.setText("counterCompleted");
      },
    },
  },
};
</script>

<style></style>
