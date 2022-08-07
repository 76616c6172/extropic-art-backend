<template>
  <div class="row justify-content-center">
    <div class="col-lg-10 col-sm-12">
      <div class="row">
        <!--{{ $store.getters.getJobs }}-->
        <div
          v-for="(job, index) in imgArray"
          :key="index"
          class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1"
        >
          <figure class="image-block">
            <img
              @click="onClickOpenFullImage(job)"
              :src="job.imgURL"
              class="img-fluid img-thumbnail"
              alt=""
            />
            <figcaption>
              <p class="prompt text-white">{{ job.prompt }}</p>
            </figcaption>
          </figure>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ImageNewestRendersComponent",
  data() {
    return {
      imgArray: [],
    };
  },
  props: ["newestJobIds"],
  methods: {
    onClickOpenFullImage(job) {
      let routeData = this.$router.resolve({
        name: "JobDetails",
        params: { jobId: job.jobid },
      });
      window.open(routeData.href, "_blank");
    },
    createImgObjectURL(jobId) {
      return new Promise((resolve) => {
        this.$store
          .dispatch("getSelectedImg", { jobId: jobId, type: "thumbnail" })
          .then((responseImg) => {
            this.$store.getters.getJobs.map((job) => {
              if (job.jobid == jobId) {
                this.imgArray.push({
                  jobid: jobId,
                  imgURL: responseImg,
                  prompt: job.prompt,
                });
              }
            });
          })
          .finally(() => {
            resolve();
          });
      });
    },
  },
  computed: {
    getIsInitialLoadStatus() {
      return this.$store.getters.getIsInitialLoadStatus;
    },
  },
  async mounted() {
    if (this.imgArray.length == 0) {
      this.newestJobIds.map((jobId) => {
        this.createImgObjectURL(jobId);
      });
    }
  },
};
</script>

<style scoped>
p.prompt {
  white-space: nowrap;
  width: 100%;
  overflow: hidden;
  -o-text-overflow: ellipsis;
  text-overflow: ellipsis;
}
/* Image Hover */
figure {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  width: 100%;
  height: auto;
  box-shadow: 0 1px 3px rgba(0, 0, 0, 0.12), 0 1px 2px rgba(0, 0, 0, 0.24);
  transition: all 0.3s cubic-bezier(0.25, 0.8, 0.25, 1);
  overflow: hidden;
}
figure:hover {
  box-shadow: 0 14px 28px rgba(0, 0, 0, 0.25), 0 10px 10px rgba(0, 0, 0, 0.22);
}
figure:hover img {
  transform: scale(1.25);
}
figure:hover figcaption {
  bottom: -25px;
}
figure:hover figcaption p {
  white-space: auto;
}
figure h1 {
  position: absolute;
  top: 50px;
  left: 20px;
  margin: 0;
  padding: 0;
  color: white;
  font-size: 60px;
  font-weight: 100;
  line-height: 1;
}
figure img {
  cursor: pointer;
  height: 100%;
  transition: 0.25s;
}
figure figcaption {
  position: absolute;
  bottom: -90px;
  left: 0;
  width: 100%;
  margin: 0;
  padding: 5px 30px 15px;
  background-color: rgba(0, 0, 0, 0.8);
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.4);
  line-height: 1;
  transition: 0.25s;
}
figure figcaption h3 {
  margin: 0 0 10px;
  padding: 0;
}
figure figcaption p {
  font-size: 14px;
  line-height: 1.75;
}
figure figcaption button {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 10px 0 0;
  border: 1px solid rgb(44, 62, 80);
}

figcaption button i {
  color: rgb(44, 62, 80);
}

figure figcaption button a {
  text-decoration: none;
  color: rgb(44, 62, 80);
}
</style>
