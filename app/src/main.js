import Vue from 'vue'
import App from './App.vue'

Vue.config.productionTip = false;

const shared = {
  getRoom() {
    let link =  window.location.href.toString();
    const room = link.substring(link.lastIndexOf("/")+1);
    return room;
  },

  getAPI() {
    return "http://localhost";
  },

  getPort() {
    return ":27522"
  }
};

new Vue({
  data : {
    room: shared.getRoom(),
    api: shared.getAPI(),
    port: shared.getPort()
  },
  render: h => h(App),
}).$mount('#app');
