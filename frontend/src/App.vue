<template>
  <div class="appContainer">
    <Navbar />
    <Home />
    <Footer />
  </div>
</template>

<script>
// Components
import { default as Home } from "./components/HomeComponent.vue";
import { default as Navbar } from "./components/general/NavbarComponent.vue";
import { default as Footer } from "./components/general/FooterComponent.vue";

export default {
  name: "App",
  components: {
    Home,
    Navbar,
    Footer,
  },
};
</script>

<style lang="scss">
@import "~bootstrap/scss/bootstrap.scss";
$fa-font-path: "~@fortawesome/fontawesome-free/webfonts";
@import "~@fortawesome/fontawesome-free/scss/fontawesome";
@import "~@fortawesome/fontawesome-free/scss/solid"; // fas
@import "~@fortawesome/fontawesome-free/scss/regular"; // far
@import "~@fortawesome/fontawesome-free/scss/brands"; // fab
@import "./styles/main.css";
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
