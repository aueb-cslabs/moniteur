<template>
    <nav class="navbar navbar-expand navbar-dark">
        <a class='navbar-brand'>
            <img v-bind:src="this.$root.$data['logo_url']"
                 alt="logo" height="60" />
        </a>
        <nav class="navbar-nav mr-auto">
            <img class="ml-2" v-bind:src="this.$root.$data['secondary_logo_url']"
                 alt="cs_logo" height="60" />
            <a class="ml-4 navbar-brand mt-2">{{room}}</a>
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
                date: '',
                room: ''
            }
        },

        created() {
            this.getTime();
            this.getRoom();

            let r = document.documentElement;

            r.style.setProperty("--navbar-bg-color", this.$root.$data['navbar_background_color']);
            r.style.setProperty("--navbar-color", this.$root.$data['navbar_color']);
        },

        methods: {
            getTime: function () {
                setInterval(() => {
                    let date = new Date();
                    this.date = date.toLocaleString()
                })
            },

            getRoom: function () {
                fetch(this.$root.$data['api'] + "/api/room/" + this.$route.params.id)
                    .then(res => res.json())
                    .then(json => {
                        this.room = json.toString()
                    })
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/TitleBar.scss";
</style>