<template>
    <div class="d-flex justify-content-center text-center">
        <div>
            <img src="https://media.giphy.com/media/WrNfErHio7ZAc/giphy.gif"
                 class="rounded mb-3" height="200" alt="facepalm-gif">
            <h1>Oops!</h1>
            <p v-html="$t('message.oops')"></p>
        </div>
    </div>
</template>

<script>
    import config from "../../config";

    export default {
        name: "Page404",

        created() {
            let r = document.documentElement;
            r.style.setProperty("--background", config.background_color);
        }
    }
</script>

<style lang="scss" scoped>
    @import "~bootstrap/scss/bootstrap";
</style>