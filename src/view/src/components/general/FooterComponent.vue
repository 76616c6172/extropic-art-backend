<template>
  <footer class="footer text-muted mt-auto py-3 bg-light">
    <div class="footer__menu text-center p-3">
      <div class="footer__content">
        <p class="footer__content-text">
          Made with <i class="fa fa-heart" aria-hidden="true"></i> by Valar and
          Zendo in {{ getYear }}
        </p>
      </div>
    </div>
  </footer>
</template>

<script>
export default {
  name: "FooterComponent",
  data() {
    return {};
  },
  computed: {
    getYear() {
      return new Date().getFullYear();
    },
  },
};
</script>
