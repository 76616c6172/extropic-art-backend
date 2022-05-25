<template>
  <VeeForm v-slot="{ handleSubmit }" as="div" ref="promptForm">
    <form @submit="handleSubmit($event, onSubmit)" name="promptForm">
      <div v-if="showCursor" class="col-12 mb-3 mb-lg-0">
        <div class="input-group">
          <i class="fa fa-arrow-right" aria-hidden="true"></i>
          <div class="w-100" ref="inputPrompt">
            <Field
              as="input"
              rules="required|commaSeperated|minLength:1|maxLength:600"
              name="vPrompt"
              type="input"
              class="form-control bg-transparent text-white"
            />
          </div>
        </div>
        <div class="input-group">
          <ErrorMessage
            name="vPrompt"
            as="div"
            class="alertbox text-start text-white"
            role="alert"
          />
        </div>
        <div v-if="promptStatus.length !== 0" class="input-group">
          <div class="alertbox text-start text-white">{{ promptStatus }}</div>
        </div>
      </div>
    </form>
  </VeeForm>
</template>

<script>
import * as Validation from "../../validation/veeValidateRules";
import { Form as VeeForm, Field, ErrorMessage } from "vee-validate";
export default {
  name: "PromptComponent",
  components: {
    VeeForm,
    Field,
    ErrorMessage,
  },
  data() {
    return {
      promptStatus: "",
      Validation,
    };
  },
  props: {
    showCursor: {
      type: Boolean,
      required: true,
    },
  },
  methods: {
    onSubmit(values, { resetForm }) {
      resetForm();
      this.setPromptStatus();
      this.$store.dispatch("sendNewJob", { prompt: values.vPrompt });
    },
    setPromptStatus() {
      this.promptStatus = "Your prompt is being processed...";
      setTimeout(() => {
        this.promptStatus = "";
      }, 3000);
    },
  },
  computed: {
    setAutofocus() {
      return this.showCursor ? this.$refs.inputPrompt.firstChild.focus() : "";
    },
  },
  watch: {
    showCursor: {
      handler() {
        setTimeout(() => {
          this.setAutofocus;
        }, 500);
      },
    },
  },
};
</script>

<style scoped>
.alertbox {
  padding: 10px 0;
}

.input-group i {
  position: absolute;
  top: 10px;
  left: 10px;
  z-index: 4;
}

.input-group input {
  padding-left: 30px;
  border: none;
}
</style>
