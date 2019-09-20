<template>
    <div v-if="!auth" class="login">
        <p>Login to Moniteur</p>
        <div class="row">
            <form @submit="login">
                <div class="form-group">
                    <label for="username">Username</label>
                    <input type="username" class="form-control" id="username" v-model="loginForm.username" placeholder="Enter username">
                </div>
                <div class="form-group">
                    <label for="password">Password</label>
                    <input type="password" class="form-control" id="password" v-model="loginForm.password" placeholder="Password">
                </div>
                <button type="submit" class="btn btn-primary">Sign in</button>
            </form>
        </div>
    </div>
    <div v-else>
        <nav class="navbar navbar-expand navbar-dark">
            <a class='navbar-brand'>
                <img v-bind:src="this.$root.$data['logo_url']"
                     alt="logo" height="60" />
            </a>
            <nav class="navbar-nav mr-auto">
                <img class="ml-2" v-bind:src="this.$root.$data['secondary_logo_url']"
                     alt="cs_logo" height="60" />
                <a class="ml-4 navbar-brand mt-2">{{room}}</a>
            </nav>
            <nav class=""
        </nav>
    </div>
</template>

<script>
    import axios from "axios";

    export default {
        name: 'Administration',

        created() {
            let r = document.documentElement;
            r.style.setProperty("--background", this.$root.$data['background_color']);
            r.style.setProperty("--navbar-bg-color", this.$root.$data['navbar_background_color']);
            r.style.setProperty("--navbar-color", this.$root.$data['navbar_color']);
        },

        data() {
            return {
                token: "",
                username: "",
                expiration: 0,
                loginForm: {
                    username: "",
                    password: ""
                },
                auth: false
            }
        },

        methods: {
            login: function (e) {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + '/api/authenticate',
                    data: {
                        username: this.loginForm.username,
                        password: this.loginForm.password
                    }
                }).then(response => {
                    this.auth = true;
                    this.username = this.loginForm.username;
                    this.token = response.data['token_type'] + " " + response.data['access_token'];
                    this.expiration = response.data['expires_in'];
                });
                e.preventDefault();
            }
        }
    }

</script>

<style lang="scss">
@import "../scss/Admin.scss";
@import "~bootstrap/scss/bootstrap.scss";
</style>