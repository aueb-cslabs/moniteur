<template>
    <div>
        <div class="alert alert-primary" role="alert">
            <h3>{{this.$props.type}} {{this.$props.announcement}}</h3>
        </div>
        <hr>
        <div class="alert alert-warning" role="alert">
            {{this.$t("message.adminExpires")}} {{this.$props.expiration}}
        </div>
    </div>
</template>

<script>
    export default {
        name: "ShowAnnouncement",

        props: {
            type: String,
            announcement: String,
            expiration: String
        }
    }
</script>

<style lang="scss">
@import "ShowAnnouncement";
</style>