<template>
    <form>
        <div class="form-group">
            <label for="message">{{$t("message.adminAnnForm1")}}</label>
            <input type="text"
                   class="form-control"
                   id="message"
                   :value="msg"
                   v-bind:placeholder="this.$t('message.adminAnnMsg')"
                   @input="update('msg', $event.target.value)"
                   required>
        </div>
        <div class="form-group">
            <label for="expiration">{{$t("message.adminAnnExpire")}}</label>
            <input type="datetime-local"
                   class="form-control"
                   id="expiration"
                   :value="end"
                   v-bind:placeholder="this.$t('message.adminAnnDateForm')"
                   @input="update('end', $event.target.value)"
                   required>
        </div>
    </form>
</template>

<script>
    export default {
        name: "CreateAnnouncement",
        props: {
            msg: String,
            end: String,
            form: {
                msg: String,
                end: String
            }
        },

        methods: {
            update(key, value) {
                this.$emit('input', {...this.$props.form, [key]: value })
            },
        }
    }
</script>

<style scoped lang="scss">
@import "./CreateAnnouncement";
</style>