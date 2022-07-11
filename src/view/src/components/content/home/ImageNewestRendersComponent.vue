<template>
<div>
    <h2 class="text-start display-7">Featured Model: Disco Diffusion 5.3</h2>
<div><br></div>

  <div><br><br></div>
    <div class="row">
      <div
        v-for="(job, index) in imgArray"
        :key="index"
        class="col-xxl-4 col-xl-4 col-lg-4 col-md-4 col-sm-1 col-xs-1"
      >
        <figure class="image-block">
          <img :src="job.imgURL" class="img-fluid img-thumbnail" alt="" />
          <figcaption>
            <h4 class="h4">^</h4>
            <p class="prompt">{{ job.prompt }}</p>
            <button class="btn text-center">
              <a :href="`${job.imgURL}`" target="_blank"
                ><i class="fa fa-eye"></i> Full image</a
              >
            </button>
          </figcaption>
        </figure>
      </div>
    </div>
  </div>
  <div><br></div>
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
    createImgObjectURL(jobId) {
      return new Promise((resolve) => {
        this.$store
          .dispatch("getSelectedImg", jobId)
          .then((responseImg) => {
            this.imgArray.push({
              jobid: jobId,
              imgURL: responseImg,
            });
          })
          .finally(() => {
            resolve();
          });
      });
    },
    async getSelectedJobsObject(jobId) {
      await this.createImgObjectURL(jobId).then(() => {
        if (this.imgArray.length == 3) {
          let newestJobIdsValues = Object.values(this.newestJobIds);
          let newestJobIdsMax = Math.max(...newestJobIdsValues);
          let newestJobIdsMin = Math.min(...newestJobIdsValues);
          this.$store
            .dispatch("getSelectedJobs", {
              jobx: newestJobIdsMin,
              joby: newestJobIdsMax,
              jobIds: newestJobIdsValues,
            })
            .then(() => {
              this.$store.getters.getSelectedJobs.forEach((storeJobElement) => {
                this.imgArray.forEach((imgArrayElement) => {
                  if (imgArrayElement.jobid == storeJobElement.jobid) {
                    imgArrayElement.prompt = storeJobElement.prompt;
                  }
                });
              });
            });
        }
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
      this.newestJobIds.map((jobId, index) =>
        index <= 3 ? this.getSelectedJobsObject(jobId) : ""
      );
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
  white-space: wrap;
  bottom: 0;
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
  height: 100%;
  transition: 0.25s;
}
figure figcaption {
  position: absolute;
  bottom: -90px;
  left: 0;
  width: 100%;
  margin: 0;
  padding: 0px 5px 0px 0px;
  background-color: rgba(25, 23, 36, 0.85);
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.4);
  line-height: 1;
  transition: 0.25s;
}
figure figcaption h4 {
  margin: 0 0 10px;
  padding: 0;
  color: white;
}
figure figcaption p {
  font-size: 14px;
  line-height: 1.75;
  color: white;
}
figure figcaption button {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: center;
  margin: 0px 0 0;
  border: 0px solid rgb(44, 62, 80);
}

figcaption button i {
  color: white;
}

figure figcaption button a {
  text-decoration: none;
  color: white;
}
</style>
