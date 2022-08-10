const { defineConfig } = require("@vue/cli-service");
const path = require("path");

module.exports = defineConfig({
  transpileDependencies: true,
  chainWebpack: (config) => {
    const types = ["vue-modules", "vue", "normal-modules", "normal"];
    types.forEach((type) =>
      addStyleResource(config.module.rule("scss").oneOf(type))
    );
  },
});

/* SCSS support (https://vinceumo.github.io/devNotes/Javascript/vue-scss-setup/) */
function addStyleResource(rule) {
  rule
    .use("style-resource")
    .loader("style-resources-loader")
    .options({
      patterns: [
        path.resolve(__dirname, "./src/styles/content/*.scss"),
        path.resolve(__dirname, "./src/styles/general/*.scss"),
        path.resolve(__dirname, "./src/styles/variables/*.scss"),
      ],
    });
}
