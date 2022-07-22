<template>
  <nav class="navbar navbar-expand-lg navbar-light">
    <div class="container-fluid p-0">
      <a class="navbar-brand text-start" href="#">Exia</a>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarNav"
        aria-controls="navbarNav"
        aria-expanded="false"
        aria-label="Toggle navigation"
        style="width: auto"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarNav">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li class="nav-item">
            <router-link to="/" class="nav-link text-start active"
              >Disco Diffusion</router-link
            >
          </li>

<!-- Navbar components that aren't implemented yet -->
<!--          <div class="nav-item">                                                               
            <router-link to="/" class="nav-link text-start"
              >Models</router-link
            >
          </div>
          <div class="nav-item">
            <router-link to="/" class="nav-link text-start">Settings</router-link>
          </div>
          <div class="nav-item">
            <router-link to="/" class="nav-link text-start">Login</router-link>
          </div>
-->

        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "NavbarComponent",
};
</script>
