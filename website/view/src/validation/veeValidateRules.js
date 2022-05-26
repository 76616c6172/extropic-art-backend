import { defineRule } from "vee-validate";

defineRule("required", (value) => {
  if (!value || !value.length) {
    return "Required field";
  }
  return true;
});

defineRule("minLength", (value, [limit]) => {
  if (value.length < limit) {
    return `Minimum length: ${limit}`;
  }
  return true;
});

defineRule("maxLength", (value, [limit]) => {
  if (value.length > limit) {
    return `Maximum length: ${limit}`;
  }
  return true;
});

defineRule("selectValue", (value) => {
  if (value == undefined) {
    return "Select value from dropdown";
  }
  return true;
});

defineRule("email", (value) => {
  const checkEmailRegex = new RegExp(
    /^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9-]+(?:\.[a-zA-Z0-9-]+)*$/
  );
  if (!checkEmailRegex.test(value)) {
    return "Email address";
  }
  return true;
});

defineRule("noWhitespace", (value) => {
  const checkNoWhitespaceRegex = new RegExp(/([\S]+)/g);
  if (!checkNoWhitespaceRegex.test(value)) {
    return "No whitespaces";
  }
  return true;
});

defineRule("commaSeperated", (value) => {
  const commaSeperatedRegex = new RegExp(/^([\S^,]+(,+)\S){1}/g);
  if (!commaSeperatedRegex.test(value)) {
    return "Format: prompt1,prompt2,... (2 prompts min)";
  }
  return true;
});

defineRule("confirmed", (value, [other]) => {
  if (value !== other) {
    return `Passwords do not match`;
  }
  return true;
});
