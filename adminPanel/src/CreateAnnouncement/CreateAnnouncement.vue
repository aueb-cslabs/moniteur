<template>
    <div>
        <div>
            <form @submit="checkForm">
                <div class="form-group">
                    <label for="message">{{$t("message.adminAnnForm1")}}</label>
                    <input type="text"
                           class="form-control"
                           id="message"
                           v-model="form.msg"
                           v-bind:placeholder="this.$t('message.adminAnnMsg')"
                           required>
                </div>
                <div class="form-group">
                    <label for="expiration">{{$t("message.adminAnnExpire")}}</label>
                    <input type="datetime-local"
                           class="form-control"
                           id="expiration"
                           v-model="form.end"
                           v-bind:placeholder="this.$t('message.adminAnnDateForm')"
                           required>
                </div>
                <button type="submit" class="btn btn-primary">{{$t("message.adminAnnSend")}}</button>
            </form>
        </div>
        <div>
            <ErrorPopup ref="error"/>
        </div>
    </div>
</template>

<script>
    import ErrorPopup from "../ErrorPopup/ErrorPopup";
    export default {
        name: "CreateAnnouncement",
        components: {ErrorPopup},
        data() {
            return {
                form: {
                    end: '',
                    msg: ''
                },
            }
        },

        methods: {
            checkForm: function (e) {
                if (this.form.msg === '' && this.form.end === '') {
                    this.$refs.error.setError(this.$t('message.adminEmptyForm'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.msg === '') {
                    this.$refs.error.setError(this.$t('message.adminNoMsg'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.form.end === '') {
                    this.$refs.error.setError(this.$t('message.adminNoDate'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                this.$parent.send(this.form);
                this.form.msg = '';
                this.form.end = '';
                e.preventDefault();
            }
        }
    }
</script>

<style scoped>

</style>