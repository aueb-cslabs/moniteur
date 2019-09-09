import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false;

const shared = {
  getRoom() {
    let link =  window.location.href.toString();
    const room = link.substring(link.lastIndexOf("/")+1);
    return room;
  }
};

new Vue({
  data : {
    room: shared.getRoom(),
    api: "http://localhost",
    port: ":27522",
    LOGO_LINK: 'https://www.aueb.gr/press/logos/1_AUEB-pantone-HR.jpg',
    DEPARTMENT_LINK: 'https://www.dept.aueb.gr/schools_department_photos/cs.png'
  },
  render: h => h(App),
}).$mount('#app');
