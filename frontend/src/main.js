import { createApp } from 'vue';
import * as bsMin from 'bootstrap/dist/js/bootstrap.min.js';
import router from './router';
import App from './App.vue';


let vueApp = createApp(App);

vueApp.use(bsMin);
vueApp.use(router);
vueApp.mount('#app');

