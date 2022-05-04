<template>
  <li
    @click="onClickSetImage"
    style="list-style-type: none; background-color: lightgrey"
  >
    <p class="text-start"><strong>Prompt: </strong>{{ prompt }}</p>
    <p class="text-start">(<strong>JobID: </strong> {{ jobId }})</p>
  </li>
</template>

<script>
export default {
  name: "StatusItemComponent",
  props: {
    prompt: {
      type: String,
      required: true,
    },
    jobId: {
      type: String,
      required: true,
    },
  },
  methods: {
    onClickSetImage() {
      console.log(this.jobId);
      this.$store.dispatch("fetchImage");
    },
  },
};
</script>
