<template>
  <popup-modal ref="popup">
    <h2 class="popupmodal__title" style="margin-top: 0" v-html="title"></h2>
    <p class="popupmodal__message" v-html="message"></p>
    <img class="img-fluid" :src="`${image}`" alt="" />
    <div class="popupmodal__buttons">
      <button class="popupmodal__buttons-cancel" @click="_cancel">
        {{ cancelButton }}
      </button>
      <button class="popupmodal__buttons-ok" @click="_confirm">
        {{ okButton }}
      </button>
    </div>
  </popup-modal>
</template>

<script>
import PopupModal from "./PopupModal";
export default {
  name: "ConfirmDialogue",

  components: { PopupModal },

  data: () => ({
    // Parameters that change depending on the type of dialogue
    title: undefined,
    message: undefined, // Main text content
    image: undefined, // Main image to display
    okButton: undefined, // Text for confirm button; leave it empty because we don't know what we're using it for
    cancelButton: "Cancel", // text for cancel button

    // Private variables
    resolvePromise: undefined,
    rejectPromise: undefined,
  }),

  methods: {
    show(opts = {}) {
      this.title = opts.title;
      this.message = opts.message;
      this.image = opts.image;
      this.okButton = opts.okButton;
      if (opts.cancelButton) {
        this.cancelButton = opts.cancelButton;
      }
      // Once we set our config, we tell the popup modal to open
      this.$refs.popup.open();
      // Return promise so the caller can get results
      return new Promise((resolve, reject) => {
        this.resolvePromise = resolve;
        this.rejectPromise = reject;
      });
    },

    _confirm() {
      this.$refs.popup.close();
      this.resolvePromise(true);
    },

    _cancel() {
      this.$refs.popup.close();
      this.resolvePromise(false);
      // Or you can throw an error
      // this.rejectPromise(new Error('User cancelled the dialogue'))
    },
  },
};
</script>

<style></style>
