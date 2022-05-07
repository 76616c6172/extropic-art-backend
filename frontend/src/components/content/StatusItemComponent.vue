<template>
  <li
    @click="onClickSetSelected(jobId) & buildModalDialogue()"
    style="list-style-type: none; background-color: lightgrey"
  >
    <p class="text-start"><strong>Prompt: </strong>{{ prompt }}</p>
    <p class="text-start">(<strong>JobID: </strong> {{ jobId }})</p>
  </li>
  <ConfirmDialogue ref="confirmDialogue" />
</template>

<script>
import ConfirmDialogue from "../modal/Confirmdialogue.vue";
export default {
  name: "StatusItemComponent",
  components: {
    ConfirmDialogue,
  },
  props: {
    prompt: {
      type: String,
      required: true,
    },
    jobId: {
      type: String,
      required: true,
    },
  },
  methods: {
    onClickSetSelected(jobId) {
      this.$store.dispatch("fetchSelectedImage", jobId);
      this.$store.dispatch("fetchSelectedPrompt", jobId);
    },
    async buildModalDialogue() {
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
