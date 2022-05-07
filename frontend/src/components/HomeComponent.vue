<template>
  <div class="container">
    <div class="mt-5">
      <div class="row justify-content-center">
        <div class="col-lg-8 col-sm-12">
          <!-- <Typing /> -->
          <StatusItemList />
          <Prompt />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import { default as Typing } from "../components/miscellaneous/TypingComponent.vue";
import { default as Prompt } from "../components/content/PromptComponent.vue";
import { default as StatusItemList } from "./content/StatusItemListComponent.vue";

export default {
  name: "HomeComponent",
  components: {
    // Typing,
    Prompt,
    StatusItemList,
  },
  props: {
    msg: String,
  },
};
</script>
