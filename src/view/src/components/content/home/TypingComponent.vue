<template>
  <div>
    <h2 class="text-start display-5">Welcome to Project Exia</h2>
    <div><br></div>
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
</template>

<script>
export default {
  name: "TypingComponent",
  data() {
    return {
      i: 0,
      textOutput: "",
      textInput:
        "Exia lets you run state of the art machine learning models in the cloud to generate high resolution images from just text! Made with love by zen and valar in 2022.",
    };
  },
  methods: {
    delay(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
    async setText() {
      if (this.i <= this.textInput.length) {
        this.textOutput += this.textInput.charAt(this.i);
        await this.delay(25);
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
