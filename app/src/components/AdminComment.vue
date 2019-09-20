<template>
    <div>
        <div class="pt-5">
            <div>
                <form @submit="checkForm">
                    <div class="form-group">
                        <label for="message">Comment</label>
                        <input type="text" class="form-control" id="message" v-model="form.msg" placeholder="Message">
                    </div>
                    <div class="form-group">
                        <label for="expiration">Expiration Date</label>
                        <input type="text" class="form-control" id="expiration" v-model="form.end" placeholder="DD/MM/YYYY">
                    </div>
                    <button type="submit" class="btn btn-primary">Send comment</button>
                    <p class="error">{{error}}</p>
                </form>
            </div>
            <div class="mt-5" v-if="comment != null">
                <div class="alert alert-primary" role="alert">
                    <h3>Current Comment: {{comment['msg']}}</h3>
                </div>
                <hr>
                <div class="alert alert-warning" role="alert">
                    Expires in: {{comment['end']}}
                </div>
                <button type="submit" class="btn btn-primary" v-on:click="removeComment">Remove Comment</button>
            </div>
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
                error: null
            }
        },

        methods: {
            send: function () {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + "/api/comment",
                    headers: this.$parent.$data['config'],
                    data: {
                        end: this.form.end,
                        msg: this.form.msg
                    }
                }).then(response => {
                    if (response.status === 200) {
                        this.fetchComment();
                    }
                });
            },

            removeComment: function () {
                axios({
                    method: 'delete',
                    url: this.$root.$data['api'] + "/api/comment",
                    headers: this.$parent.$data['config'],
                }).then(response => {
                    if (response.status === 200) {
                        this.comment = null;
                    }
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
                    this.error = "Form cannot be empty!";
                    e.preventDefault();
                    return;
                }
                if (this.form.msg === '') {
                    this.error = "No message provided!";
                    e.preventDefault();
                    return;
                }
                if (!this.isGoodDate(this.form.end)) {
                    this.error = "Invalid date!";
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
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/AdminComment";
@import "~bootstrap/scss/bootstrap.scss";
</style>