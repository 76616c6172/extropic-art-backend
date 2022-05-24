<template>
  <div class="rounded terminalWrapper">
    <div class="p-4 border rounded">
      <div class="ui-container">
        <div class="ui-circle ui-circle-red rounded-pill"></div>
        <div class="ui-circle ui-circle-green rounded-pill"></div>
      </div>
      <ItemList />
      <Prompt :showCursor="showCursor" />
    </div>
  </div>
</template>

<script>
import { default as Prompt } from "./PromptComponent.vue";
import { default as ItemList } from "./ItemListComponent.vue";
export default {
  name: "TerminalWrapperComponent",
  components: {
    Prompt,
    ItemList,
  },
  props: {
    showCursor: {
      type: Boolean,
      required: true,
    },
  },
};
</script>

<style>
.terminalWrapper {
  background: #191724;
}

.ui-container {
  display: flex;
}

.ui-circle {
  width: 15px;
  height: 15px;
  margin-left: 10px;
}

.ui-circle-red {
  background: #ff5f56;
}

.ui-circle-green {
  background: #27c93f;
}
</style>
