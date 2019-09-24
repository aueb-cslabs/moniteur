import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'
import VueCookies from 'vue-cookies'

import config from './config'
import messages from './translations'

Vue.config.productionTip = false;
Vue.use(VueI18n);
Vue.use(VueCookies);

const i18n = new VueI18n({
  locale: 'el',
  messages,
});

new Vue({
  i18n,
  data : {
    ...config,
  },
  render: h => h(App),
}).$mount('#app');
