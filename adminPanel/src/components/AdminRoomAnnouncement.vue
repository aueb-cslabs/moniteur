<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminRAHeader")}}</h2>
            <div>
                <label>{{$t("message.adminRASelect")}}</label>
                <b-form-select @change="fetchRoomAnnouncement" v-model="room" :options="options"></b-form-select>
            </div>
            <CreateAnnouncement v-if="this.room != null" class="pt-5"/>
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
    </div>
</template>

<script>
    import axios from 'axios';
    import CreateAnnouncement from "../CreateAnnouncement/CreateAnnouncement";

    const config = require('electron').remote.getGlobal('config');

    export default {
        name: 'AdminRoomAnnouncement',
        components: {CreateAnnouncement},
        created() {
            axios({
                method: 'get',
                url: config.api + '/api/rooms'
            }).then(response => {
                this.options = response.data;
            });
        },

        data() {
            return {
                announcement: null,
                searchError: null,
                room: null,
                roomBool: false,
                options: []
            }
        },

        methods: {
            send: function (form) {
                axios({
                    method: 'post',
                    url: config.api + "/api/announcement/" + this.room,
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: {
                        end: new Date(form.end)/1000,
                        msg: form.msg
                    }
                }).then(() => {
                    this.fetchRoomAnnouncement();
                });
            },

            removeRoomAnnouncement: function () {
                axios({
                    method: 'delete',
                    url: config.api + "/api/announcement/" + this.room,
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                }).then(() => {
                    this.announcement = null;
                });
            },

            fetchRoomAnnouncement: function () {
                axios.get(config.api + "/api/announcement/" + this.room)
                    .then(response => {
                        if (response.data !== null) {
                            this.announcement = response.data;
                            let date = new Date(this.announcement['end']*1000).toLocaleDateString("el-GR");
                            this.announcement['end'] = date;
                        }
                    });
            },

            checkSearch: function (e) {
                if (this.room === '') {
                    this.$refs.error.setError(this.$t('message.adminRoomAnnSearchErr'));
                    this.roomBool = false;
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                this.fetchRoomAnnouncement();
                this.roomBool = true;
                e.preventDefault();
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/AdminRoomAnnouncement";
@import "~bootstrap/scss/bootstrap.scss";
</style>