<template>
  <div>
    <p>
      {{ getText() }}
      <span class="blink">|</span>
    </p>
  </div>
</template>

<script>
export default {
  name: "TypingComponent",
  data() {
    return {
      i: 0,
      txtOut: "",
      txtIn:
        "We leverage an AI Image generating technique called CLIP-Guided Diffusion to allow you to create compelling and beautiful images from just text inputs. Made with love by Zen and Valar!.",
    };
  },
  methods: {
    getText() {
      return this.txtOut;
    },
    delay(ms) {
      return new Promise((resolve) => setTimeout(resolve, ms));
    },
    async setText() {
      if (this.i <= this.txtIn.length) {
        this.txtOut += this.txtIn.charAt(this.i);
        await this.delay(40);
        this.i++;
        this.setText();
      }
    },
  },
  mounted() {
    this.setText();
  },
};
</script>
