import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'

import messages from './translations'

Vue.config.productionTip = false;

Vue.use(VueI18n);

const i18n = new VueI18n({
  locale: 'el',
  messages,
});

new Vue({
  i18n,

  render: h => h(App),
}).$mount('#app');
