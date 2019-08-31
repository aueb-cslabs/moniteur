<template>
    <nav class="navbar navbar-expand navbar-dark">
        <a class='navbar-brand'>
            <img v-bind:src="'https://www.aueb.gr/press/logos/1_AUEB-pantone-HR.jpg'"
                 alt="logo" height="44" />
        </a>
        <nav class="navbar-nav mr-auto">
            <a class="navbar-brand">AUEB CSLabs</a>
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
                    let date = new Date()

                    this.date = date.toLocaleString()
                })
            }
        }
    }
</script>

<style lang="scss">
@import "../css/TitleBar.scss";
</style>