<template>
    <Admin/>
</template>

<script>
    import Admin from "../Admin/Admin";

    const config = require('electron').remote.getGlobal('config');

    export default {

        name: 'app',
        components: {
            Admin
        },

        created() {
            let r = document.documentElement;
            r.style.setProperty("--background", config.background_color);
            r.style.setProperty("--navbar-bg-color", config.navbar_background_color);
            r.style.setProperty("--navbar-color", config.navbar_color);
        },
    }
</script>

<style>
@import "./App.scss";
</style>
