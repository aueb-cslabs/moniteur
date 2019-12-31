<template>
    <div v-if="!this.$root.$data.authToken.auth" class="login">
        <Login/>
    </div>
    <div v-else class="actions">
        <AdminBar/>
        <div class="admin">
            <AdminSideBar/>
            <div id="main">
                <router-view></router-view>
            </div>
        </div>
    </div>
</template>

<script>
    import Login from "../Login/Login";
    import AdminBar from "../AdminBar/AdminBar";
    import AdminSideBar from "../AdminSideBar/AdminSideBar";

    export default {

        name: 'Administration',

        components: {
            Login,
            AdminBar,
            AdminSideBar
        }
    }

</script>

<style lang="scss">
@import "./Admin.scss";
@import "~bootstrap/scss/bootstrap.scss";
</style>