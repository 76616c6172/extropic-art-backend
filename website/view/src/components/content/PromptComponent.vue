<template>
  <div>
    <div class="input-group">
      <input
        v-model="vPrompt"
        type="text"
        class="form-control"
        placeholder="Enter your prompt"
        aria-label="Enter your prompt"
      />
      <button
        @click="onClickSendNewJob()"
        class="btn btn-outline-secondary"
        type="submit"
      >
        Generate
      </button>
    </div>
  </div>
</template>

<script>
export default {
  name: "PromptComponent",
  data() {
    return {
      vPrompt: "",
    };
  },
  methods: {
    onClickSendNewJob() {
      this.$store.dispatch("sendNewJob", { prompt: this.vPrompt });
    },
  },
};
</script>
