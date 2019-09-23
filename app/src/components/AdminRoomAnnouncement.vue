<template>
    <div>
        <div class="pt-5">
            <div>
                <form @submit="checkSearch">
                    <div class="row">
                        <div class="col-xs-10 col-md-10 form-group">
                            <input type="text" class="form-control" v-model="room" v-bind:placeholder="this.$t('message.adminRoomAnnRN')">
                        </div>
                        <div class="col-xs-2 col-md-2 form-group pl-4">
                            <button type="submit" class="btn btn-primary">{{$t("message.adminSearch")}}</button>
                        </div>
                    </div>
                </form>
                <p class="error">{{searchError}}</p>
            </div>
            <div v-if="roomBool">
                <form @submit="checkForm">
                    <div class="form-group">
                        <label for="message">{{$t("message.adminAnnForm1")}}</label>
                        <input type="text" class="form-control" id="message" v-model="form.msg" v-bind:placeholder="this.$t('message.adminAnnMsg')">
                    </div>
                    <div class="form-group">
                        <label for="expiration">{{$t("message.adminAnnExpire")}}</label>
                        <input type="text" class="form-control" id="expiration" v-model="form.end" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                    </div>
                    <button type="submit" class="btn btn-primary">{{$t("message.adminAnnSend")}}</button>
                    <p class="error">{{error}}</p>
                </form>
            </div>
            <div class="mt-5" v-if="announcement != null">
                <div class="alert alert-primary" role="alert">
                    <h3>{{$t("message.adminRoomAnnCurrent")}} {{announcement['msg']}}</h3>
                </div>
                <hr>
                <div class="alert alert-warning" role="alert">
                    {{$t("message.adminExpires")}} {{announcement['end']}}
                </div>
                <button type="submit" class="btn btn-danger" v-on:click="removeRoomAnnouncement">{{$t("message.removeAnn")}}</button>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: 'AdminRoomAnnouncement',

        data() {
            return {
                announcement: null,
                form: {
                    end: '',
                    msg: ''
                },
                error: null,
                searchError: null,
                room: '',
                roomBool: false
            }
        },

        methods: {
            send: function () {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + "/api/announcement/" + this.room,
                    headers: this.$parent.$data['config'],
                    data: {
                        end: this.form.end,
                        msg: this.form.msg
                    }
                }).then(response => {
                    if (response.status === 200) {
                        this.fetchRoomAnnouncement();
                    }
                });
            },

            removeRoomAnnouncement: function () {
                axios({
                    method: 'delete',
                    url: this.$root.$data['api'] + "/api/announcement/" + this.room,
                    headers: this.$parent.$data['config'],
                }).then(response => {
                    if (response.status === 200) {
                        this.announcement = null;
                    }
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
                    e.preventDefault();
                    return;
                }
                if (this.form.msg === '') {
                    this.error = this.$t('message.adminNoMsg');
                    e.preventDefault();
                    return;
                }
                if (this.form.end === '') {
                    this.error = this.$t('message.adminNoDate');
                    e.preventDefault();
                    return;
                }
                if (!this.isGoodDate(this.form.end)) {
                    this.error = this.$t('message.adminInvalidDate');
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
                    this.searchError = this.$t('message.adminRoomAnnSearchErr');
                    this.roomBool = false;
                    e.preventDefault();
                    return;
                }
                this.fetchRoomAnnouncement();
                this.searchError = '';
                this.roomBool = true;
                e.preventDefault();
            },

            isGoodDate: function(dt){
                let reGoodDate = /([0-3]?\d\/{1})([01]?\d\/{1})([12]{1}\d{3}\/?)/g;
                return reGoodDate.test(dt);
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/AdminRoomAnnouncement";
@import "~bootstrap/scss/bootstrap.scss";
</style>