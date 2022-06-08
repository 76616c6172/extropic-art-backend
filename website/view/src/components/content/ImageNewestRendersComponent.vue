<template>
  <div>
    <div class="row">
      <div class="col-12">
        <h2 class="display-8 text-start">Most Recent renders</h2>
      </div>
      <!-- <div
        v-for="(img, index) in blobObjArrayLocalStorage"
        :key="index"
        class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1"
      >
        <img :src="img" class="img-fluid img-thumbnail" alt="" />
      </div> -->
      <div
        v-for="(img, index) in blobObjArray"
        :key="index"
        class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1"
      >
        <img :src="img" class="img-fluid img-thumbnail" alt="" />
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
      blobObjArrayLocalStorage: [],
    };
  },
  props: ["newestJobIds"],
  methods: {
    createImgObjectURL(jobId) {
      this.$store.dispatch("getSelectedImg", jobId).then((response) => {
        this.blobObjArray.push(response);
        // LocalStorage
        // localStorage.setItem(`job-${jobId}`, response);
      });
    },
  },
  computed: {
    getIsInitialLoadStatus() {
      return this.$store.getters.getIsInitialLoadStatus;
    },
  },
  async mounted() {
    if (this.blobObjArray.length == 0)
      this.newestJobIds.map((jobId, index) =>
        index < 3
          ? this.createImgObjectURL(jobId)
          : this.blobObjArrayLocalStorage.push(
              localStorage.getItem(`job-${jobId}`)
            )
      );
  },
};
</script>

<style></style>
