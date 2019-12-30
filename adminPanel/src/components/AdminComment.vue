<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminComHeader")}}</h2>
            <CreateAnnouncement></CreateAnnouncement>
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
    </div>
</template>

<script>
    import axios from 'axios';
    import CreateAnnouncement from "../CreateAnnouncement/CreateAnnouncement";

    const config = require('electron').remote.getGlobal('config');

    export default {
        name: 'AdminComment',
        components: {CreateAnnouncement},

        created() {
            this.fetchComment();
        },

        data() {
            return {
                comment: null,
            }
        },

        methods: {
            send: function (form) {
                axios({
                    method: 'post',
                    url: config.api + "/api/comment",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: {
                        end: new Date(form.end)/1000,
                        msg: form.msg
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
                        if (response.data !== null) {
                            this.comment = response.data;
                            let date = new Date(this.comment['end']*1000).toLocaleDateString("el-GR");
                            this.comment['end'] = date;
                        }
                    });
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/AdminComment";
    @import "~bootstrap/scss/bootstrap.scss";
</style>