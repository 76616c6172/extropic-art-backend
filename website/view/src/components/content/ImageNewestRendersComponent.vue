<template>
  <div>
    <div class="row">
      <div
        v-for="(job, index) in imgArray"
        :key="index"
        class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1"
      >
        <img
          @click="onClickSetSelected(job.jobId)"
          :src="job.imgURL"
          class="img-fluid img-thumbnail"
          alt=""
        />
        <p>{{ job.prompt }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ImageNewestRendersComponent",
  data() {
    return {
      imgArray: [],
    };
  },
  props: ["newestJobIds"],
  methods: {
    createImgObjectURL(jobId) {
      return new Promise((resolve) => {
        this.$store
          .dispatch("getSelectedImg", jobId)
          .then((responseImg) => {
            this.imgArray.push({
              jobid: jobId,
              imgURL: responseImg,
            });
          })
          .finally(() => {
            resolve();
          });
      });
    },
    async getSelectedJobsObject(jobId) {
      await this.createImgObjectURL(jobId).then(() => {
        if (this.imgArray.length == 3) {
          let newestJobIdsValues = Object.values(this.newestJobIds);
          let newestJobIdsMax = Math.max(...newestJobIdsValues);
          let newestJobIdsMin = Math.min(...newestJobIdsValues);
          this.$store
            .dispatch("getSelectedJobs", {
              jobx: newestJobIdsMin,
              joby: newestJobIdsMax,
              jobIds: newestJobIdsValues,
            })
            .then(() => {
              this.$store.getters.getSelectedJobs.forEach((storeJobElement) => {
                this.imgArray.forEach((imgArrayElement) => {
                  if (imgArrayElement.jobid == storeJobElement.jobid) {
                    imgArrayElement.prompt = storeJobElement.prompt;
                  }
                });
              });
            });
        }
      });
    },
    onClickSetSelected() {},
  },
  computed: {
    getIsInitialLoadStatus() {
      return this.$store.getters.getIsInitialLoadStatus;
    },
  },
  async mounted() {
    if (this.imgArray.length == 0) {
      this.newestJobIds.map((jobId, index) =>
        index <= 3 ? this.getSelectedJobsObject(jobId) : ""
      );
    }
  },
};
</script>

<style scoped>
p {
  white-space: nowrap;
  width: 100%;
  overflow: hidden;
  -o-text-overflow: ellipsis;
  text-overflow: ellipsis;
}
</style>
