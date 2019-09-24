<template>
    <div style="width: 25%;">
        <p>{{ $t("message.loginWelcome") }}</p>
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
                loginForm: {
                    username: "",
                    password: ""
                },
                error: "",
                cookie: {
                    username: '',
                    token: ''
                }
            }
        },

        created() {
            if (this.$cookies.get('session') != null) {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + '/api/validate',
                    headers: {
                        username: this.$cookies.get('session').cookie.username,
                        authorization: this.$cookies.get('session').cookie.token
                    }
                }).then(res => {
                    if (res.status === 200) {
                        this.$parent.$data['authToken'].auth = true;
                        this.$parent.$data['authToken'].username = this.$cookies.get('session').cookie.username;
                        this.$parent.$data['authToken'].token = this.$cookies.get('session').cookie.token;
                        this.$parent.$data['authToken'].expiration = this.$cookies.get('session').expiration;
                    }
                    else {
                        this.$parent.$data['authToken'].auth = false;
                        this.$parent.$data['authToken'].username = '';
                        this.$parent.$data['authToken'].token = '';
                        this.$parent.$data['authToken'].expiration = '';
                    }
                })
            }
            if (this.$cookies.get('test') != null) {
                this.$parent.$data['authToken'].auth = true;
            }
        },

        methods: {
            login: function () {
                this.$cookies.set('test', 'abc');
                this.$parent.$data['authToken'].auth=true;
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
                        this.cookie.token = this.$parent.$data['authToken'].token;
                        this.cookie.username = this.$parent.$data['authToken'].username;
                        this.$cookies.set('session', this.cookie, new Date(this.$parent.$data['authToken'].expiration));
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