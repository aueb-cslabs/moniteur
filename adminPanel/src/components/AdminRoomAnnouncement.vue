<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminRAHeader")}}</h2>
            <div>
                <label>{{$t("message.adminRASelect")}}</label>
                <b-form-select @change="fetchRoomAnnouncement" v-model="room" :options="options"></b-form-select>
            </div>
            <div v-if="this.room != null" class="pt-5">
                <CreateAnnouncement
                        v-model="form"
                        :form="form"
                        :msg="form.msg"
                        :end="form.end"/>
                <button type="submit" class="btn btn-primary">{{$t("message.adminAnnSend")}}</button>
            </div>
            <div class="pt-5" v-if="this.announcement != null">
                <ShowAnnouncement
                        :type="this.$t('message.adminRoomAnnCurrent')"
                        :announcement="this.announcement['msg']"
                        :expiration="this.announcement['end']"/>
                <button type="submit"
                        class="btn btn-danger float-right"
                        v-on:click="removeRoomAnnouncement">{{this.$t("message.removeAnn")}}</button>
            </div>
        </div>
        <div>
            <ErrorPopup ref="error"/>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import CreateAnnouncement from "../CreateAnnouncement/CreateAnnouncement";
    import ErrorPopup from "../ErrorPopup/ErrorPopup";
    import ShowAnnouncement from "../ShowAnnouncement/ShowAnnouncement";

    const config = require('electron').remote.getGlobal('config');

    export default {
        name: 'AdminRoomAnnouncement',
        components: {ErrorPopup, CreateAnnouncement, ShowAnnouncement},
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
                options: [],
                form: {
                    msg: '',
                    end: ''
                }
            }
        },

        methods: {
            send: function () {
                axios({
                    method: 'post',
                    url: config.api + "/api/announcement/" + this.room,
                    headers: {
                        Username: this.$root.$data['authToken'].username,
                        Authorization: this.$root.$data['authToken'].token
                    },
                    data: {
                        end: new Date(this.form.end)/1000,
                        msg: this.form.msg
                    }
                }).then(() => {
                    this.fetchRoomAnnouncement();
                }).catch(()=> {
                    this.$refs.error.setError(this.$t('message.calFail'));
                    this.$refs.error.showAlert();
                });
            },

            removeRoomAnnouncement: function () {
                axios({
                    method: 'delete',
                    url: config.api + "/api/announcement/" + this.room,
                    headers: {
                        Username: this.$root.$data['authToken'].username,
                        Authorization: this.$root.$data['authToken'].token
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