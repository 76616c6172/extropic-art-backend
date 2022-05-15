<template>
  <div>
    <div class="mb-5 mt-5">
      <!-- <div v-if="isLoading" class="spinner-border text-secondary" role="status">
        <span class="visually-hidden">Loading...</span>
      </div> -->
      <div class="imgContainer">
        <div v-if="isLoading" class="loader">Loading...</div>
        <img :src="imgObjectURL" class="img-fluid img-thumbnail" alt="" />
      </div>
    </div>
  </div>
</template>

<script>
export default {
  name: "ImageComponent",
  data() {
    return {
      imgObjectURL:
        "https://via.placeholder.com/1920x1024.png?text=This%20is%20zen%27s%20placeholder",
      isLoading: false,
    };
  },
  methods: {
    createImgObjectURL() {
      this.isLoading = true;
      this.imgObjectURL =
        "https://via.placeholder.com/1920x1024.png?text=Loading%20image";
      this.$store
        .dispatch("getSelectedImg")
        .then((response) => {
          var img = new Image();
          let blob = new Blob([response], { type: "image/png" });
          let url = URL.createObjectURL(blob);
          img.src = url;
          this.imgObjectURL = img.src;
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
  computed: {
    getSelectedJob() {
      return this.$store.getters.getSelectedJob;
    },
  },
  watch: {
    getSelectedJob: {
      handler() {
        this.createImgObjectURL();
      },
    },
  },
};
</script>

<style>
.imgContainer {
  position: relative;
}

.loader,
.loader:before,
.loader:after {
  border-radius: 50%;
  width: 2.5em;
  height: 2.5em;
  -webkit-animation-fill-mode: both;
  animation-fill-mode: both;
  -webkit-animation: load7 1.8s infinite ease-in-out;
  animation: load7 1.8s infinite ease-in-out;
}
.loader {
  position: absolute;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  color: #327959;
  font-size: 10px;
  margin: 80px auto;
  text-indent: -9999em;
  -webkit-transform: translateZ(0);
  -ms-transform: translateZ(0);
  transform: translateZ(0);
  -webkit-animation-delay: -0.16s;
  animation-delay: -0.16s;
}
.loader:before,
.loader:after {
  content: "";
  position: absolute;
  top: 0;
}
.loader:before {
  left: -3.5em;
  -webkit-animation-delay: -0.32s;
  animation-delay: -0.32s;
}
.loader:after {
  left: 3.5em;
}
@-webkit-keyframes load7 {
  0%,
  80%,
  100% {
    box-shadow: 0 2.5em 0 -1.3em;
  }
  40% {
    box-shadow: 0 2.5em 0 0;
  }
}
@keyframes load7 {
  0%,
  80%,
  100% {
    box-shadow: 0 2.5em 0 -1.3em;
  }
  40% {
    box-shadow: 0 2.5em 0 0;
  }
}
</style>
