import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'
import VueRouter from "vue-router";

import messages from './translations'

import Moniteur from "./components/Moniteur";

Vue.config.productionTip = false;

Vue.use(VueI18n);
Vue.use(VueRouter);

const i18n = new VueI18n({
  locale: 'el',
  messages,
});

const routes = [
  {
    path: '/:room',
    component: Moniteur
  }
];

const router = new VueRouter({
  mode: 'history',
  routes
});

new Vue({
  i18n,
  router,

  render: h => h(App),
}).$mount('#app');
