<template>
    <div class="wrapper" v-if="!this.isWeekend && !this.isBreak">
        <div class="text-center current schedule">
            <p class="title"> {{ $t("message.nowMsg") }} <i class="fas fa-chalkboard-teacher"></i></p>
            <p class="center-message common fade-in" v-if="current['now'] != null">
                <u>
                    {{nowStart}} - {{nowEnd}}
                </u>
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
            <p class="title"> {{ $t("message.nextMsg") }} <i class="fas fa-chalkboard-teacher next-icon"></i></p>
            <p class="center-message common fade-in" v-if="current['next'] != null">
                <u>
                    {{nextStart}} - {{nextEnd}}
                </u>
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
    import config from "../../config";
    import axios from 'axios';

    export default {
        name: "Schedule",

        data() {
            return {
                current: [],
                isExam: false,
                isWeekend: false,
                isBreak: false,
                examIntervalId: null,
                normIntervalId: null
            }
        },

        created() {
            this.checkExam();
            this.checkWeekend();
            this.checkBreak();
        },

        computed: {
            nowStart() {
                let date = new Date(this.current['now']['start'] * 1000);
                return date.getHours() + ":" +
                    this.getCorrectedMinutes(date.getMinutes())
            },
            nowEnd() {
                let date = new Date(this.current['now']['end'] * 1000);
                return date.getHours() + ":" +
                    this.getCorrectedMinutes(date.getMinutes())
            },
            nextStart() {
                let date = new Date(this.current['next'][0]['start'] * 1000);
                return date.getHours() + ":" +
                    this.getCorrectedMinutes(date.getMinutes())
            },
            nextEnd() {
                let date = new Date(this.current['next'][0]['end'] * 1000);
                return date.getHours() + ":" +
                    this.getCorrectedMinutes(date.getMinutes())
            }
        },

        methods: {
            /* Checks if we are in examination period */
            checkExam: function () {
                setInterval(() => {
                    axios.get(config.api + "/api/calendarInfo")
                        .then(res => this.isExam = res.data['exams'])
                        .finally(() => {
                            if(this.isExam) {
                                if(this.normIntervalId != null) {
                                    clearInterval(this.normIntervalId);
                                    this.normIntervalId = null
                                }
                                if(this.examIntervalId == null)
                                    this.fetchExamSched();
                            } else {
                                if(this.examIntervalId != null) {
                                    clearInterval(this.examIntervalId);
                                    this.examIntervalId = null
                                }
                                if(this.normIntervalId == null)
                                    this.fetchNormSched();
                                this.$parent.$emit('exam', false)
                            }
                        });
                }, 30000);
            },

            /* Checks if it is weekend */
            checkWeekend: function () {
                setInterval(() => {
                    axios.get(config.api + "/api/calendarInfo")
                        .then(res => this.isWeekend = res.data['weekend']);
                }, 60000);
            },

            /* Checks if there is a break or a national holiday */
            checkBreak: function () {
                setInterval(() => {
                    axios.get(config.api + "/api/calendarInfo")
                        .then(res => this.isBreak = res.data['break']);
                }, 60000);
            },

            /* Fetches examination schedule */
            fetchExamSched: function() {
                this.examIntervalId = setInterval(() => {
                    axios.get(config.api + "/api/exams/" + this.$route.params.room + "/now")
                        .then(res => {
                            this.current = res.data;

                            this.checkNext(this.current);

                            /* If now is not null then an exam is taking place
                             * Fire up exam sign!
                             */
                            if(this.current['now'] != null)
                                this.$parent.$emit('exam', true);
                            else
                                this.$parent.$emit('exam', false)
                        });
                }, 30000)
            },

            /* Fetches normal schedule */
            fetchNormSched: function() {
                this.normIntervalId = setInterval(() => {
                    axios.get(config.api + "/api/schedule/" + this.$route.params.room + "/now")
                        .then(res => {
                            this.current = res.data;
                            this.checkNext(this.current);
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

            /* Returns minutes with leading zero
             * if minutes < 10
             */
            getCorrectedMinutes: function (min) {
                return (min < 10) ? ("0" + min) : min;
            }
        }
    }
</script>

<style lang="scss">
    @import "~bootstrap/scss/bootstrap";
    @import "Schedule";
    @import "Schedule_phone";
    @import "Schedule_tablet";
</style>