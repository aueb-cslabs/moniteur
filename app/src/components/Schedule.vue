<template>
    <div class="wrapper" v-if="normal">
        <div id="current" class="text-center current schedule">
            <h3>Τώρα <i class="fas fa-users-class"></i></h3>
            <p class="center common fade-in" v-if="current['now'] != null">
                <u>
                    {{time['now_start'].getUTCHours()}}:{{getMinutes(this.time['now_start'])}}
                    - {{time['now_end'].getUTCHours()}}:{{getMinutes(this.time['now_end'])}}
                </u><br>
                {{current['now']['title']}}<br>
                {{current['now']['host']}}
            </p>
            <p class="center common fade-in subject" v-else-if="isExam">
                Δεν πραγματοποιείται εξέταση.
            </p>
            <p class="center common fade-in subject" v-else>
                Δεν πραγματοποιείται μάθημα.
            </p>
        </div>

        <div id="next" class="text-center next schedule">
            <h3>Επόμενο <i class="fas fa-users-class nextsub"></i></h3>
            <p class="center common fade-in" v-if="current['next'] != null">
                <u>
                    {{time['next_start'].getUTCHours()}}:{{getMinutes(this.time['next_start'])}}
                    - {{time['next_end'].getUTCHours()}}:{{getMinutes(this.time['next_end'])}}
                </u><br>
                {{current['next'][0]['title']}}<br>
                {{current['next'][0]['host']}}
            </p>
            <p class="center common fade-in subject" v-else-if="isExam">
                Δεν θα πραγματοποιείται εξέταση.
            </p>
            <p class="center common fade-in subject" v-else>
                Δεν θα πραγματοποιείται μάθημα.
            </p>
        </div>
    </div>
    <div class="wrapper" v-else>
        <div class="center-message message mt-0">
            <h3 class="common fade-on subject" v-if="isBreak">Αργία ή διακοπές. Καλή ξεκούραση!</h3>
            <h3 class="center-message common fade-on subject" v-if="isWeekend">Καλό Σαββατοκύριακο!</h3>
        </div>
    </div>
</template>

<script>
    import { EventBus} from "./EventBus";

    export default {
        name: "Schedule",

        data() {
            return {
                current: [],
                isExam: {},
                time: {},
                isWeekend: {},
                isBreak: {},
                normal: {}
            }
        },

        created() {
            this.checkExam();
            this.checkWeekend();
            this.checkBreak();
            this.checkDefault();

            if (this.isExam) {
                this.fetchExamSched();
            } else {
                this.fetchNormSched();
                EventBus.$emit('exam', false);
            }
        },

        methods: {
            /* Checks if we are in examination period */
            checkExam: function () {
                setInterval(() => {
                    fetch(this.$root.$data['api'] + this.$root.$data['port'] + "/api/calendarInfo")
                        .then(res => res.json())
                        .then(json => {
                            this.isExam = json['exams'];
                        });
                }, 86400);
            },

            /* Checks if it is weekend */
            checkWeekend: function () {
                setInterval(() => {
                    fetch(this.$root.$data['api'] + this.$root.$data['port'] + "/api/calendarInfo")
                        .then(res => res.json())
                        .then(json => {
                            this.isWeekend = json['weekend'];
                        });
                }, 7200000);
            },

            /* Checks if there is a break or a national holiday */
            checkBreak: function () {
                setInterval(() => {
                    fetch(this.$root.$data['api'] + this.$root.$data['port'] + "/api/calendarInfo")
                        .then(res => res.json())
                        .then(json => {
                            this.isBreak = json['break'];
                        })
                }, 7200000);
            },

            /* Checks for normal period */
            checkDefault: function () {
                setInterval(() => {
                    this.normal = !this.isWeekend && !this.isBreak
                }, 7200000);
            },

            /* Fetches examination schedule */
            fetchExamSched: function() {
                setInterval(() => {
                    fetch(this.$root.$data['api'] + this.$root.$data['port'] + "/api/exams/" + this.$root.$data['room'] + "/now")
                        .then(response => response.json())
                        .then(json => {
                            this.current = json;

                            this.checkNext(this.current);
                            this.convertSecToTime();

                            if(this.current['now'] != null) {
                                EventBus.$emit('exam', this.isExam);
                            } else {
                                EventBus.$emit('exam', false);
                            }
                        })
                }, 30000)
            },

            /* Fetches normal schedule */
            fetchNormSched: function() {
                setInterval(() => {
                    fetch(this.$root.$data['api'] + this.$root.$data['port'] + "/api/schedule/" + this.$root.$data['room']+ "/now")
                        .then(response => response.json())
                        .then(json => {
                            this.current = json;
                            this.checkNext(this.current);
                            this.convertSecToTime();
                        })
                }, 30000)
            },

            /* Checks if it's morning. If not don't render next
             * class until the time is 8AM
             */
            checkNext: function (schedule) {
                if(schedule['next'] != null) {
                    let date =  new Date();
                    if(date.getHours() < 8) {
                        schedule['next'] = null;
                    }
                }
            },

            convertSecToTime: function() {
                let d = new Date(0);
                if(this.current['now'] != null) {
                    d = new Date(0);
                    d.setSeconds(this.current['now']['start']);
                    this.time['now_start'] = d;


                    d = new Date(0);
                    d.setSeconds(this.current['now']['end']);
                    this.time['now_end'] = d;
                }

                if(this.current['next'] != null) {
                    d = new Date(0);
                    d.setSeconds(this.current['next'][0]['start']);
                    this.time['next_start'] = d;

                    d = new Date(0);
                    d.setSeconds(this.current['next'][0]['end']);
                    this.time['next_end'] = d;
                }
            },

            /* Returns minutes with leading zero
             * if minutes < 10
             */
            getMinutes: function (time) {
                let min = time.getUTCMinutes();
                return (min < 10) ? ("0" + min) : min;
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/Schedule.scss";
</style>