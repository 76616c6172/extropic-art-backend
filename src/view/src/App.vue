<template>
  <div class="appContainer">
    <div class="mainWrapper row">
      <div class="mainWrapper__col col-12">
        <div class="mainWrapper__container container-fluid">
          <router-view></router-view>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// Components

export default {
  name: "App",
};
</script>

<style lang="scss">
@font-face {
  font-family: "Rubik";
  src: local("Rubik"), url(./fonts/Rubik/Rubik-Regular.ttf) format("truetype");
}
@font-face {
  font-family: "Rubik-Light";
  src: local("Rubik-Light"),
    url(./fonts/Rubik/Rubik-Light.ttf) format("truetype");
}

@import "~bootstrap/scss/bootstrap.scss";
$fa-font-path: "~@fortawesome/fontawesome-free/webfonts";
@import "~@fortawesome/fontawesome-free/scss/fontawesome";
@import "~@fortawesome/fontawesome-free/scss/solid"; // fas
@import "~@fortawesome/fontawesome-free/scss/regular"; // far
@import "~@fortawesome/fontawesome-free/scss/brands"; // fab
@import "./styles/main.css";
#app {
  // font-family: Avenir, Helvetica, Arial, sans-serif;
  font-family: "Rubik", sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
