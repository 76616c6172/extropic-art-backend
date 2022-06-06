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
          <StatsComponent v-if="jobStatus" :jobStatus="jobStatus" />
        </div>
      </div>
      <div class="row justify-content-center bg-light pt-5 pb-5">
        <div class="col-lg-6 col-sm-12">
          <TerminalWrapper
            @set-newJob="reloadJobStatus()"
            :showCursor="showCursor"
          />
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
import { mapGetters } from "vuex";
import { default as Navbar } from "../components/general/NavbarComponent.vue";
import { default as Footer } from "../components/general/FooterComponent.vue";
import { default as Typing } from "../components/miscellaneous/TypingComponent.vue";
import { default as Image } from "./content/ImageComponent.vue";
import { default as TerminalWrapper } from "./content/TerminalWrapperComponent.vue";
import { default as Instructions } from "../components/miscellaneous/InstructionsComponent.vue";
import { default as StatsComponent } from "../components/miscellaneous/StatsComponent.vue";

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
      jobStatus: {},
    };
  },
  methods: {
    setCursor() {
      this.showCursor = true;
    },
    reloadJobStatus() {
      this.$store.dispatch("fetchJobStatus", "initial").then(() => {
        this.jobStatus = this.getJobStatus;
      });
    },
  },
  computed: {
    ...mapGetters(["getJobStatus"]),
  },
  async mounted() {
    this.$store.dispatch("fetchJobStatus", "initial").then(() => {
      this.jobStatus = this.getJobStatus;
    });
  },
};
</script>
