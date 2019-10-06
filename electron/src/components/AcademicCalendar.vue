<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.calReg")}}</h2>
            <div class="pt-3">
                <div class="row">
                    <div class="col">
                        <h5>{{$t('message.calWinterS')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Semesters.Winter.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Semesters.Winter.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSprintS')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Semesters.Sprint.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Semesters.Sprint.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
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
                                <input type="text" class="form-control" v-model="Calendar.Exams.Winter.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Exams.Winter.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSpr')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Exams.Sprint.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Exams.Sprint.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSep')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Exams.September.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Exams.September.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
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
                                <input type="text" class="form-control" v-model="Calendar.Breaks.Winter.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Breaks.Winter.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                        </form>
                    </div>
                    <div class="col">
                        <h5>{{$t('message.calSprintS')}}</h5>
                        <form @submit="checkForm">
                            <div class="form-group">
                                <label>{{$t("message.calSemesterStart")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Breaks.Sprint.Start" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
                            </div>
                            <div class="form-group">
                                <label>{{$t("message.calSemesterEnd")}}</label>
                                <input type="text" class="form-control" v-model="Calendar.Breaks.Sprint.End" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
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
                                <div class="col-md-10 form-group">
                                    <input type="text" class="form-control" v-model="date" v-bind:placeholder="this.$t('message.adminAnnDateForm')">
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
                                <button @click="checkForm" type="submit" class="btn btn-primary ml-3 mt-3 float-right">{{this.$t("message.calSend")}}</button>
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
    </div>
</template>

<script>
    import axios from "axios";

    export default {
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
                error: null,
                date: ''
            }
        },

        methods: {
            checkForm: function () {

                //TODO CHANGE THIS, IT IS STUPID!!!

                for (let key in this.Calendar.Semesters.Winter) {
                    if (this.Calendar.Semesters.Winter.hasOwnProperty(key)) {
                        if(!this.isGoodDate(this.Calendar.Semesters.Winter[key])){
                            this.errorF();
                            return;
                        }
                    }
                }
                for (let key in this.Calendar.Semesters.Sprint) {
                    if (this.Calendar.Semesters.Sprint.hasOwnProperty(key)) {
                        if(!this.isGoodDate(this.Calendar.Semesters.Sprint[key])){
                            this.errorF();
                            return;
                        }
                    }
                }
                for (let key in this.Calendar.Exams.Winter) {
                    if (this.Calendar.Exams.Winter.hasOwnProperty(key)) {
                        if(!this.isGoodDate(this.Calendar.Exams.Winter[key])){
                            this.errorF();
                            return;
                        }
                    }
                }
                for (let key in this.Calendar.Exams.Sprint) {
                    if (this.Calendar.Exams.Sprint.hasOwnProperty(key)) {
                        if(!this.isGoodDate(this.Calendar.Exams.Sprint[key])){
                            this.errorF();
                            return;
                        }
                    }
                }
                for (let key in this.Calendar.Exams.September) {
                    if (this.Calendar.Exams.September.hasOwnProperty(key)) {
                        if(!this.isGoodDate(this.Calendar.Exams.September[key])){
                            this.errorF();
                            return;
                        }
                    }
                }
                for (let key in this.Calendar.Breaks.Winter) {
                    if (this.Calendar.Breaks.Winter.hasOwnProperty(key)) {
                        if(!this.isGoodDate(this.Calendar.Breaks.Winter[key])){
                            this.errorF();
                            return;
                        }
                    }
                }
                for (let key in this.Calendar.Breaks.Sprint) {
                    if (this.Calendar.Breaks.Sprint.hasOwnProperty(key)) {
                        if(!this.isGoodDate(this.Calendar.Breaks.Sprint[key])){
                            this.errorF();
                            return;
                        }
                    }
                }
                this.Calendar.Breaks.Various.forEach(value => {
                    if (!this.isGoodDate(value)) {
                        this.errorF();
                        return;
                    }
                });
                this.error = null;
                this.send();
            },

            errorF: function() {
                this.error = this.$t('message.adminInvalidDate');
                this.showAlert();
            },

            send: function() {
                axios({
                    method: 'post',
                    url: this.$root.$data['api'] + '/api/calendarInfo',
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: this.Calendar
                }).then(() => {
                    //TODO Show success result
                })
            },

            isGoodDate: function(dt){
                let reGoodDate = /([0-3]?\d\/{1})([01]?\d\/{1})([12]{1}\d{3}\/?)/g;
                return reGoodDate.test(dt);
            },

            retrieveCalendar: function() {
                axios({
                    method: 'get',
                    url: this.$root.$data['api'] + '/api/calendarInfo/full'
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
                if (!this.isGoodDate(this.date)) {
                    this.error = this.$t('message.adminInvalidDate');
                    this.showAlert();
                    this.date = '';
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
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/AcademicCalendar";
@import "~bootstrap/scss/bootstrap";
</style>