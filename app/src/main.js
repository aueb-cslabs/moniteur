import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false;

const shared = {
  getRoom() {
    let link =  window.location.href.toString();
    const room = link.substring(link.lastIndexOf("/")+1);
    return room
  }
};

new Vue({
  data : {
    room: shared.getRoom()
  },
  render: h => h(App),
}).$mount('#app');
