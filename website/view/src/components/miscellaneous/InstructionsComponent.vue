<template>
  <div>
    <div class="text-start">
      <button
        @click="this.visible = !this.visible"
        type="button"
        class="btn btn-outline-secondary"
      >
        How to
      </button>
    </div>
    <div v-if="visible" class="mt-5">
      <div class="row">
        <template v-for="(instruction, index) in instructions" :key="index">
          <div class="col-sm-1 col-lg-1 mt-2 mb-3">
            <h4 class="text-start">0{{ index + 1 }}</h4>
          </div>
          <div class="col-sm-11 col-lg-5 mt-2 mb-3">
            <h5 class="text-start">{{ instruction.content }}</h5>
          </div>
          <template v-if="index == 1 || index == 3">
            <hr class="d-lg-block d-sm-none d-xs-none" />
          </template>
        </template>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "InstructionsComponent",
  data() {
    return {
      instructions: [
        { content: "Enter search term" },
        { content: "Click generate or hit enter" },
        { content: "Wait the image to be finished" },
        { content: "Enjoy and feel energized" },
      ],
      visible: false,
    };
  },
};
</script>

<style scoped></style>
