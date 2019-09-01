<template v-if="current">
    <div class="wrapper">
        <div id="current" class="text-center current schedule">

            <h3>Τώρα <i class="fas fa-users-class"></i></h3>
            <p class="mt-4 pt-2 common fade-in" v-if="current['now'] != null">
                {{current['now']}}
            </p>
            <p class="mt-4 pt-2 common fade-in subject" v-else>
                Δεν πραγματοποιείται μάθημα.
            </p>
        </div>

        <div id="next" class="text-center next schedule">
            <h3>Επόμενο <i class="fas fa-users-class nextsub"></i></h3>
            <p class="mt-4 pt-2 common fade-in" v-if="current['now'] != null">
                {{current['now']}}
            </p>
            <p class="mt-4 pt-2 common fade-in subject" v-else>
                Δεν θα πραγματοποιείται μάθημα.
            </p>
        </div>
    </div>
</template>

<script>
    export default {
        name: "Schedule",

        data() {
            return {
                current: []
            }
        },

        created() {
            this.getNow()
        },

        methods: {
            getNow: function () {
                setInterval(() => {
                    fetch("http://localhost:1323/api/schedule/a/now")
                        .then(response => response.json())
                        .then(json => this.current = json)
                }, 5000)
            }
        }
    }
</script>

<style lang="scss">
    @import "../css/Schedule.scss";
</style>