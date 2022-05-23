<template>
  <div v-if="showCursor" class="pb-5 col-12 mb-3 mb-lg-0">
    <div class="input-group">
      <i class="fa fa-arrow-right" aria-hidden="true"></i>
      <input
        @keyup.enter="onClickSendNewJob()"
        v-model="vPrompt"
        ref="inputPrompt"
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
        }, 500);
      },
    },
  },
};
</script>

<style scoped>
.input-group i {
  position: absolute;
  top: 10px;
  left: 10px;
  z-index: 4;
}

.input-group input {
  padding-left: 30px;
}
</style>
