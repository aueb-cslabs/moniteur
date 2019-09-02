<template>
    <nav class="navbar navbar-expand navbar-dark">
        <a class='navbar-brand'>
            <img v-bind:src="'https://www.aueb.gr/press/logos/1_AUEB-pantone-HR.jpg'"
                 alt="logo" height="44" />
        </a>
        <nav class="navbar-nav mr-auto">
            <a class="navbar-brand mt-2">AUEB CSLabs</a>
            <img class="ml-2" v-bind:src="'https://www.dept.aueb.gr/schools_department_photos/cs.png'"
                 alt="cs_logo" height="55" />
        </nav>
        <a class="navbar-brand">
            {{date}}
        </a>
    </nav>
</template>

<script>
    export default {
        name: "TitleBar",

        data() {
            return {
                date: ''
            }
        },

        created() {
            this.getTime();
        },

        methods: {
            getTime: function () {
                setInterval(() => {
                    let date = new Date();

                    this.date = date.toLocaleString()
                })
            }
        }
    }
</script>

<style lang="scss">
@import "../css/TitleBar.scss";
</style>