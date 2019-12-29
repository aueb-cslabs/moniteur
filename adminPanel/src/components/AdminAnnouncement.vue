<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminAnnHeader")}}</h2>
            <div>
                <form @submit="checkForm">
                    <div class="form-group">
                        <label for="message">{{$t("message.adminAnnForm1")}}</label>
                        <input type="text" class="form-control" id="message" v-model="form.msg" v-bind:placeholder="this.$t('message.adminAnnMsg')">
                    </div>
                    <div class="form-group">
                        <label for="expiration">{{$t("message.adminAnnExpire")}}</label>
                        <input type="datetime-local" class="form-control" id="expiration" v-model="form.end" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                    </div>
                    <button type="submit" class="btn btn-primary">{{$t("message.adminAnnSend")}}</button>
                </form>
            </div>
            <div class="mt-5" v-if="this.announcement != null">
                <div class="alert alert-primary" role="alert">
                    <h3>{{this.$t("message.adminAnnCurrent")}} {{this.announcement['msg']}}</h3>
                </div>
                <hr>
                <div class="alert alert-warning" role="alert">
                    {{this.$t("message.adminExpires")}} {{this.announcement['end']}}
                </div>
                <button type="submit" class="btn btn-danger float-right" v-on:click="removeAnnouncement">{{this.$t("message.removeAnn")}}</button>
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
        name: 'AdminAnnouncement',

        components: {
            ErrorPopup
        },

        created() {
            this.fetchAnnouncement();
        },

        data() {
            return {
                announcement: null,
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
                    url: config.api + "/api/announcement",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: {
                        end: new Date(this.form.end)/1000,
                        msg: this.form.msg
                    }
                }).then(() => {
                    this.fetchAnnouncement();
                });
            },

            removeAnnouncement: function () {
                axios({
                    method: 'delete',
                    url: config.api + "/api/announcement",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                }).then(() => {
                    this.announcement = null;
                });
            },

            fetchAnnouncement: function () {
                axios.get(config.api + "/api/announcement")
                    .then(response => {
                        this.announcement = response.data;
                        let date = new Date(this.announcement['end']*1000).toLocaleDateString("el-GR");
                        this.announcement['end'] = date;
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
                    this.$refs.error.setError(this.$t('message.adminNoMsg'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.end === '') {
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
                this.form.msg = '';
                this.form.end = '';
                e.preventDefault();
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/AdminAnnouncement";
@import "~bootstrap/scss/bootstrap.scss";
</style>