<template>
  <div>
    <div class="mb-5 mt-5">
      <div v-if="isLoading" class="spinner-border text-secondary" role="status">
        <span class="visually-hidden">Loading...</span>
      </div>
      <img :src="imgObjectURL" class="img-fluid" alt="" />
    </div>
  </div>
</template>

<script>
export default {
  name: "ImageComponent",
  data() {
    return {
      imgObjectURL: "",
      isLoading: false,
    };
  },
  methods: {
    createImgObjectURL() {
      this.isLoading = true;
      this.$store
        .dispatch("getSelectedImg")
        .then((response) => {
          var img = new Image();
          let blob = new Blob([response], { type: "image/png" });
          let url = URL.createObjectURL(blob);
          img.src = url;
          this.imgObjectURL = img.src;
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
  computed: {
    getSelectedJob() {
      return this.$store.getters.getSelectedJob;
    },
  },
  watch: {
    getSelectedJob: {
      handler() {
        this.createImgObjectURL();
      },
    },
  },
};
</script>

<style></style>
