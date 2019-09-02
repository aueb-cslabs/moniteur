<template>
    <div class="wrapper">
        <div id="current" class="text-center current schedule">

            <h3>Τώρα <i class="fas fa-users-class"></i></h3>
            <p class="mt-2 common fade-in" v-if="current['now'] != null">
                <u>{{current['now']['start']/60/60}} - {{current['now']['end']/60/60}}</u><br>
                {{current['now']['title']}}<br>
                {{current['now']['host']}}
            </p>
            <p class="mt-4 pt-2 common fade-in subject" v-else>
                Δεν πραγματοποιείται μάθημα.
            </p>
        </div>

        <div id="next" class="text-center next schedule">
            <h3>Επόμενο <i class="fas fa-users-class nextsub"></i></h3>
            <p class="mt-2 common fade-in" v-if="current['next'] != null">
                <u>{{current['next'][0]['start']/60/60}} - {{current['next'][0]['end']/60/60}}</u><br>
                {{current['next'][0]['title']}}<br>
                {{current['next'][0]['host']}}
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
            this.getNow();
        },

        methods: {
            getNow: function () {
                setInterval(() => {
                    fetch("http://localhost:1323/api/schedule/a/now")
                        .then(response => response.json())
                        .then(json =>  {
                            this.current = json;
                            this.checkNext(this.current);
                        })
                }, 5000)
            },

            checkNext: function (schedule) {
                if(schedule['now'] != null) {
                    if(schedule['now']['end'] !== schedule['next'][0]['start']) {
                        schedule['next'] = null;
                    }
                }
            }
        }
    }
</script>

<style lang="scss">
    @import "../css/Schedule.scss";
</style>