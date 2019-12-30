<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.ecGeneral")}}</h2>
            <div class="pt-3">
                <div class="row">
                    <div class="col">
                        <h5>{{$t('message.calWinterS')}}</h5>
                        <PeriodDefinition
                                v-model="Calendar.Semesters.Winter"
                                :period="Calendar.Semesters.Winter"
                                :start="Calendar.Semesters.Winter.Start"
                                :end="Calendar.Semesters.Winter.End"/>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSprintS')}}</h5>
                        <PeriodDefinition
                                v-model="Calendar.Semesters.Sprint"
                                :period="Calendar.Semesters.Sprint"
                                :start="Calendar.Semesters.Sprint.Start"
                                :end="Calendar.Semesters.Sprint.End"/>
                    </div>
                </div>
            </div>
            <h2 class="pt-3">{{$t("message.calExams")}}</h2>
            <div class="pt-3">
                <div class="row">
                    <div class="col">
                        <h5>{{$t('message.calWin')}}</h5>
                        <PeriodDefinition
                                v-model="Calendar.Exams.Winter"
                                :period="Calendar.Exams.Winter"
                                :start="Calendar.Exams.Winter.Start"
                                :end="Calendar.Exams.Winter.End"/>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSpr')}}</h5>
                        <PeriodDefinition
                                v-model="Calendar.Exams.Sprint"
                                :period="Calendar.Exams.Sprint"
                                :start="Calendar.Exams.Sprint.Start"
                                :end="Calendar.Exams.Sprint.End"/>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSep')}}</h5>
                        <PeriodDefinition
                                v-model="Calendar.Exams.September"
                                :period="Calendar.Exams.September"
                                :start="Calendar.Exams.September.Start"
                                :end="Calendar.Exams.September.End"/>
                    </div>
                </div>
            </div>
            <h2 class="pt-3">{{$t("message.calBreaks")}}</h2>
            <div class="pt-3">
                <div class="row">
                    <div class="col">
                        <h5>{{$t('message.calWinterS')}}</h5>
                        <PeriodDefinition
                                v-model="Calendar.Breaks.Winter"
                                :period="Calendar.Breaks.Winter"
                                :start="Calendar.Breaks.Winter.Start"
                                :end="Calendar.Breaks.Winter.End"/>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSprintS')}}</h5>
                        <PeriodDefinition
                                v-model="Calendar.Breaks.Sprint"
                                :period="Calendar.Breaks.Sprint"
                                :start="Calendar.Breaks.Sprint.Start"
                                :end="Calendar.Breaks.Sprint.End"/>
                    </div>
                </div>
                <h5>{{$t('message.calVarious')}}</h5>
                <div>
                    <label>{{$t("message.calAddDate")}}</label>
                    <div>
                        <form>
                            <div class="row">
                                <div class="col-md-8 col-lg-10 form-group">
                                    <input type="date" class="form-control" v-model="date" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                                </div>
                                <div class="col form-group">
                                    <button @click="addDate" class="btn btn-primary float-right">{{$t("message.umAdd")}}</button>
                                </div>
                            </div>
                        </form>
                    </div>
                    <div class="mb-4">
                        <h5>{{$t("message.calRemoveDate")}}</h5>
                        <div>
                            <label>{{$t("message.calSelectDate")}}</label>
                            <b-form-select v-model="dateOption" :options="Calendar.Breaks.Various"></b-form-select>
                        </div>
                        <div class="row">
                            <div class="col">
                                <b-button @click="removeDate" variant="danger" class="mt-3">{{$t("message.calRemoveDate")}}</b-button>
                                <b-button @click="removeAll" variant="danger" class="ml-3 mt-3">{{$t("message.calRemoveAll")}}</b-button>
                                <b-button @click="send" variant="primary" class="ml-3 mt-3 float-right">{{this.$t("message.calSend")}}</b-button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <SuccessPopup ref="success"/>
        <ErrorPopup ref="error"/>
    </div>
</template>

<script>
    import axios from "axios";
    import ErrorPopup from "../ErrorPopup/ErrorPopup";
    import SuccessPopup from "../SuccessPopup/SuccessPopup";
    import PeriodDefinition from "../PeriodDefinition/PeriodDefinition";

    const config = require('electron').remote.getGlobal('config');

    export default {

        components: {
            PeriodDefinition,
            SuccessPopup,
            ErrorPopup
        },

        created() {
            this.retrieveCalendar();
        },

        data() {
            return {
                Calendar: {
                    Semesters: {
                        Winter: {
                            Start: '',
                            End: ''
                        },
                        Sprint: {
                            Start: '',
                            End: ''
                        }
                    },
                    Exams: {
                        September: {
                            Start: '',
                            End: ''
                        },
                        Winter: {
                            Start: '',
                            End: ''
                        },
                        Sprint: {
                            Start: '',
                            End: ''
                        }
                    },
                    Breaks: {
                        Winter: {
                            Start: '',
                            End: ''
                        },
                        Sprint: {
                            Start: '',
                            End: ''
                        },
                        Various: []
                    }
                },
                dateOption: null,
                date: ''
            }
        },

        methods: {
            send: function() {
                axios({
                    method: 'post',
                    url: config.api + '/api/calendarInfo',
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: this.Calendar
                }).then(() => {
                    this.$refs.success.showSuccess();
                }).catch(() => {
                    this.$refs.error.setError(this.$t('message.calFail'));
                    this.$refs.error.showAlert();
                })
            },

            retrieveCalendar: function() {
                axios({
                    method: 'get',
                    url: config.api + '/api/calendarInfo/full'
                }).then(result => {
                    this.Calendar = result.data;
                })
            },

            removeDate: function() {
                let index = this.Calendar.Breaks.Various.indexOf(this.dateOption);
                if (index > -1) {
                    this.Calendar.Breaks.Various.splice(index, 1);
                }
            },

            removeAll: function() {
                this.Calendar.Breaks.Various = [];
            },

            addDate: function(e) {
                if (this.date === '') {
                    this.$refs.error.setError(this.$t('message.adminInvalidDate'));
                    this.$refs.error.showAlert();
                    this.date = '';
                    e.preventDefault();
                    return;
                }
                this.Calendar.Breaks.Various.push(this.date);
                this.date = '';
                e.preventDefault();
            },
        }
    }
</script>

<style lang="scss">
@import "../scss/AcademicCalendar";
@import "~bootstrap/scss/bootstrap";
</style>