<template>
  <div class="container">
    <div class="mt-5">
      <div class="row justify-content-center">
        <div class="col-lg-8 col-sm-12">
          <Typing />
        </div>
      </div>
      <div class="row justify-content-center">
        <div class="col-lg-8 col-sm-12">
          <Instructions />
        </div>
      </div>
      <div class="row justify-content-center">
        <div class="col-lg-8 col-sm-12">
          <StatusItemList />
        </div>
      </div>
      <div class="row justify-content-center">
        <div class="col-lg-8 col-sm-12">
          <Prompt />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { default as Typing } from "../components/miscellaneous/TypingComponent.vue";
import { default as Instructions } from "../components/miscellaneous/InstructionsComponent.vue";
import { default as Prompt } from "../components/content/PromptComponent.vue";
import { default as StatusItemList } from "./content/StatusItemListComponent.vue";

export default {
  name: "HomeComponent",
  components: {
    Typing,
    Instructions,
    Prompt,
    StatusItemList,
  },
  props: {
    msg: String,
  },
};
</script>
