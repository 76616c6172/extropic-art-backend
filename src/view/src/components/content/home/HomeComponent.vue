<template>
  <Typing @set-cursor="setCursor()" />
  <Instructions />
  <ImageNewestRendersComponent
    v-if="jobStatus.newestCompletedJobs"
    :newestJobIds="jobStatus.newestCompletedJobs"
  />
  <StatsComponent v-if="jobStatus" :jobStatus="jobStatus" />
  <TerminalWrapper :showCursor="showCursor" />
  <Image />
</template>

<script>
import { mapGetters } from "vuex";
import { default as Typing } from "./TypingComponent.vue";
import { default as ImageNewestRendersComponent } from "./ImageNewestRendersComponent";
import { default as Image } from "./ImageComponent.vue";
import { default as TerminalWrapper } from "./TerminalWrapperComponent.vue";
import { default as Instructions } from "./InstructionsComponent.vue";
import { default as StatsComponent } from "./StatsComponent.vue";

export default {
  name: "HomeComponent",
  components: {
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
