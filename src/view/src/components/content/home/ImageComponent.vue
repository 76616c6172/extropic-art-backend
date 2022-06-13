<template>
  <div>
    <div class="imgContainer">
      <div class="imgRendered">
        <div v-if="isLoading" class="loader">Loading...</div>
        <img :src="imgObjectURL" class="img-fluid img-thumbnail" alt="" />
      </div>
      <div v-show="Object.keys(getSelectedJob).length !== 0" class="imgTextbox">
        <div class="p-5 imgTextboxContent">
          <hr />
          <p class="fs-4">
            {{ getSelectedJob.prompt }}
          </p>
          <hr />
          <span>
            <i class="fa-solid fa-link"></i>
            <a :href="imgObjectURL" class="text-white" target="_blank">
              Goto Image</a
            >
          </span>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
import { mapGetters } from "vuex";
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
          this.imgObjectURL = response;
        })
        .finally(() => {
          this.isLoading = false;
        });
    },
  },
  computed: {
    ...mapGetters(["getSelectedJob"]),
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
.imgContainer:hover .imgTextbox {
  opacity: 1;
}

.imgRendered {
  z-index: 0;
}
.imgRendered img {
  display: block;
  position: relative;
  width: 100%;
  height: auto;
}

.imgTextbox {
  position: absolute;
  top: 0px;
  left: 0px;
  right: 0px;
  bottom: 0px;
  z-index: 1;
  opacity: 0;
  background-color: rgba(0, 0, 0, 0.6);
  width: 100%;
  height: 100%;
  transition: 0.4s ease-in-out;
  display: block;
  transform: scale(1);
  backface-visibility: hidden;
  color: #f2f2f2;
}

.imgTextboxContent {
  padding: 5px;
  top: 50%;
  left: 50%;
  float: left;
  display: block;
  position: relative;
  transform: translateX(-50%) translateY(-50%);
}

/* Loader Start */
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
  z-index: 1;
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

/* Loader End */
</style>
