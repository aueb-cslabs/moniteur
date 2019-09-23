<template>
    <div style="width: 25%;">
        <p>{{ $t("message.loginWelcome" )}}</p>
        <form @submit="checkLogin">
            <div class="form-group">
                <input type="text" class="form-control" id="username" v-model="loginForm.username" v-bind:placeholder="this.$t('message.loginUsername')">
            </div>
            <div class="form-group">
                <input type="password" class="form-control" id="password" v-model="loginForm.password" v-bind:placeholder="this.$t('message.loginPassword')">
            </div>
            <button type="submit" class="btn btn-primary">{{$t("message.loginSignIn")}}</button>
        </form>
        <p>{{error}}</p>
    </div>
</template>

<script>
    import axios from "axios";

    export default {
        data() {
            return {
                token: "",
                username: "",
                expiration: 0,
                loginForm: {
                    username: "",
                    password: ""
                },
                error: ""
            }
        },

        methods: {
            login: function () {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + '/api/authenticate',
                    data: {
                        username: this.loginForm.username,
                        password: this.loginForm.password
                    }
                }).then(response => {
                    if (response.status===200){
                        this.$parent.$data['authToken'].username = this.loginForm.username;
                        this.$parent.$data['authToken'].token = response.data['token_type'] + " " + response.data['access_token'];
                        this.$parent.$data['authToken'].expiration = response.data['expires_in'];
                        this.$parent.$data['authToken'].auth = true;
                        this.$parent.$data['config'].Authorization = response.data['token_type'] + " " + response.data['access_token'];
                        this.$parent.$data['config'].Username = this.loginForm.username;
                        this.error = "";
                    }
                    else {
                        this.error = this.$t('message.loginError');
                    }
                    this.$parent.checkLogin();
                });
            },

            checkLogin: function (e) {
                if (this.loginForm.username === '' || this.loginForm.password === '') {
                    this.error = this.$t('message.loginFormError');
                    e.preventDefault();
                    return;
                }
                this.error = '';
                this.login();
                e.preventDefault();
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/Login.scss";
@import "~bootstrap/scss/bootstrap.scss";
</style>