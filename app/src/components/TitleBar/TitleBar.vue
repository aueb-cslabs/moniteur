<template>
    <nav class="navbar navbar-expand navbar-dark">
        <a class='navbar-brand'>
            <img v-bind:src="logo"
                 alt="logo" height="90" />
        </a>
        <nav class="navbar-nav mr-auto">
            <img class="ml-2" v-bind:src="secondary_logo"
                 alt="cs_logo" height="95" />
            <a class="ml-4 navbar-brand title">{{$t("message.room")}} {{room}}</a>
        </nav>
        <a class="navbar-brand title">
            {{date}}
        </a>
    </nav>
</template>

<script>
    import config from "../../config";

    export default {
        name: "TitleBar",

        data() {
            return {
                logo: config.logo_url,
                secondary_logo: config.secondary_logo_url,

                date: '',
                room: ''
            }
        },

        created() {
            this.getTime();
            this.getRoom();

            let r = document.documentElement;

            r.style.setProperty("--navbar-bg-color", config.navbar_background_color);
            r.style.setProperty("--navbar-color", config.navbar_color);
        },

        methods: {
            getTime: function () {
                setInterval(() => {
                    let date = new Date();
                    this.date = date.toLocaleString(undefined, {hour12: false })
                })
            },

            getRoom: function () {
                fetch(config.api + "/api/room/" + config.room)
                    .then(res => res.json())
                    .then(json => {
                        this.room = json.toString()
                    })
            }
        }
    }
</script>

<style lang="scss">
    @import "TitleBar";
</style>