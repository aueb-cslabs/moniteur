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
                        <!--<input type="text" class="form-control" id="expiration" v-model="form.end" v-bind:placeholder="this.$t('message.adminAnnDateForm')">-->
                        <label>{{$t("message.adminAnnExpire")}}</label>
                        <datetime v-bind:placeholder="this.$t('message.adminAnnDateForm')" format="DD/MM/YYYY H:i" v-model="form.end"></datetime>
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
    import axios from 'axios';
    import functions from "../functions";
    import datetime from 'vuejs-datetimepicker';

    export default {
        name: 'AdminAnnouncement',

        components: { datetime },

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
                error: null,
                dismissSecs: 5,
                dismissCountDown: 0,
            }
        },

        methods: {
            send: function () {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + "/api/announcement",
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
                    url: this.$root.$data['api'] + "/api/announcement",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                }).then(() => {
                    this.announcement = null;
                });
            },

            fetchAnnouncement: function () {
                axios.get(this.$root.$data['api'] + "/api/announcement")
                    .then(response => {
                        this.announcement = response.data;
                        let date = new Date(this.announcement['end']*1000).toLocaleDateString("el-GR");
                        this.announcement['end'] = date;
                    });
            },

            checkForm: function (e) {
                if (this.form.msg === '' && this.form.end === '') {
                    this.error = this.$t('message.adminEmptyForm');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.msg === '') {
                    this.error = this.$t('message.adminNoMsg');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.end === '') {
                    this.error = this.$t('message.adminNoDate');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (!functions.isGoodDate(this.form.end)) {
                    this.error = this.$t('message.adminInvalidDate');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                this.error = null;
                this.send();
                this.form.msg = '';
                this.form.end = '';
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
@import "../scss/AdminAnnouncement";
@import "~bootstrap/scss/bootstrap.scss";
</style>