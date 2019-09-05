import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false;

const shared = {
  getRoom() {
    let link =  window.location.href.toString();
    let room = link.substring(link.lastIndexOf("/")+1);
    return room
  }
};

new Vue({
  data : {
    shared: shared.getRoom()
  },
  render: h => h(App),
}).$mount('#app');
