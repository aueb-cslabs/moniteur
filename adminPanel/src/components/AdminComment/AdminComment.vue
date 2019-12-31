<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.adminComHeader")}}</h2>
            <CreateAnnouncement
                    v-model="form"
                    :form="form"
                    :msg="form.msg"
                    :end="form.end"/>
            <button type="submit" class="btn btn-primary">{{$t("message.adminAnnSend")}}</button>
            <div class="mt-5" v-if="this.comment != null">
                <ShowAnnouncement
                        :type="this.$t('message.adminCommCurrent')"
                        :announcement="this.comment['msg']"
                        :expiration="this.comment['end']"/>
                <button type="submit"
                        class="btn btn-primary float-right"
                        v-on:click="removeComment">{{this.$t("message.removeComm")}}</button>
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
        name: 'AdminComment',
        components: {ErrorPopup, CreateAnnouncement, ShowAnnouncement},

        created() {
            this.fetchComment();
        },

        data() {
            return {
                comment: null,
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
                    url: config.api + "/api/comment",
                    headers: {
                        Username: this.$root.$data['authToken'].username,
                        Authorization: this.$root.$data['authToken'].token
                    },
                    data: {
                        end: new Date(this.form.end)/1000,
                        msg: this.form.msg
                    }
                }).then(() => {
                    this.fetchComment();
                }).catch(()=> {
                    this.$refs.error.setError(this.$t('message.calFail'));
                    this.$refs.error.showAlert();
                });
            },

            removeComment: function () {
                axios({
                    method: 'delete',
                    url: config.api + "/api/comment",
                    headers: {
                        Username: this.$root.$data['authToken'].username,
                        Authorization: this.$root.$data['authToken'].token
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
    @import "./AdminComment";
    @import "~bootstrap/scss/bootstrap.scss";
</style>