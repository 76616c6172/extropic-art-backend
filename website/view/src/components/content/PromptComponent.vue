<template>
  <div class="pb-5">
    <div class="input-group">
      <input
        @keyup.enter="onClickSendNewJob()"
        v-model="vPrompt"
        ref="inputPrompt"
        type="text"
        class="form-control"
        placeholder="Enter your prompt"
        aria-label="Enter your prompt"
      />
      <!-- <span v-if="showCursor" class="blink">|</span> -->
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
  props: {
    showCursor: {
      type: Boolean,
      required: true,
    },
  },
  methods: {
    onClickSendNewJob() {
      var regex = /^\s+$/;
      if (this.vPrompt != "" && !this.vPrompt.match(regex)) {
        this.$store.dispatch("sendNewJob", { prompt: this.vPrompt });
        this.vPrompt = "";
      }
    },
  },
  computed: {
    setAutofocus() {
      return this.showCursor ? this.$refs.inputPrompt.focus() : "";
    },
  },
  watch: {
    showCursor: {
      handler() {
        setTimeout(() => {
          this.setAutofocus;
        }, 1500);
      },
    },
  },
};
</script>
