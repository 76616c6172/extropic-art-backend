<template>
  <div class="container">
    <div class="mt-5">
      <div class="row justify-content-center">
        <div class="col-lg-8 col-sm-12">
          <TypingComponent />
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import TypingComponent from "../components/miscellaneous/TypingComponent.vue";

export default {
  name: "HomeComponent",
  components: {
    TypingComponent,
  },
  props: {
    msg: String,
  },
};
</script>
