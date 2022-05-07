<template>
  <li
    @click="onClickSetSelected(jobId)"
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
    onClickSetSelected(jobId) {
      this.$store.dispatch("fetchSelectedImage", jobId);
      this.$store.dispatch("fetchSelectedPrompt", jobId);
    },
  },
};
</script>
