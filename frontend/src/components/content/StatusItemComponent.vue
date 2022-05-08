<template>
  <li
    @click="onClickSetSelected(jobId) & buildModalDialogue()"
    class="list-group-item list-group-item-action"
  >
    <div class="row">
      <div class="col-10">
        <p class="text-start"><strong>Prompt: </strong>{{ prompt }}</p>
      </div>
      <div class="col-2">asdf</div>
    </div>
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
      this.$store.dispatch("getSelectedJob", jobId);
    },
    async buildModalDialogue() {
      console.log("triggered");
      const ok = await this.$refs.confirmDialogue.show({
        title: this.getSelectedJob.prompt,
        image: this.getSelectedJob.img,
        message: "",
        okButton: "Löschen",
      });
      if (ok) {
        // this.$store.dispatch("deleteTodo", "todoObj._id");
      }
    },
  },
  computed: {
    getSelectedJob() {
      return this.$store.getters.getSelectedJob;
    },
  },
};
</script>

<style scoped>
.list-group-item {
  background: transparent;
  padding: 10px 5px;
}
</style>
