<template>
  <div>
    <div class="row">
      <div
        v-for="(el, index) in blobObjArray"
        :key="index"
        class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1"
      >
        <img
          @click="onClickSetSelected(el.jobId)"
          :src="el.imgURL"
          class="img-fluid img-thumbnail"
          alt=""
        />
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
      this.$store.dispatch("getSelectedImg", jobId).then((response) => {
        this.blobObjArray.push({ jobId: jobId, imgURL: response });
      });
    },
    onClickSetSelected(jobId) {
      this.$store.dispatch("fetchSingleJob", jobId).then(() => {
        console.log("triggered");
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
        index < 3 ? this.createImgObjectURL(jobId) : ""
      );
  },
};
</script>

<style></style>
