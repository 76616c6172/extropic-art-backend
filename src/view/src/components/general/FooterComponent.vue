<template>
  <footer class="footer text-muted mt-auto py-3 bg-light">
    <div class="text-center p-3">
      Made with <i class="fa fa-heart" aria-hidden="true"></i> by valar and
      zen in {{ getYear }}
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
