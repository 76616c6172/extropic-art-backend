<template>
  <div class="container-fluid">
    <div class="mt-5">
      <div class="row justify-content-center">
        <div class="col-lg-6 col-sm-12">
          <Typing @set-cursor="setCursor()" />
        </div>
      </div>
      <div class="row justify-content-center">
        <div class="col-lg-6 col-sm-12">
          <Instructions />
        </div>
      </div>
      <div class="row justify-content-center bg-light">
        <div class="col-lg-6 col-sm-12">
          <ItemList />
        </div>
      </div>
      <div class="row justify-content-center bg-light">
        <div class="col-lg-6 col-sm-12">
          <Prompt :showCursor="showCursor" />
        </div>
      </div>
      <div class="row justify-content-center">
        <div class="col-lg-6 col-sm-12">
          <Image />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { default as Typing } from "../components/miscellaneous/TypingComponent.vue";
import { default as Instructions } from "../components/miscellaneous/InstructionsComponent.vue";
import { default as Prompt } from "../components/content/PromptComponent.vue";
import { default as ItemList } from "./content/ItemListComponent.vue";
import { default as Image } from "./content/ImageComponent.vue";

export default {
  name: "HomeComponent",
  components: {
    Typing,
    Instructions,
    Prompt,
    ItemList,
    Image,
  },
  props: {
    msg: String,
  },
  data() {
    return {
      showCursor: false,
    };
  },
  methods: {
    setCursor() {
      this.showCursor = true;
    },
  },
};
</script>
