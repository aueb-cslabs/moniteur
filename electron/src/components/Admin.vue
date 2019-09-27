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
    import AdminSideBar from "./AdminSideBar";

    export default {
        name: 'Administration',

        created() {
            let r = document.documentElement;
            r.style.setProperty("--background", this.$root.$data['background_color']);
            r.style.setProperty("--navbar-bg-color", this.$root.$data['navbar_background_color']);
            r.style.setProperty("--navbar-color", this.$root.$data['navbar_color']);
        },

        components: {
            Login,
            AdminBar,
            AdminAnnouncement,
            AdminComment,
            AdminRoomAnnouncement,
            AdminSideBar
        },

        data() {
            return {
                authToken: authToken,
                auth: false,
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

            showSidebar: function () {
                if (!this.sideOpen){
                    let mw = window.matchMedia( "(max-width: 500px) ");
                    if (mw.matches){
                        document.getElementById("sidebar").style.width = "100%";
                        document.getElementById("main").style.marginLeft = "100%";
                    }
                    else {
                        document.getElementById("sidebar").style.width = "250px";
                        document.getElementById("main").style.marginLeft = "250px";
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