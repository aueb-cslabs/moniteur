import Vue from 'vue'
import App from './App.vue'
import VueI18n from 'vue-i18n'
import VueCookies from 'vue-cookies'
import VueRouter from 'vue-router'
import BootstrapVue from 'bootstrap-vue'

import authToken from './auth';
import messages from './translations';
import AdminAnnouncement from "./components/AdminAnnouncement";
import AdminRoomAnnouncement from "./components/AdminRoomAnnouncement";
import AdminComment from "./components/AdminComment";
import UserManagement from "./components/UserManagement";
import AcademicCalendar from "./components/AcademicCalendar";
import UnscheduledLessons from "./components/UnscheduledLessons";
import Settings from "./components/Settings";

Vue.config.productionTip = false;
Vue.use(VueI18n);
Vue.use(VueCookies);
Vue.use(BootstrapVue);
Vue.use(VueRouter);

window.sideOpen = false;

const i18n = new VueI18n({
    locale: 'el',
    messages,
});

const routes = [
    { path: '/announcement', component: AdminAnnouncement },
    { path: '/roomAnnouncement', component: AdminRoomAnnouncement },
    { path: '/comment', component: AdminComment },
    { path: '/users', component: UserManagement },
    { path: '/calendar', component: AcademicCalendar },
    { path: '/unscheduled', component: UnscheduledLessons },
    { path: '/settings', component: Settings },
];

const router = new VueRouter({
    routes
});

new Vue({
    i18n,
    router,
    data() {
        return {
            authToken: authToken,
            config: {
                Authorization: authToken.token,
                Username: authToken.username
            },
        }
    },
    render: h => h(App),
}).$mount('#app');
