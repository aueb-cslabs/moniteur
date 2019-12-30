<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminAnnHeader")}}</h2>
            <CreateAnnouncement/>
            <div class="mt-5" v-if="this.announcement != null">
                <ShowAnnouncement
                        :type="this.$t('message.adminAnnCurrent')"
                        :announcement="this.announcement['msg']"
                        :expiration="this.announcement['end']"/>
                <button type="submit" class="btn btn-danger float-right" v-on:click="removeAnnouncement">{{this.$t("message.removeAnn")}}</button>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';
    import CreateAnnouncement from "../CreateAnnouncement/CreateAnnouncement";
    import ShowAnnouncement from "../ShowAnnouncement/ShowAnnouncement";

    const config = require('electron').remote.getGlobal('config');

    export default {
        name: 'AdminAnnouncement',

        components: {
            ShowAnnouncement,
            CreateAnnouncement
        },

        created() {
            this.fetchAnnouncement();
        },

        data() {
            return {
                announcement: null,
            }
        },

        methods: {
            send: function (form) {
                axios({
                    method: 'post',
                    url: config.api + "/api/announcement",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: {
                        end: new Date(form.end)/1000,
                        msg: form.msg
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
                        if (response.data !== null){
                            this.announcement = response.data;
                            let date = new Date(this.announcement['end']*1000).toLocaleDateString("el-GR");
                            this.announcement['end'] = date;
                        }
                    });
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/AdminAnnouncement";
@import "~bootstrap/scss/bootstrap.scss";
</style>