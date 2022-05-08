<template>
  <div>
    <h2 class="text-start mt-5">
      <strong> How to </strong>
    </h2>
    <div class="row">
      <template v-for="(instruction, index) in instructions" :key="index">
        <div class="col-sm-3 col-lg-1 mt-2 mb-3">
          <h4 class="text-start">0{{ index }}</h4>
        </div>
        <div class="col-sm-9 col-lg-5 mt-2 mb-3">
          <h5 class="text-start">{{ instruction.content }}</h5>
        </div>
        <template v-if="index == 1 || index == 3">
          <hr class="d-lg-block d-sm-none d-xs-none" />
        </template>
      </template>
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
    };
  },
};
</script>
