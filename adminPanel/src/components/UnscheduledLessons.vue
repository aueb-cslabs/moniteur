<template>
    <div>
        <div class="pt-5 pb-5">
            <h2>{{$t("message.adminBarUL")}}</h2>
            <div class="=pt-3">
                <h5>{{$t('message.ecSubmit')}}</h5>
                <div class="row">
                    <div class="col col-lg-5">
                        <!--                        NAME-->
                        <b-form-input v-model="title" v-bind:placeholder="this.$t('message.ecLessonName')"></b-form-input>
                    </div>
                    <div class="col col-lg-4">
                        <!--                        NAME-->
                        <b-form-input v-model="host" v-bind:placeholder="this.$t('message.ecProf')"></b-form-input>
                    </div>
                    <div class="col col-lg-3">
                        <!--                        ROOMS-->
                        <multiselect v-model="room" v-bind:placeholder="this.$t('message.ecSelectRooms')"
                                     :close-on-select="false" :options="roomOptions" :multiple="false"
                                     v-bind:select-label="this.$t('message.ecSelect')" v-bind:deselect-label="this.$t('message.ecDeselect')">
                            <span slot="noResult">Oops! No elements found. Consider changing the search query.</span>
                        </multiselect>
                    </div>
                </div>
                <div class="row pt-3">
                    <div class="col col-lg-6">
                        <!--                        START_ΤΙΜΕ-->
                        <b-form-input type="datetime-local" v-model="start" v-bind:placeholder="this.$t('message.ecStart')"></b-form-input>
                    </div>
                    <div class="col col-lg-6">
                        <!--                        END_TIME-->
                        <b-form-input type="datetime-local" v-model="end" v-bind:placeholder="this.$t('message.ecEnd')"></b-form-input>
                    </div>
                </div>
                <div class="row pt-3">
                    <div class="col">
                        <!--                        ADD_BUTTON-->
                        <b-button @click="deleteLesson" class="float-left" variant="danger">{{$t('message.ecRemove')}}</b-button>
                        <b-button @click="checkEntry" class="float-right" variant="primary">{{$t('message.ecAdd')}}</b-button>
                    </div>
                </div>
            </div>
            <h2 class="pt-3">{{$t('message.ecRegExams')}}</h2>
            <div class="pt-3">
                <b-table ref="unscheduledTable" striped
                         hover
                         :fields="fields"
                         :items="unscheduledLessons"
                         selectable
                         @row-selected="getRow"
                         select-mode="single"
                ></b-table>
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
    import Multiselect from 'vue-multiselect';
    import axios from 'axios';

    const config = require('electron').remote.getGlobal('config');

    export default {
        components: {
            Multiselect,
        },
        
        created() {
            this.loadRooms();
            this.fetchUnscheduled();
        },

        data() {
            return {
                title: '',
                host: '',
                start: '',
                end: '',
                date: '',
                roomOptions: [],
                room: '',
                unscheduledLessons: [],
                selectedRow: null,
                fields: [
                    { key: 'title', label: this.$t('message.ecLessonName') },
                    { key: 'host', label: this.$t('message.ecProf') },
                    { key: 'start', label: this.$t('message.ecStart'), formatter: "transformTimestamp" },
                    { key: 'end', label: this.$t('message.ecEnd'), formatter: "transformTimestamp" },
                    { key: 'room', label: this.$t('message.ecSelectRooms')},
                ],
                exam: {
                    title: '',
                    examiner: '',
                    start: '',
                    end: '',
                    date: '',
                    room: '',
                    semester: 0,
                    department: '',
                    extraDepartments: []
                },
                dismissSecs: 5,
                dismissCountDown: 0,
                error: null
            }
        },

        methods: {
            addExam: function () {
                let newExam = {
                    title : this.title,
                    host : this.host,
                    start : this.start,
                    end : this.end,
                    rooms : this.room,
                };
                this.unscheduledLessons.push(newExam);
                axios({
                    method: 'post',
                    url: config.api + "/api/override",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    },
                    data: {
                        room: this.room,
                        start: new Date(this.start)/1000,
                        end: new Date(this.end)/1000,
                        title: this.title,
                        host: this.host
                    }
                }).then(() => {
                    this.clearForm();
                    this.fetchUnscheduled();
                });
            },

            clearForm: function () {
                this.title = '';
                this.host = '';
                this.start = '';
                this.end = '';
                this.room = '';
            },

            checkEntry: function (e) {
                if (this.title === '') {
                    this.error = this.$t('message.ecFormError');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.examiner === '') {
                    this.host = this.$t('message.ecFormError');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.start === '') {
                    this.error = this.$t('message.ecFormError');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.end === '') {
                    this.error = this.$t('message.ecFormError');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.room === '') {
                    this.error = this.$t('message.ecFormError');
                    this.showAlert();
                    e.preventDefault();
                    return;
                }
                this.addExam();
                e.preventDefault();
            },

            countDownChanged(dismissCountDown) {
                this.dismissCountDown = dismissCountDown
            },

            showAlert() {
                this.dismissCountDown = this.dismissSecs
            },
            
            loadRooms: function () {
                axios({
                    method: 'get',
                    url: config.api + '/api/rooms'
                }).then(response => {
                    this.roomOptions = response.data;
                });
            },

            fetchUnscheduled: function () {
                axios({
                    method: 'get',
                    url: config.api + "/api/override",
                    headers: {
                        Username: this.$parent.$data['authToken'].username,
                        Authorization: this.$parent.$data['authToken'].token
                    }
                }).then((response) => {
                    this.unscheduledLessons = response.data;
                    this.$refs.unscheduledTable.refresh();
                });
            },

            transformTimestamp: function (table) {
                return new Date(table*1000).toLocaleDateString() + " " + new Date(table*1000).toLocaleTimeString();
            },

            getRow: function (row) {
                this.selectedRow = row;
            },

            deleteLesson: function () {
                if (this.selectedRow[0].db_info !== null) {
                    axios({
                        method: 'delete',
                        url: config.api + "/api/override",
                        headers: {
                            Username: this.$parent.$data['authToken'].username,
                            Authorization: this.$parent.$data['authToken'].token
                        },
                        data: {
                            db_info: this.selectedRow[0].db_info,
                            room: this.selectedRow[0].room,
                            start: this.selectedRow[0].start,
                            end: this.selectedRow[0].end,
                            title: this.selectedRow[0].title,
                            host: this.selectedRow[0].host,
                        }
                    }).then(() => {
                        this.fetchUnscheduled();
                    });
                }
                else {
                    this.error = this.$t('message.ecNoSelected');
                    this.showAlert();
                }
            }
        }
    }
</script>

<style lang="scss">
@import "~vue-multiselect/dist/vue-multiselect.min.css";
@import "../scss/ExamsCalendar";
@import "../scss/AcademicCalendar";
</style>