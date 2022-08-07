<template>
  <div>
    <img :src="imgObjectURL" class="img-fluid img-thumbnail" alt="" />
  </div>
</template>

<script>
export default {
  name: "FullSizeImageComponent",
  props: {
    imgObjectURL: {
      Type: String,
      required: true,
    },
  },
};
</script>

<style></style>
