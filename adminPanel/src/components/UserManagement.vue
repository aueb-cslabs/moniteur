<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.umAddUser")}}</h2>
            <div>
                <form @submit="checkUser">
                    <div class="row">
                        <div class="col-xs-10 col-md-10 form-group">
                            <input type="text" class="form-control" v-model="user" v-bind:placeholder="this.$t('message.umUsername')">
                        </div>
                        <div class="col-xs-2 col-md-2 form-group">
                            <button type="submit" class="btn btn-primary float-right">{{$t("message.umAdd")}}</button>
                        </div>
                    </div>
                </form>
            </div>
            <div class="pt-5">
                <h2>{{$t("message.umRemoveUser")}}</h2>
                <div>
                    <label>{{$t("message.umSelect")}}</label>
                    <b-form-select v-model="userOption" :options="options"></b-form-select>
                </div>
                <b-button @click="removeUser" variant="danger" class="mt-3">{{$t("message.umRemove")}}</b-button>
            </div>
        </div>
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
        created() {
            this.updateUsers();
        },

        data() {
            return {
                user: '',
                dismissSecs: 5,
                dismissCountDown: 0,
                error: null,
                options: [],
                userOption: null
            }
        },

        methods: {
            checkUser: function (e) {
                if (this.user === '') {
                    this.error = this.$t('message.adminEmptyForm');
                    this.showAlert();
                }
                this.addUser();
                this.user = null;
                e.preventDefault()
            },

            addUser: function() {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + '/api/register/' + this.user,
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    }
                }).then(() => {
                    this.updateUsers();
                })
            },

            updateUsers: function() {
                axios({
                    method: 'get',
                    url: this.$root.$data['api'] + '/api/users',
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    }
                }).then(response => {
                    this.options = response.data;
                })
            },

            removeUser: function() {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + '/api/unregister/' + this.userOption,
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    }
                }).then(() => {
                    this.updateUsers();
                })
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
@import "../scss/UserManagement";
@import "~bootstrap/scss/bootstrap";
</style>