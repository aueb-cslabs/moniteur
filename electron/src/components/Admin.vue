<template>
    <div v-if="!this.authToken.auth" class="login">
        <Login/>
    </div>
    <div v-else class="actions">
        <AdminBar/>
        <div class="admin">
            <div v-if="announcementBool">
                <AdminAnnouncement/>
            </div>
            <div v-if="roomAnnouncementBool">
                <AdminRoomAnnouncement/>
            </div>
            <div v-if="commentBool">
                <AdminComment/>
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
            AdminRoomAnnouncement
        },

        data() {
            return {
                authToken: authToken,
                auth: false,
                announcementBool: false,
                roomAnnouncementBool: false,
                commentBool: false,
                config: {
                    Authorization: authToken.token,
                    Username: authToken.username
                }
            }
        },

        methods: {
            showAnnouncement: function () {
                this.announcementBool = true;
                this.roomAnnouncementBool = false;
                this.commentBool = false;
            },

            showRoomAnnouncement: function () {
                this.announcementBool = false;
                this.roomAnnouncementBool = true;
                this.commentBool = false;
            },

            showComment: function () {
                this.announcementBool = false;
                this.roomAnnouncementBool = false;
                this.commentBool = true;
            },
        }
    }

</script>

<style lang="scss">
@import "../scss/Admin.scss";
@import "~bootstrap/scss/bootstrap.scss";
</style>