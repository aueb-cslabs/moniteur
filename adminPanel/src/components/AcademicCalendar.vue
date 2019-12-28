<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.ecGeneral")}}</h2>
            <div class="pt-3">
                <div class="row">
                    <div class="col">
                        <h5>{{$t('message.calWinterS')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Semesters.Winter.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Semesters.Winter.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSprintS')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Semesters.Sprint.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Semesters.Sprint.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            <h2 class="pt-3">{{$t("message.calExams")}}</h2>
            <div class="pt-3">
                <div class="row">
                    <div class="col">
                        <h5>{{$t('message.calWin')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Exams.Winter.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Exams.Winter.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSpr')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Exams.Sprint.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Exams.Sprint.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSep')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Exams.September.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Exams.September.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                </div>
            </div>
            <h2 class="pt-3">{{$t("message.calBreaks")}}</h2>
            <div class="pt-3">
                <div class="row">
                    <div class="col">
                        <h5>{{$t('message.calWinterS')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Breaks.Winter.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Breaks.Winter.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSprintS')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Breaks.Sprint.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="datetime-local" class="form-control" v-model="Calendar.Breaks.Sprint.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                </div>
                <h5>{{$t('message.calVarious')}}</h5>
                <div>
                    <label>{{$t("message.calAddDate")}}</label>
                    <div>
                        <form>
                            <div class="row">
                                <div class="col-md-8 col-lg-10 form-group">
                                    <input type="datetime-local" class="form-control" v-model="date" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
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
                                <b-button @click="checkForm" variant="primary" class="ml-3 mt-3 float-right">{{this.$t("message.calSend")}}</b-button>
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <div class="error">
            <b-alert
                    :show="dismissCountDown"
                    dismissible
                    variant="danger"
                    @dismissed="dismissCountDown=0"
                    @dismiss-count-down="countDownChanged"
            >
                <p>Error! {{error}}</p>
                <b-progress
                        variant="dark"
                        :max="dismissSecs"
                        :value="dismissCountDown"
                        height="4px"
                ></b-progress>
            </b-alert>
        </div>
        <div class="success">
            <b-alert
                    :show="successCountDown"
                    dismissible
                    variant="success"
                    @dismissed="successCountDown=0"
                    @dismiss-count-down="successChanged"
            >
                <p>{{$t('message.calSuccess')}}</p>
                <b-progress
                        variant="dark"
                        :max="dismissSecs"
                        :value="successCountDown"
                        height="4px"
                ></b-progress>
            </b-alert>
        </div>
    </div>
</template>

<script>
    import axios from "axios";
    import functions from "../functions";
    import traverse from "traverse";
    import datetime from 'vuejs-datetimepicker';

    const config = require('electron').remote.getGlobal('config');

    export default {
        components: { datetime },

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
                dismissSecs: 5,
                dismissCountDown: 0,
                successCountDown: 0,
                error: null,
                date: ''
            }
        },

        methods: {
            checkForm: function () {

                let dates = traverse(this.Calendar).reduce(function (dates, x) {
                    if (this.isLeaf) {
                        if (Array.isArray(x)) {
                            dates = dates.concat(x);
                        }
                        else {
                            dates.push(x);
                        }
                    }
                    return dates;
                }, []);
                for (let i in dates) {
                    if (!functions.isGoodDate(dates[i])) {
                        this.error = this.$t('message.adminInvalidDate');
                        this.showAlert();
                        return;
                    }
                }
                this.error = null;
                this.send();
            },

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
                    this.showSuccess();
                }).catch(() => {
                    this.error = this.$t('message.calFail');
                    this.showAlert();
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
                if (!functions.isGoodDate(this.date)) {
                    this.error = this.$t('message.adminInvalidDate');
                    this.showAlert();
                    this.date = '';
                    datetime.methods.clearDate();
                    e.preventDefault();
                    return;
                }
                this.Calendar.Breaks.Various.push(this.date);
                this.date = '';
                e.preventDefault();
            },

            countDownChanged(dismissCountDown) {
                this.dismissCountDown = dismissCountDown
            },

            showAlert() {
                this.dismissCountDown = this.dismissSecs
            },

            successChanged(successCountDown) {
                this.successCountDown = successCountDown
            },

            showSuccess() {
                this.successCountDown = this.dismissSecs;
            },
        }
    }
</script>

<style lang="scss">
@import "../scss/AcademicCalendar";
@import "~bootstrap/scss/bootstrap";
</style>