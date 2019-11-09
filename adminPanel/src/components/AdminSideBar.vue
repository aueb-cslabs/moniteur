<template>
    <div id="sidebar" class="sidebar">
        <a class="nav-item" v-on:click="this.$parent.showAnnouncement">{{$t("message.adminBarGA")}}</a>
        <a class="nav-item" v-on:click="this.$parent.showRoomAnnouncement">{{$t("message.adminBarRA")}}</a>
        <a class="nav-item" v-on:click="this.$parent.showComment">{{$t("message.adminBarGC")}}</a>
        <a class="nav-item" v-on:click="this.$parent.showUserManagement">{{$t("message.adminBarUM")}}</a>
        <a class="nav-item" v-on:click="this.$parent.showAcademicCalendar">{{$t("message.adminBarAC")}}</a>
<!--        <a class="nav-item" v-on:click="this.$parent.showExamsCalendar">{{$t("message.adminBarEC")}}</a>-->
        <a class="nav-item" v-on:click="this.$parent.showSettings">{{$t("message.adminBarSettings")}}</a>
    </div>
</template>

<script>
    export default {
        name: 'AdminSideBar',
    }
</script>

<style>
@import "../scss/AdminSideBar.scss";
</style>