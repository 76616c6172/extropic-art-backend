<template>
  <div class="terminalWrapper rounded-3 shadow border border-4 border-white">
    <div class="terminalWrapperBg p-4 rounded-3">
      <div class="ui-container">
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
}

.terminalWrapperBg {
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

.ui-circle-yellow {
  background: #ffbd2e;
}
</style>
