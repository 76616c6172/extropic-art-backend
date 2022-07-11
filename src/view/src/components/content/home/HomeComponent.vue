<template>
  <div class="container-fluid">
    <div class="row justify-content-center">
      <div class="col-lg-10 col-sm-12">
        <Navbar />
      </div>
    </div>
    <div class="mt-5">
      <div class="row justify-content-center">
        <div class="col-lg-10 col-sm-12">
          <Typing @set-cursor="setCursor()" />
        </div>
      </div>
      <div class="row justify-content-center pb-5">
        <div class="col-lg-10 col-sm-12">
          <Instructions />
        </div>
      </div>
      <div class="row justify-content-center">
        <div class="col-lg-10 col-sm-12">
          <ImageNewestRendersComponent
            v-if="jobStatus.newestCompletedJobs"
            :newestJobIds="jobStatus.newestCompletedJobs"
          />
        </div>
      </div>
      <div class="row justify-content-center pb-5 pt-6">
        <div class="col-lg-10 col-sm-12">
          <StatsComponent v-if="jobStatus" :jobStatus="jobStatus" />
        </div>
      </div>
      <div class="row justify-content-center pt-5 pb-6">
        <div class="col-lg-10 col-sm-12">
          <Image />
        </div>
      </div>
      <div class="row justify-content-center bg-light pt-7 pb-7">
        <div class="col-lg-10 col-sm-12">
          <TerminalWrapper :showCursor="showCursor" />
        </div>
      </div>
    </div>
  </div>
  <Footer />
</template>

<script>
document.title = 'Exia'

import { mapGetters } from "vuex";
import { default as Navbar } from "../../general/NavbarComponent.vue";
import { default as Footer } from "../../general/FooterComponent.vue";
import { default as Typing } from "./TypingComponent.vue";
import { default as ImageNewestRendersComponent } from "./ImageNewestRendersComponent";
import { default as Image } from "./ImageComponent.vue";
import { default as TerminalWrapper } from "./TerminalWrapperComponent.vue";
import { default as Instructions } from "./InstructionsComponent.vue";
import { default as StatsComponent } from "./StatsComponent.vue";

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
    ImageNewestRendersComponent,
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
  },
  computed: {
    ...mapGetters(["getJobStatus"]),
  },
  async mounted() {
    this.$store.dispatch("fetchJobStatus", "initial").then(() => {
      this.jobStatus = this.getJobStatus;
    });
  },
  watch: {
    getJobStatus: {
      handler() {
        this.jobStatus = this.getJobStatus;
      },
    },
  },
};
</script>
