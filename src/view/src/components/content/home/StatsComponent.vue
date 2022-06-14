<template>
  <div class="row justify-content-center pb-5 pt-5">
    <div class="col-lg-10 col-sm-12">
      <div class="row">
        <div class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1">
          <!-- Total Jobs -->
          <h2 v-if="!isPageLoaded" class="display-3 text-start">
            {{ getText("counterTotal") }}
          </h2>
          <h2 v-else class="display-3 text-start">
            {{ jobStatus.newestJobId }}
          </h2>
          <p v-if="jobStatus.newestJobId" class="text-start">Images total</p>
        </div>
        <div class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1">
          <!-- Jobs Queued -->
          <h2 v-if="!isPageLoaded" class="display-3 text-start">
            {{ getText("counterQueued") }}
          </h2>
          <h2 v-else class="display-3 text-start">
            {{ jobStatus.jobsQueued }}
          </h2>
          <p v-if="jobStatus.jobsQueued" class="text-start">Images queued</p>
        </div>
        <div class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1">
          <!-- Jobs Completed -->
          <h2 v-if="!isPageLoaded" class="display-3 text-start">
            {{ getText("counterCompleted") }}
          </h2>
          <h2 v-else class="display-3 text-start">
            {{ jobStatus.jobsCompleted }}
          </h2>
          <p v-if="jobStatus.jobsCompleted" class="text-start">
            Images completed
          </p>
        </div>
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
      jobStatusProperty: "",
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
      if (type !== this.jobStatus[this.jobStatusProperty]) {
        switch (type) {
          case "counterTotal":
            this.jobStatusProperty = "newestJobId";
            break;
          case "counterQueued":
            this.jobStatusProperty = "jobsQueued";
            break;
          case "counterCompleted":
            this.jobStatusProperty = "jobsCompleted";
            break;
          default:
            break;
        }
      }
      if (this.counterObj[type] < this.jobStatus[this.jobStatusProperty]) {
        await this.delay(1);
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
        this.setText("counterTotal");
        this.setText("counterQueued");
        this.setText("counterCompleted");
      },
    },
  },
};
</script>

<style></style>
