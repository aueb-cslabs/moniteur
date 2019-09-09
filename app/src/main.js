import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'

Vue.config.productionTip = false;
Vue.use(VueI18n);

const shared = {
  getRoom() {
    let link =  window.location.href.toString();
    const room = link.substring(link.lastIndexOf("/")+1);
    return room;
  }
};

const messages = {
  en: {
    message: {
      touchscreen: "The monitor has no touch functionality. Don't touch!",
      examMsg: "Examination!",
      nowMsg: "Now ",
      noExamNow: "No examination right now.",
      noLessonNow: "No lesson carried out right now.",
      nextMsg: "Next ",
      noExamNext: "No examination after.",
      noLessonNext: "No lesson carried out afterwards.",
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
    room: shared.getRoom(),
    api: "http://localhost",
    port: ":27522",
    LOGO_LINK: 'https://www.aueb.gr/press/logos/1_AUEB-pantone-HR.jpg',
    DEPARTMENT_LINK: 'https://www.dept.aueb.gr/schools_department_photos/cs.png'
  },
  render: h => h(App),
}).$mount('#app');
