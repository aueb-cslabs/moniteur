<template>
    <div>
        <div class="pt-5">
            <div>
                <form @submit="checkForm">
                    <div class="form-group">
                        <label for="message">Announcement</label>
                        <input type="text" class="form-control" id="message" v-model="form.msg" placeholder="Message">
                    </div>
                    <div class="form-group">
                        <label for="expiration">Expiration Date</label>
                        <input type="text" class="form-control" id="expiration" v-model="form.end" placeholder="DD/MM/YYYY">
                    </div>
                    <button type="submit" class="btn btn-primary">Send announcement</button>
                    <p class="error">{{error}}</p>
                </form>
            </div>
            <div class="mt-5" v-if="announcement != null">
                <div class="alert alert-primary" role="alert">
                    <h3>Current Announcement: {{announcement['msg']}}</h3>
                </div>
                <hr>
                <div class="alert alert-warning" role="alert">
                    Expires in: {{announcement['end']}}
                </div>
                <button type="submit" class="btn btn-primary" v-on:click="removeAnnouncement">Remove Announcement</button>
            </div>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: 'AdminAnnouncement',

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
                error: null
            }
        },

        methods: {
            send: function () {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + "/api/announcement",
                    headers: this.$parent.$data['config'],
                    data: {
                        end: this.form.end,
                        msg: this.form.msg
                    }
                }).then(response => {
                    if (response.status === 200) {
                        this.fetchAnnouncement();
                    }
                });
            },

            removeAnnouncement: function () {
                axios({
                    method: 'delete',
                    url: this.$root.$data['api'] + "/api/announcement",
                    headers: this.$parent.$data['config'],
                }).then(response => {
                    if (response.status === 200) {
                        this.announcement = null;
                    }
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
                    this.error = "Form cannot be empty!";
                    return;
                }
                if (this.form.msg === '') {
                    this.error = "No message provided!";
                    return;
                }
                if (this.form.end === '') {
                    this.error = "No date provided!";
                    return;
                }
                this.error = null;
                this.send();
                e.preventDefault();
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/AdminAnnouncement";
@import "~bootstrap/scss/bootstrap.scss";
</style>