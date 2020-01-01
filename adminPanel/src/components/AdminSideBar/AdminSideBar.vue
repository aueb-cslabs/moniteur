<template>
    <div id="sidebar" class="sidebar">
        <router-link to="/" class="nav-item" v-on:click.native="showSidebar">{{$t("message.adminHome")}}</router-link>
        <router-link to="/announcement" class="nav-item" v-on:click.native="showSidebar">{{$t("message.adminBarGA")}}</router-link>
        <router-link to="/roomAnnouncement" class="nav-item" v-on:click.native="showSidebar">{{$t("message.adminBarRA")}}</router-link>
        <router-link to="/comment" class="nav-item" v-on:click.native="showSidebar">{{$t("message.adminBarGC")}}</router-link>
        <router-link to="/users" class="nav-item" v-on:click.native="showSidebar">{{$t("message.adminBarUM")}}</router-link>
        <router-link to="/calendar" class="nav-item" v-on:click.native="showSidebar">{{$t("message.adminBarAC")}}</router-link>
        <router-link to="/unscheduled" class="nav-item" v-on:click.native="showSidebar">{{$t("message.adminBarUL")}}</router-link>
        <router-link to="/settings" class="nav-item" v-on:click.native="showSidebar">{{$t("message.adminBarSettings")}}</router-link>
    </div>
</template>

<script>
    import functions from "../../functions";

    export default {
        name: 'AdminSideBar',

        methods: {
            showSidebar: functions.showSidebar
        }
    }
</script>

<style>
@import "./AdminSideBar.scss";
</style>