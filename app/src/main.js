import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'

import config from './config'

Vue.config.productionTip = false;
Vue.use(VueI18n);

const room = () => {
  let url =  new URL(window.location.href.toString());
  return url.searchParams.get('room');
}

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

new Vue({
  i18n,
  data : {
    ...config,
    room: room(),
  },
  render: h => h(App),
}).$mount('#app');
