import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'
import VueRouter from 'vue-router'

import Moniteur from "./components/Moniteur";
import P404 from "./components/404";

import config from './config'
import messages from './translations'

Vue.config.productionTip = false;
Vue.use(VueI18n);
const routes = [
  {path: '/notfound', component: P404},
  {path: '/:id', component: Moniteur},
  {path: '/', redirect: '/notfound'}
];

Vue.use(VueRouter);

const i18n = new VueI18n({
  locale: 'el',
  messages,
});

const router = new VueRouter({
  mode: 'history',
  routes
});

new Vue({
  i18n,
  router,
  data : {
    ...config,
  },
  render: h => h(App),
}).$mount('#app');
