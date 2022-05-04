<template>
  <div class="container">
    <div class="mt-5">
      <div class="row justify-content-center">
        <div class="col-lg-8 col-sm-12">
          <!-- <TypingComponent /> -->
          <StatusListComponent />
          <!-- <PromptComponent /> -->
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// import TypingComponent from "../components/miscellaneous/TypingComponent.vue";
// import PromptComponent from "../components/content/PromptComponent.vue";
import StatusListComponent from "../components/content/StatusListComponent.vue";

export default {
  name: "HomeComponent",
  components: {
    // TypingComponent,
    // PromptComponent,
    StatusListComponent,
  },
  props: {
    msg: String,
  },
};
</script>
