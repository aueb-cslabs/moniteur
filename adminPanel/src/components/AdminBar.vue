<template>
    <nav class="navbar navbar-expand navbar-dark">
        <button type="button" class="btn" v-on:click="this.$parent.showSidebar">
            <i class="fas fa-list-ul fa-2x" style="color: white;"></i>
        </button>
        <a class='navbar-brand d-none d-sm-block'>
            <img v-bind:src="this.logo_url"
            alt="logo" height="60" />
        </a>
        <a class="navbar-nav mr-auto d-none d-sm-block">
            <img class="ml-2" v-bind:src="this.secondary_logo_url"
                 alt="cs_logo" height="60" />
        </a>
        <ul class="navbar-nav ml-auto">
            <li class="nav-item pr-3 d-none d-sm-block" v-if="this.$parent.$data['authToken'].auth">{{$t("message.adminBarLI")}} {{this.$parent.$data['authToken'].username}}</li>
            <li class="nav-item" v-on:click="logout">{{$t("message.adminBarLogout")}}</li>
        </ul>
    </nav>
</template>

<script>
    import axios from "axios";

    const config = require('electron').remote.getGlobal('config');

    export default {
        name: 'AdminBar',

        created() {
            this.username = this.$cookies.get('session').Username;
        },

        data() {
            return {
                username: '',
                sideOpen: false,
                logo_url: config.logo_url,
                secondary_logo_url: config.secondary_logo_url
            }
        },

        methods: {
            logout: function (e) {
                //let cookie = this.$cookies.get('session');
                axios({
                    method: 'post',
                    url: config.api + "/api/invalidate",
                    headers: {
                        Authorization: this.$parent.$data['authToken'].token,
                        Username: this.$parent.$data['authToken'].username
                    }
                }).then(() => {
                    this.$parent.$data['authToken'].auth = false;
                    this.$cookies.remove('session');
                });
                e.preventDefault();
            },
        }
    }
</script>

<style lang="scss">
@import "../scss/AdminBar.scss";
@import "~bootstrap/scss/bootstrap.scss";
@import "~@fortawesome/fontawesome-free/css/all.css";
</style>