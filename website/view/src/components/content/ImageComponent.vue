<template>
  <div>
    <div class="mb-5 mt-5">
      <img :src="getSelectedImage" class="img-fluid" alt="..." />
    </div>
  </div>
</template>

<script>
export default {
  name: "ImageComponent",
  computed: {
    getSelectedImage() {
      let selectedImg = this.$store.getters.getSelectedJob.img;
      return selectedImg == undefined
        ? "https://via.placeholder.com/1920x1024.png?text=This is zen's placeholder"
        : selectedImg;
    },
  },
};
</script>

<style></style>
