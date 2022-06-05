<template>
  <div class="container-fluid">
    <div class="row justify-content-center">
      <div class="col-lg-6 col-sm-12">
        <Navbar />
      </div>
    </div>
    <div class="mt-5">
      <div class="row justify-content-center">
        <div class="col-lg-6 col-sm-12">
          <Typing @set-cursor="setCursor()" />
        </div>
      </div>
      <div class="row justify-content-center pb-5">
        <div class="col-lg-6 col-sm-12">
          <Instructions />
        </div>
      </div>
      <div class="row justify-content-center pb-5">
        <div class="col-lg-6 col-sm-12">
          <StatsComponent />
        </div>
      </div>
      <div class="row justify-content-center bg-light pt-5 pb-5">
        <div class="col-lg-6 col-sm-12">
          <TerminalWrapper :showCursor="showCursor" />
        </div>
      </div>
      <div class="row justify-content-center pt-5 pb-5">
        <div class="col-lg-6 col-sm-12">
          <Image />
        </div>
      </div>
    </div>
  </div>
  <Footer />
</template>

<script>
import { default as Navbar } from "../components/general/NavbarComponent.vue";
import { default as Footer } from "../components/general/FooterComponent.vue";
import { default as Typing } from "../components/miscellaneous/TypingComponent.vue";
import { default as Instructions } from "../components/miscellaneous/InstructionsComponent.vue";
import { default as Image } from "./content/ImageComponent.vue";
import { default as TerminalWrapper } from "./content/TerminalWrapperComponent.vue";
import { default as StatsComponent } from "./content/StatsComponent.vue";

export default {
  name: "HomeComponent",
  components: {
    Navbar,
    Footer,
    Typing,
    Instructions,
    Image,
    TerminalWrapper,
    StatsComponent,
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
