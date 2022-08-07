<template>
  <div>
    <div class="row justify-content-center">
      <div class="col-lg-10 col-sm-12">
        <h1 class="text-start display-5">Welcome to Project Exia</h1>
        <div class="pt-3 pb-3">
          <div v-if="!isPageLoaded">
            <p class="text-start fs-5">
              {{ getText }}
              <span v-if="showCursor" class="blink">|</span>
            </p>
          </div>
          <div v-else>
            <p class="text-start fs-5">{{ textInput }}</p>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "TypingComponent",
  data() {
    return {
      i: 0,
      textOutput: "",
      textInput:
        "Run state of the art machine learning models in the cloud to generate high resolution images from just text!",
    };
  },
  methods: {
    delay(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
    async setText() {
      if (this.i <= this.textInput.length) {
        this.textOutput += this.textInput.charAt(this.i);
        await this.delay(20);
        this.i++;
        this.setText();
      } else {
        sessionStorage.setItem(this.$options.name, "typingComponent");
      }
    },
  },
  computed: {
    getText() {
      return this.textOutput;
    },
    showCursor() {
      if (this.i == this.textOutput.length + 1) {
        this.$emit("set-cursor");
      }
      return this.i <= this.textOutput.length ? true : false;
    },
    isPageLoaded() {
      return sessionStorage.getItem(this.$options.name) == "typingComponent";
    },
  },
  mounted() {
    this.setText();
  },
};
</script>
