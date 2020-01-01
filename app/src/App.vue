<template>
    <Moniteur class="moniteur"></Moniteur>
</template>

<script>
    import Moniteur from "./components/Moniteur";

    export default {
        name: 'app',

        components: {
            Moniteur
        }
    }

</script>

<style lang="scss">
    @import "~bootstrap/scss/bootstrap.scss";
    @import "~@fortawesome/fontawesome-free/css/all.css";
    @import "App";
</style>
