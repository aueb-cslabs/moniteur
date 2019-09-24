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
                    <li class="nav-item pr-3" v-on:click="this.$parent.showAnnouncement"> {{$t("message.adminBarGA")}}</li>
                    <li class="nav-item pr-3" v-on:click="this.$parent.showRoomAnnouncement">{{$t("message.adminBarRA")}}</li>
                    <li class="nav-item pr-3" v-on:click="this.$parent.showComment">{{$t("message.adminBarGC")}}</li>
                </ul>
                <ul class="navbar-nav ml-auto">
                    <li class="nav-item pr-3" v-if="this.$parent.$data['authToken'].auth">{{$t("message.adminBarLI")}} {{this.$parent.$data['authToken'].username}}</li>
                    <li class="nav-item" v-on:click="logout">{{$t("message.adminBarLogout")}}</li>
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
                this.$parent.$data['authToken'].auth = false;
                this.$cookies.remove('test');
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