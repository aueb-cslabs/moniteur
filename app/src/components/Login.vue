<template>
    <div style="width: 25%;">
        <p>Login to Moniteur</p>
        <form @submit="checkLogin">
            <div class="form-group">
                <label for="username">Username</label>
                <input type="text" class="form-control" id="username" v-model="loginForm.username" placeholder="Enter username">
            </div>
            <div class="form-group">
                <label for="password">Password</label>
                <input type="password" class="form-control" id="password" v-model="loginForm.password" placeholder="Password">
            </div>
            <button type="submit" class="btn btn-primary">Sign in</button>
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
                        this.error = "Invalid username or password.";
                    }
                    this.$parent.checkLogin();
                });
            },

            checkLogin: function (e) {
                if (this.loginForm.username === '' || this.loginForm.password === '') {
                    this.error = "You have not entered username or password!";
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