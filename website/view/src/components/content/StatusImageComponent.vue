<template>
  <div>
    <div @click="buildModalDialogue()">Open Modal</div>
    <ConfirmDialogue ref="confirmDialogue" />
    <p>{{ getSelectedPrompt }}</p>
    <img :src="`${getSelectedImageUrl}`" alt="" />
  </div>
</template>

<script>
import ConfirmDialogue from "../modal/Confirmdialogue.vue";

export default {
  name: "StatusImageComponent",
  components: {
    ConfirmDialogue,
  },
  methods: {
    async buildModalDialogue() {
      console.log(this.getSelectedImageUrl);
      console.log("triggered");
      const ok = await this.$refs.confirmDialogue.show({
        title: this.getSelectedPrompt,
        image: this.getSelectedImageUrl,
        message: "",
        okButton: "Löschen",
      });
      if (ok) {
        // this.$store.dispatch("deleteTodo", "todoObj._id");
      }
    },
  },
  computed: {
    getSelectedImageUrl() {
      return this.$store.getters.getSelectedImageUrl;
    },
    getSelectedPrompt() {
      return this.$store.getters.getSelectedPrompt;
    },
  },
};
</script>

<style></style>
