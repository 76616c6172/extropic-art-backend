<template>
  <div>
    {{ job }}
  </div>
</template>

<script>
export default {
  name: "SideBarComponent",
  props: {
    job: {
      type: Object,
      required: true,
    },
  },
};
</script>

<style></style>
