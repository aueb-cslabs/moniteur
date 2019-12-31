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
        <ErrorPopup ref="error"/>
    </div>
</template>

<script>
    import axios from "axios";
    import ErrorPopup from "../ErrorPopup/ErrorPopup";

    const config = require('electron').remote.getGlobal('config');

    export default {
        components: {ErrorPopup},
        created() {
            this.updateUsers();
        },

        data() {
            return {
                user: '',
                options: [],
                userOption: null
            }
        },

        methods: {
            checkUser: function (e) {
                if (this.user === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                }
                this.addUser();
                this.user = null;
                e.preventDefault()
            },

            addUser: function() {
                axios({
                    method: 'post',
                    url: config.api + '/api/register/' + this.user,
                    headers: {
                        Username: this.$root.$data['authToken'].username,
                        Authorization: this.$root.$data['authToken'].token
                    }
                }).then(() => {
                    this.updateUsers();
                })
            },

            updateUsers: function() {
                axios({
                    method: 'get',
                    url: config.api + '/api/users',
                    headers: {
                        Username: this.$root.$data['authToken'].username,
                        Authorization: this.$root.$data['authToken'].token
                    }
                }).then(response => {
                    this.options = response.data;
                })
            },

            removeUser: function() {
                axios({
                    method: 'post',
                    url: config.api + '/api/unregister/' + this.userOption,
                    headers: {
                        Username: this.$root.$data['authToken'].username,
                        Authorization: this.$root.$data['authToken'].token
                    }
                }).then(() => {
                    this.updateUsers();
                })
            }
        }
    }

</script>

<style lang="scss">
@import "../scss/UserManagement";
@import "~bootstrap/scss/bootstrap";
</style>