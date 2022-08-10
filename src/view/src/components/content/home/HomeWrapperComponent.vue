<template>
  <div>
    <NavbarComponent />
    <div class="mt-5 mb-5">
      <Typing @set-cursor="setCursor()" />
      <Instructions />
      <ImageNewestRendersComponent
        v-if="jobStatus.newestCompletedJobs"
        :newestJobIds="jobStatus.newestCompletedJobs"
      />
      <StatsComponent v-if="jobStatus" :jobStatus="jobStatus" />
      <TerminalWrapper :showCursor="showCursor" />
    </div>
    <FooterComponent />
  </div>
</template>

<script>
import { mapGetters } from "vuex";
import { default as Typing } from "./TypingComponent.vue";
import { default as ImageNewestRendersComponent } from "./ImageNewestRendersComponent";
import { default as TerminalWrapper } from "./terminal/TerminalWrapperComponent.vue";
import { default as Instructions } from "./InstructionsComponent.vue";
import { default as StatsComponent } from "./StatsComponent.vue";
import { default as NavbarComponent } from "../../general/NavbarComponent.vue";
import { default as FooterComponent } from "../../general/FooterComponent.vue";

export default {
  name: "HomeComponent",
  components: {
    Typing,
    Instructions,
    TerminalWrapper,
    StatsComponent,
    ImageNewestRendersComponent,
    NavbarComponent,
    FooterComponent,
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
