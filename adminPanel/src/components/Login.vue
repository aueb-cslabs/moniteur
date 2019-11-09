<template>
    <div class="form">
        <p>{{ $t("message.loginWelcome") }}</p>
        <form @submit="checkLogin">
            <div class="form-group">
                <input type="text" class="form-control" id="username" v-model="loginForm.username" v-bind:placeholder="this.$t('message.loginUsername')">
            </div>
            <div class="form-group">
                <input type="password" class="form-control" id="password" v-model="loginForm.password" v-bind:placeholder="this.$t('message.loginPassword')">
            </div>
            <button @click="checkLogin" class="btn btn-primary">{{$t("message.loginSignIn")}}</button>
        </form>
        <div class="error">
            <b-alert
                    :show="dismissCountDown"
                    dismissible
                    variant="danger"
                    @dismissed="dismissCountDown=0"
                    @dismiss-count-down="countDownChanged"
            >
                <p>Error! {{error}}</p>
                <b-progress
                        variant="dark"
                        :max="dismissSecs"
                        :value="dismissCountDown"
                        height="4px"
                ></b-progress>
            </b-alert>
        </div>
    </div>
</template>

<script>
    import axios from "axios";

    export default {
        data() {
            return {
                loginForm: {
                    username: "",
                    password: ""
                },
                error: "",
                cookie: {
                    Username: '',
                    Authorization: ''
                },
                dismissSecs: 5,
                dismissCountDown: 0,
            }
        },

        created() {
            let cookie = this.$cookies.get('session');
            if (cookie != null) {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + '/api/validate',
                    headers: cookie
                }).then(() => {
                    this.$parent.$data['authToken'].auth = true;
                    this.$parent.$data['authToken'].username = cookie.Username;
                    this.$parent.$data['authToken'].token = cookie.Authorization;
                })
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
                    this.cookie.Authorization = response.data['token_type'] + " " + response.data['access_token'];
                    this.cookie.Username = this.loginForm.username;
                    this.$parent.$data['authToken'].auth = true;
                    this.$parent.$data['authToken'].username = this.loginForm.username;
                    this.$parent.$data['authToken'].token = response.data['token_type'] + " " + response.data['access_token'];
                        let date = new Date(response.data['expires_in']*1000);
                    this.$cookies.set('session', this.cookie, date);
                }).catch(() => {
                    this.error = this.$t('message.loginError');
                    this.showAlert();
                })
            },

            checkLogin: function (e) {
                if (this.loginForm.username === '' || this.loginForm.password === '') {
                    this.error = this.$t('message.loginFormError');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                this.error = '';
                this.login();
                e.preventDefault();
            },

            countDownChanged(dismissCountDown) {
                this.dismissCountDown = dismissCountDown
            },

            showAlert() {
                this.dismissCountDown = this.dismissSecs
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/Login.scss";
@import "~bootstrap/scss/bootstrap.scss";
</style>