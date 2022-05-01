<template>
  <nav class="navbar navbar-expand-lg navbar-dark fixed-top bg-dark">
    <div class="container">
      <a class="navbar-brand" href="#">Exia</a>
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
        <ul class="navbar-nav">
          <li class="nav-item">
            <router-link to="/" class="nav-link text-start active">Home</router-link>
          </li>
          <div class="nav-item">
            <router-link to="/" class="nav-link text-start">Settings</router-link>
          </div>
          <div class="nav-item">
            <router-link to="/" class="nav-link text-start">About</router-link>
          </div>
          <div class="nav-item">
            <router-link to="/" class="nav-link text-start">Login</router-link>
          </div>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  name: "NavbarComponent"
}
</script>