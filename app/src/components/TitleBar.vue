<template>
    <nav class="navbar navbar-expand navbar-dark">
        <a class='navbar-brand'>
            <img v-bind:src="'https://www.aueb.gr/press/logos/1_AUEB-pantone-HR.jpg'"
                 alt="logo" height="44" />
        </a>
        <nav class="navbar-nav mr-auto">
            <img class="ml-2" v-bind:src="'https://www.dept.aueb.gr/schools_department_photos/cs.png'"
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
        },

        methods: {
            getTime: function () {
                setInterval(() => {
                    let date = new Date();

                    this.date = date.toLocaleString()
                })
            },

            getRoom: function () {
                fetch(this.$root.$data['api']+":27522/api/room/" + this.$root.$data['room'])
                    .then(res => res.json())
                    .then(json => {
                        this.room = json.toString()
                    })
            }
        }
    }
</script>

<style lang="scss">
@import "../css/TitleBar.scss";
</style>