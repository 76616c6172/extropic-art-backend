<template>
  <div>
    <div class="row">
      <Sidebar :job="{ job }" />
      <FullSizeImage :imgObjectURL="imgObjectURL" />
    </div>
  </div>
</template>

<script>
import { default as Sidebar } from "./SidebarComponent.vue";
import { default as FullSizeImage } from "./FullSizeImageComponent.vue";

export default {
  name: "FullSizeImageWrapperComponent",
  components: {
    Sidebar,
    FullSizeImage,
  },
  data() {
    return {
      job: "",
      jobId: "",
      imgObjectURL: "",
    };
  },
  methods: {
    createImgObjectURL(jobId) {
      this.$store
        .dispatch("getSelectedImg", { jobId: jobId, type: "full" })
        .then((response) => {
          this.imgObjectURL = response;
        });
    },
  },
  async mounted() {
    this.jobId = this.$route.params.jobId;
    this.createImgObjectURL(this.jobId);
    this.$store.dispatch("fetchJob", this.jobId).then(() => {
      this.job = this.$store.getters.getSelectedJob;
      console.log(this.job);
    });
  },
};
</script>

<style></style>
