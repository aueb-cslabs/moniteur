<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminRAHeader")}}</h2>
            <div>
                <label>{{$t("message.adminRASelect")}}</label>
                <b-form-select @change="fetchRoomAnnouncement" v-model="room" :options="options"></b-form-select>
            </div>
            <div v-if="this.room != null" class="pt-5">
                <form @submit="checkForm">
                    <div class="form-group">
                        <label for="message">{{$t("message.adminAnnForm1")}}</label>
                        <input type="text" class="form-control" id="message" v-model="form.msg" v-bind:placeholder="this.$t('message.adminAnnMsg')">
                    </div>
                    <div class="form-group">
                        <label>{{$t("message.adminAnnExpire")}}</label>
                        <datetime v-bind:placeholder="this.$t('message.adminAnnDateForm')" format="DD/MM/YYYY H:i" v-model="form.end"></datetime>
                    </div>
                    <button type="submit" class="btn btn-primary float-right mb-4">{{$t("message.adminAnnSend")}}</button>
                </form>
            </div>
            <div class="pt-5" v-if="this.announcement != null">
                <div class="alert alert-primary" role="alert">
                    <h3>{{this.$t("message.adminRoomAnnCurrent")}} {{this.announcement['msg']}}</h3>
                </div>
                <hr>
                <div class="alert alert-warning" role="alert">
                    {{this.$t("message.adminExpires")}} {{this.announcement['end']}}
                </div>
                <button type="submit" class="btn btn-danger float-right" v-on:click="removeRoomAnnouncement">{{this.$t("message.removeAnn")}}</button>
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
    import functions from '../functions';
    import datetime from 'vuejs-datetimepicker';

    export default {
        name: 'AdminRoomAnnouncement',

        components: { datetime },

        created() {
            axios({
                method: 'get',
                url: this.$root.$data['api'] + '/api/rooms'
            }).then(response => {
                this.options = response.data;
            });
        },

        data() {
            return {
                announcement: null,
                form: {
                    end: '',
                    msg: ''
                },
                error: null,
                searchError: null,
                room: null,
                roomBool: false,
                dismissSecs: 5,
                dismissCountDown: 0,
                options: []
            }
        },

        methods: {
            send: function () {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + "/api/announcement/" + this.room,
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: {
                        end: new Date(this.form.end)/1000,
                        msg: this.form.msg
                    }
                }).then(() => {
                    this.fetchRoomAnnouncement();
                });
            },

            removeRoomAnnouncement: function () {
                axios({
                    method: 'delete',
                    url: this.$root.$data['api'] + "/api/announcement/" + this.room,
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                }).then(() => {
                    this.announcement = null;
                });
            },

            fetchRoomAnnouncement: function () {
                axios.get(this.$root.$data['api'] + "/api/announcement/" + this.room)
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
                this.form.end = '';
                this.form.msg = '';
                e.preventDefault();
            },

            checkSearch: function (e) {
                if (this.room === '') {
                    this.error = this.$t('message.adminRoomAnnSearchErr');
                    this.roomBool = false;
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                this.fetchRoomAnnouncement();
                this.error = '';
                this.roomBool = true;
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
@import "../scss/AdminRoomAnnouncement";
@import "~bootstrap/scss/bootstrap.scss";
</style>