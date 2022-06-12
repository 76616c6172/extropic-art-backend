<template>
  <transition name="fade">
    <div class="popupmodal" v-if="isVisible">
      <div class="popupmodal__window">
        <slot></slot>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "PopupModal",

  data: () => ({
    isVisible: false,
  }),

  methods: {
    open() {
      this.isVisible = true;
    },

    close() {
      this.isVisible = false;
    },
  },
};
</script>
