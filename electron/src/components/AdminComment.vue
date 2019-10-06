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
                        <input type="text" class="form-control" id="expiration" v-model="form.end" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
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
                <button type="submit" class="btn btn-primary" v-on:click="removeComment">{{this.$t("message.removeComm")}}</button>
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

    export default {
        name: 'AdminComment',

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
                error: null,
                dismissSecs: 5,
                dismissCountDown: 0,
            }
        },

        methods: {
            send: function () {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + "/api/comment",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: {
                        end: this.form.end,
                        msg: this.form.msg
                    }
                }).then(() => {
                    this.fetchComment();
                });
            },

            removeComment: function () {
                axios({
                    method: 'delete',
                    url: this.$root.$data['api'] + "/api/comment",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                }).then(() => {
                    this.comment = null;
                });
            },

            fetchComment: function () {
                axios.get(this.$root.$data['api'] + "/api/comment")
                    .then(response => {
                        this.comment = response.data;
                        let date = new Date(this.comment['end']*1000).toLocaleDateString("el-GR");
                        this.comment['end'] = date;
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
                    this.error = this.$t('message.adminNoDate');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (!this.isGoodDate(this.form.end)) {
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

            isGoodDate: function(dt){
                let reGoodDate = /([0-3]?\d\/{1})([01]?\d\/{1})([12]{1}\d{3}\/?)/g;
                return reGoodDate.test(dt);
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
    @import "../scss/AdminComment";
    @import "~bootstrap/scss/bootstrap.scss";
</style>