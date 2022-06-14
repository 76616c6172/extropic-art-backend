<template>
  <div class="navBar__container row justify-content-center">
    <div class="navBar__container-inner col-lg-10 col-sm-12">
      <header>
        <nav class="navBar__container-nav navbar navbar-expand-lg navbar-light">
          <div class="navBar__list container-fluid p-0">
            <a class="navBar__list-link navbar-brand text-start" href="#"
              >Exia</a
            >
            <button
              class="navBar__list-btn navbar-toggler"
              type="button"
              data-bs-toggle="collapse"
              data-bs-target="#navbarNav"
              aria-controls="navbarNav"
              aria-expanded="false"
              aria-label="Toggle navigation"
              style="width: auto"
            >
              <span class="navBar__list-toggleSpan navbar-toggler-icon"></span>
            </button>
            <div class="navBar__listContainer navbar-collapse" id="navbarNav">
              <ul class="navigation__list navbar-nav me-auto mb-2 mb-lg-0">
                <li class="navigation__list-item nav-item">
                  <router-link to="/" class="nav-link text-start active"
                    >Home</router-link
                  >
                </li>
                <div class="navigation__list-item nav-item">
                  <router-link
                    to="/"
                    class="navigation__list-link nav-link text-start"
                    >Settings</router-link
                  >
                </div>
                <div class="navigation__list-item nav-item">
                  <router-link
                    to="/"
                    class="navigation__list-link nav-link text-start"
                    >About</router-link
                  >
                </div>
                <div class="navigation__list-item nav-item">
                  <router-link
                    to="/"
                    class="navigation__list-link nav-link text-start"
                    >Login</router-link
                  >
                </div>
              </ul>
            </div>
          </div>
        </nav>
      </header>
    </div>
  </div>
</template>

<script>
export default {
  name: "NavbarComponent",
};
</script>
