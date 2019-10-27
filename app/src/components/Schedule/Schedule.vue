<template>
    <div class="wrapper" v-if="normal">
        <div class="text-center current schedule">
            <h3> {{ $t("message.nowMsg") }} <i class="fas fa-chalkboard-teacher"></i></h3>
            <p class="center-message common fade-in" v-if="current['now'] != null">
                <u>
                    {{time['now_start'].getUTCHours()}}:{{getMinutes(this.time['now_start'])}}
                    - {{time['now_end'].getUTCHours()}}:{{getMinutes(this.time['now_end'])}}
                </u><br>
                {{current['now']['title']}}<br>
                {{current['now']['host']}}
            </p>
            <p class="center-message common fade-in" v-else-if="isExam">
                {{ $t("message.noExamNow") }}
            </p>
            <p class="center-message common fade-in" v-else>
                {{ $t("message.noLessonNow") }}
            </p>
        </div>
        <div id="next" class="text-center next schedule">
            <h3> {{ $t("message.nextMsg") }} <i class="fas fa-chalkboard-teacher next-icon"></i></h3>
            <p class="center-message common fade-in" v-if="current['next'] != null">
                <u>
                    {{time['next_start'].getUTCHours()}}:{{getMinutes(this.time['next_start'])}}
                    - {{time['next_end'].getUTCHours()}}:{{getMinutes(this.time['next_end'])}}
                </u><br>
                {{current['next'][0]['title']}}<br>
                {{current['next'][0]['host']}}
            </p>
            <p class="center-message common fade-in" v-else-if="isExam">
                {{ $t("message.noExamNext") }}
            </p>
            <p class="center-message common fade-in" v-else>
                {{ $t("message.noLessonNext") }}
            </p>
        </div>
    </div>
    <div class="wrapper" v-else>
        <div class="message">
            <h3 class="center-message common fade-in" v-if="isBreak">{{ $t("message.holiday") }}</h3>
            <h3 class="center-message common fade-in" v-if="isWeekend">{{ $t("message.weekend") }}</h3>
        </div>
    </div>
</template>

<script>
    import axios from 'axios';

    export default {
        name: "Schedule",

        data() {
            return {
                current: [],
                isExam: false,
                time: {},
                isWeekend: false,
                isBreak: false,
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
                this.$emit('exam', false);
            }
        },

        methods: {
            /* Checks if we are in examination period */
            checkExam: function () {
                setInterval(() => {
                    axios.get(this.$root.$data['api']+"/api/calendarInfo")
                        .then(res => this.isExam = res.data['exams']);
                }, 60000);
            },

            /* Checks if it is weekend */
            checkWeekend: function () {
                setInterval(() => {
                    axios.get(this.$root.$data['api']+"/api/calendarInfo")
                        .then(res => this.isWeekend = res.data['weekend']);
                }, 60000);
            },

            /* Checks if there is a break or a national holiday */
            checkBreak: function () {
                setInterval(() => {
                    axios.get(this.$root.$data['api']+"/api/calendarInfo")
                        .then(res => this.isBreak = res.data['break']);
                }, 60000);
            },

            /* Checks for normal period */
            checkDefault: function () {
                setInterval(() => {
                    this.normal = !this.isWeekend && !this.isBreak
                }, 60000);
            },

            /* Fetches examination schedule */
            fetchExamSched: function() {
                setInterval(() => {
                    axios.get(this.$root.$data['api']+"/api/exams/" + this.$route.params.id)
                        .then(res => {
                            this.current = res.data;

                            this.checkNext(this.current);
                            this.convertSecToTime();

                            if(this.current['now'] != null) {
                                this.$emit('exam', this.isExam);
                            } else {
                                this.$emit('exam', false);
                            }
                        });
                }, 30000)
            },

            /* Fetches normal schedule */
            fetchNormSched: function() {
                setInterval(() => {
                    axios.get(this.$root.$data['api']+"/api/schedule/" + this.$route.params.id + "/now")
                        .then(res => {
                            this.current = res.data;
                            this.checkNext(this.current);
                            this.convertSecToTime();
                        });
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
    @import "Schedule";
</style>