<template>
  <div>
    <div class="row">
      <div
        v-for="(job, index) in blobObjArray"
        :key="index"
        class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1"
      >
        <img
          @click="onClickSetSelected(job.jobId)"
          :src="job.imgURL"
          class="img-fluid img-thumbnail"
          alt=""
        />
        <p>{{ job.jobObject.prompt }}</p>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ImageNewestRendersComponent",
  data() {
    return {
      blobObjArray: [],
    };
  },
  props: ["newestJobIds"],
  methods: {
    createImgObjectURL(jobId) {
      this.$store.dispatch("getSelectedImg", jobId).then((responseImg) => {
        this.$store
          .dispatch("getSelectedJob", { jobId: jobId, type: "newRequest" })
          .finally(() => {
            this.blobObjArray.push({
              jobId: jobId,
              imgURL: responseImg,
              jobObject: this.$store.getters.getSelectedJob,
            });
          });
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
    if (this.blobObjArray.length == 0) {
      this.newestJobIds.map((jobId, index) =>
        index < 3 ? this.createImgObjectURL(jobId) : ""
      );
    }
  },
};
</script>

<style scoped>
p {
  white-space: nowrap;
  width: 100%; /* IE6 needs any width */
  overflow: hidden; /* "overflow" value must be different from  visible"*/
  -o-text-overflow: ellipsis; /* Opera < 11*/
  text-overflow: ellipsis; /* IE, Safari (WebKit), Opera >= 11, FF > 6 */
}
</style>
