<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminComHeader")}}</h2>
            <div>
                <form @submit="checkForm">
                    <div class="form-group">
                        <label for="message">{{$t("message.adminCommForm1")}}</label>
                        <input type="text" class="form-control" id="message" v-model="form.msg" v-bind:placeholder="this.$t('message.adminAnnMsg')">
                    </div>
                    <div class="form-group">
                        <label for="expiration">{{$t("message.adminAnnExpire")}}</label>
                        <input type="datetime-local" class="form-control" id="expiration" v-model="form.end" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                    </div>
                    <button type="submit" class="btn btn-primary">{{$t("message.adminCommSend")}}</button>
                </form>
            </div>
            <div class="mt-5" v-if="this.comment != null">
                <div class="alert alert-primary" role="alert">
                    <h3>{{this.$t("message.adminCommCurrent")}} {{this.comment['msg']}}</h3>
                </div>
                <hr>
                <div class="alert alert-warning" role="alert">
                    {{this.$t("message.adminExpires")}} {{this.comment['end']}}
                </div>
                <button type="submit" class="btn btn-primary float-right" v-on:click="removeComment">{{this.$t("message.removeComm")}}</button>
            </div>
        </div>
        <ErrorPopup ref="error"/>
    </div>
</template>

<script>
    import axios from 'axios';
    import functions from "../functions";
    import ErrorPopup from "../ErrorPopup/ErrorPopup";

    const config = require('electron').remote.getGlobal('config');

    export default {
        name: 'AdminComment',
        components: {ErrorPopup},

        created() {
            this.fetchComment();
        },

        data() {
            return {
                comment: null,
                form: {
                    end: '',
                    msg: ''
                },
            }
        },

        methods: {
            send: function () {
                axios({
                    method: 'post',
                    url: config.api + "/api/comment",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: {
                        end: new Date(this.form.end)/1000,
                        msg: this.form.msg
                    }
                }).then(() => {
                    this.fetchComment();
                });
            },

            removeComment: function () {
                axios({
                    method: 'delete',
                    url: config.api + "/api/comment",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                }).then(() => {
                    this.comment = null;
                });
            },

            fetchComment: function () {
                axios.get(config.api + "/api/comment")
                    .then(response => {
                        this.comment = response.data;
                        let date = new Date(this.comment['end']*1000).toLocaleDateString("el-GR");
                        this.comment['end'] = date;
                    });
            },

            checkForm: function (e) {
                if (this.form.msg === '' && this.form.end === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.msg === '') {
                    this.$refs.error.setError(this.$t('message.adminNoDate'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (!functions.isGoodDate(this.form.end)) {
                    this.$refs.error.setError(this.$t('message.adminInvalidDate'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                this.send();
                this.form.end = '';
                this.form.msg = '';
                e.preventDefault();
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/AdminComment";
    @import "~bootstrap/scss/bootstrap.scss";
</style>