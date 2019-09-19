import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'
import VueRouter from 'vue-router'
import Inview from 'vueinview'

import Admin from './components/Admin'
import Moniteur from "./components/Moniteur";
import P404 from "./components/404";

import config from './config'

Vue.config.productionTip = false;
Vue.use(VueI18n);
Vue.use(VueRouter);
Vue.use(Inview);

const routes = [
  {path: '/admin', component: Admin},
  {path: '/notfound', component: P404},
  {path: '/:id', component: Moniteur},
  {path: '/', redirect: '/notfound'}
];

const messages = {
  en: {
    message: {
      touchscreen: "The monitor has no touch functionality. Refrain from touching it!",
      examMsg: "Ongoing course examination!",
      nowMsg: "Now ",
      noExamNow: "No course examination right now.",
      noLessonNow: "No lesson right now.",
      nextMsg: "Next ",
      noExamNext: "No course examination with take place the following hours.",
      noLessonNext: "No lesson with take place the following hours.",
      holiday: "Holidays.",
      weekend: "Have a nice weekend!"
    }
  },
  el: {
    message: {
      touchscreen: "Η οθόνη δεν είναι touchscreen. Μην αγγίζετε!",
      examMsg: "Πραγματοποιείται εξέταση!",
      nowMsg: "Τώρα ",
      noExamNow: "Δεν πραγματοποιείται εξέταση.",
      noLessonNow: "Δεν πραγματοποιείται μάθημα.",
      nextMsg: "Επόμενο ",
      noExamNext: "Δεν θα πραγματοποιείται εξέταση.",
      noLessonNext: "Δεν θα πραγματοποιείται μάθημα.",
      holiday: "Αργία ή διακοπές. Καλή ξεκούραση!",
      weekend: "Καλό Σαββατοκύριακο!"
    }
  }
};

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
