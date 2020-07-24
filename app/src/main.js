import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'
import VueRouter from "vue-router";

import messages from './translations'
import config from "./config";

import Moniteur from "./components/Moniteur";
import Page404 from "./components/Page404/Page404";

import * as Sentry from '@sentry/browser';
import { Vue as VueIntegration } from '@sentry/integrations';

Vue.config.productionTip = false;

Vue.use(VueI18n);
Vue.use(VueRouter);

const i18n = new VueI18n({
  locale: config.locale,
  messages,
});

const routes = [
  {
    path: '/:room',
    component: Moniteur
  },
  {
    path: '*',
    component: Page404
  }
];

const router = new VueRouter({
  mode: 'history',
  routes
});

Sentry.init({
  dsn: 'https://564a5fc484f1487aac89ad3490fd2473@o424515.ingest.sentry.io/5357396',
  release: 'frontend@sha',
  integrations: [new VueIntegration({Vue, attachProps: true})],
});

new Vue({
  i18n,
  router,

  render: h => h(App),
}).$mount('#app');
