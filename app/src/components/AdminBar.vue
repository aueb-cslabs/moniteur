<template>
    <div class="top">
        <nav class="navbar navbar-expand navbar-dark">
            <a class='navbar-brand'>
                <img v-bind:src="this.$root.$data['logo_url']"
                     alt="logo" height="60" />
            </a>
            <nav class="navbar-nav mr-auto">
                <img class="ml-2" v-bind:src="this.$root.$data['secondary_logo_url']"
                     alt="cs_logo" height="60" />
            </nav>
            <div class="collapse navbar-collapse ml-3">
                <ul class="navbar-nav mr-auto">
                    <li class="nav-item">
                        <a class="nav-link" v-on:click="this.$parent.showAnnouncement">General Announcement</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" v-on:click="this.$parent.showRoomAnnouncement">Room Announcement</a>
                    </li>
                    <li class="nav-item">
                        <a class="nav-link" v-on:click="this.$parent.showComment">General Comment</a>
                    </li>
                </ul>
                <ul class="navbar-nav ml-auto">
                    <li class="nav-link">
                        <a class="nav-text" v-on:click="logout">Logout</a>
                    </li>
                </ul>
            </div>
        </nav>
    </div>
</template>

<script>
    import axios from "axios";

    export default {
        name: 'AdminBar',

        methods: {
            logout: function (e) {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + "/api/invalidate",
                    headers: {
                        Authorization: this.$parent.$data['authToken'].token,
                        Username: this.$parent.$data['authToken'].username
                    }
                }).then(response => {
                    if (response.status === 200) {
                        this.$parent.$data['authToken'].username = "";
                        this.$parent.$data['authToken'].token = "";
                        this.$parent.$data['authToken'].expiration = 0;
                        this.$parent.$data['authToken'].auth = false;
                    }
                });
                e.preventDefault();
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/AdminBar.scss";
@import "~bootstrap/scss/bootstrap.scss";
</style>