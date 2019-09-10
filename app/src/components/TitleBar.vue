<template>
    <nav class="navbar navbar-expand navbar-dark">
        <a class='navbar-brand'>
            <img v-bind:src="this.$root.$data['LOGO_LINK']"
                 alt="logo" height="44" />
        </a>
        <nav class="navbar-nav mr-auto">
            <img class="ml-2" v-bind:src="this.$root.$data['DEPARTMENT_LINK']"
                 alt="cs_logo" height="55" />
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

            r.style.setProperty("--navbar", this.$root.$data['NAVBAR']);
        },

        methods: {
            getTime: function () {
                setInterval(() => {
                    let date = new Date();

                    this.date = date.toLocaleString()
                })
            },

            getRoom: function () {
                fetch(this.$root.$data['api'] + this.$root.$data['port'] + "/api/room/" + this.$root.$data['room'])
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