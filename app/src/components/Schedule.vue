<template v-if="current">
    <div class="wrapper">
        <div id="current" class="text-center w-25 current schedule">

            <h3>Τώρα <i class="fas fa-users-class"></i></h3>
            <p class="mt-4 pt-2" v-if="current['now'] != null">
                {{current['now']}}
            </p>
            <p class="mt-4 pt-2 subject" v-else>
                Δεν πραγματοποιείται μάθημα.
            </p>
        </div>

        <div id="next" class="text-center w-25 next schedule">
            <h3>Επόμενο <i class="fas fa-users-class nextsub"></i></h3>
            <p class="mt-4 pt-2" v-if="current['now'] != null">
                {{current['now']}}
            </p>
            <p class="mt-4 pt-2 subject" v-else>
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
                fetch("http://localhost:1323/api/schedule/a/now")
                    .then(response => response.json())
                    .then(json => {
                        this.current = json;
                    })
            }
        }
    }
</script>

<style lang="scss">
    @import "../css/Schedule.scss";
</style>