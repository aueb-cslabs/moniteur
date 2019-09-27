<template>
    <div id="sidebar" class="sidebar">
        <a class="nav-item" v-on:click="this.$parent.showAnnouncement">{{$t("message.adminBarGA")}}</a>
        <a class="nav-item" v-on:click="this.$parent.showRoomAnnouncement">{{$t("message.adminBarRA")}}</a>
        <a class="nav-item" v-on:click="this.$parent.showComment">{{$t("message.adminBarGC")}}</a>
    </div>
</template>

<script>
    export default {
        name: 'AdminSideBar',

        created() {

        }
    }
</script>

<style>
@import "../scss/AdminSideBar.scss";
</style>