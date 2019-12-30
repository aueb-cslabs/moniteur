<template>
    <div v-if="!this.authToken.auth" class="login">
        <Login/>
    </div>
    <div v-else class="actions">
        <AdminBar/>
        <div class="admin">
            <AdminSideBar/>
            <div id="main">
                <div v-if="type === 'aa'">
                    <AdminAnnouncement/>
                </div>
                <div v-if="type === 'ara'">
                    <AdminRoomAnnouncement/>
                </div>
                <div v-if="type === 'ac'">
                    <AdminComment/>
                </div>
                <div v-if="type === 'um'">
                    <UserManagement/>
                </div>
                <div v-if="type === 'cal'">
                    <AcademicCalendar/>
                </div>
                <div v-if="type === 'unles'">
                    <UnscheduledLessons/>
                </div>
                <div v-if="type === 'settings'">
                    <Settings/>
                </div>
            </div>
        </div>
    </div>
</template>

<script>
    import authToken from "../auth";
    import Login from "./Login";
    import AdminBar from "./AdminBar";
    import AdminAnnouncement from "./AdminAnnouncement";
    import AdminComment from "./AdminComment";
    import AdminRoomAnnouncement from "./AdminRoomAnnouncement";
    import UserManagement from "./UserManagement";
    import AdminSideBar from "./AdminSideBar";
    import AcademicCalendar from "./AcademicCalendar";
    import Settings from "./Settings";
    import UnscheduledLessons from "./UnscheduledLessons";

    const config = require('electron').remote.getGlobal('config');

    export default {

        name: 'Administration',

        created() {
            let r = document.documentElement;
            r.style.setProperty("--background", config.background_color);
            r.style.setProperty("--navbar-bg-color", config.navbar_background_color);
            r.style.setProperty("--navbar-color", config.navbar_color);
        },

        components: {
            UnscheduledLessons,
            Settings,
            Login,
            AdminBar,
            AdminAnnouncement,
            AdminComment,
            AdminRoomAnnouncement,
            AdminSideBar,
            UserManagement,
            AcademicCalendar,
            //ExamsCalendar,
        },

        data() {
            return {
                authToken: authToken,
                type: '',
                config: {
                    Authorization: authToken.token,
                    Username: authToken.username
                }
            }
        },

        methods: {
            showAnnouncement: function () {
                this.type = 'aa';
                this.showSidebar();
            },

            showRoomAnnouncement: function () {
                this.type = 'ara';
                this.showSidebar();
            },

            showComment: function () {
                this.type = 'ac';
                this.showSidebar();
            },

            showUserManagement: function() {
                this.type = 'um';
                this.showSidebar();
            },

            showAcademicCalendar: function() {
                this.type = 'cal';
                this.showSidebar();
            },

            showExamsCalendar: function() {
                this.type = 'exams';
                this.showSidebar();
            },

            showUnscheduledLessons: function() {
                this.type = 'unles';
                this.showSidebar();
            },

            showSettings: function() {
                this.type = 'settings';
                this.showSidebar();
            },

            showSidebar: function () {
                if (!this.sideOpen){
                    let mw = window.matchMedia( "(max-width: 500px) ");
                    if (mw.matches){
                        document.getElementById("sidebar").style.width = "100%";
                        document.getElementById("main").style.marginLeft = "100%";
                    }
                    else {
                        document.getElementById("sidebar").style.width = "280px";
                        document.getElementById("main").style.marginLeft = "280px";
                    }
                    this.sideOpen = true;
                }
                else {
                    document.getElementById("sidebar").style.width = "0px";
                    document.getElementById("main").style.marginLeft = "0px";
                    this.sideOpen = false;
                }
            }
        }
    }

</script>

<style lang="scss">
@import "../scss/Admin.scss";
@import "~bootstrap/scss/bootstrap.scss";
</style>