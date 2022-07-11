<template>
  <footer class="footer text-default mt-auto py-2">
    <div class="text-center p-10">
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
