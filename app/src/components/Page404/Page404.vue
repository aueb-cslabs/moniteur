<template>
    <div class="d-flex justify-content-center text-center">
        <div>
            <i class="fas fa-hammer fa-5x mb-3 text-danger"></i>
            <h1>Oops!</h1>
            Something went wrong <br>
            Our team is trying to fix it! <br><br>
            Are you sure you entered a classroom?<br>
            Try <code>moniteur.aueb.gr/a</code> for classroom A!<br><br>
            For more information take a look at our documentation: <code>docs.moniteur.aueb.gr</code>
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