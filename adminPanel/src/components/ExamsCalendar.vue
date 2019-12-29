<template>
    <div>
        <div class="pt-5 pb-5">
            <h2>{{$t("message.ecGeneral")}}</h2>
            <div class="=pt-3">
                <h5>{{$t('message.ecSubmit')}}</h5>
                <div class="row">
                    <div class="col col-lg-4">
<!--                        NAME-->
                        <b-form-input v-model="name" v-bind:placeholder="this.$t('message.ecLessonName')"></b-form-input>
                    </div>
                    <div class="col col-lg-2">
                        <b-form-input v-model="examiner" v-bind:placeholder="this.$t('message.ecExaminer')"></b-form-input>
                    </div>
                    <div class="col col-lg-3">
<!--                        ROOMS-->
                        <multiselect v-model="roomSelections" v-bind:placeholder="this.$t('message.ecSelectRooms')"
                                     :close-on-select="false" :options="roomOptions" :multiple="true"
                                     v-bind:select-label="this.$t('message.ecSelect')" v-bind:deselect-label="this.$t('message.ecDeselect')">
                            <span slot="noResult">Oops! No elements found. Consider changing the search query.</span>
                        </multiselect>
                    </div>
                    <div class="col col-lg-3">
<!--                        DATE-->
                        <b-form-input v-model="date" v-bind:placeholder="this.$t('message.ecDate')"></b-form-input>
                    </div>
                </div>
                <div class="row pt-3">
                    <div class="col col-lg-2">
<!--                        START_ΤΙΜΕ-->
                        <b-form-input v-model="startTime" v-bind:placeholder="this.$t('message.ecStart')"></b-form-input>
                    </div>
                    <div class="col col-lg-2">
<!--                        END_TIME-->
                        <b-form-input v-model="endTime" v-bind:placeholder="this.$t('message.ecEnd')"></b-form-input>
                    </div>
                    <div class="col col-lg-2">
<!--                        SEMESTER_ID-->
                        <multiselect v-model="sidSelection" v-bind:placeholder="this.$t('message.ecSelectSem')"
                                     :close-on-select="true" :options="sidOptions" :multiple="true" :max="1"
                                     v-bind:select-label="this.$t('message.ecSelect')" v-bind:deselect-label="this.$t('message.ecDeselect')">
                            <span slot="noResult">Oops! No elements found. Consider changing the search query.</span>
                        </multiselect>
                    </div>
                    <div class="col col-lg-2">
<!--                        DEPARTMENT_ID_UNIQUE-->
                        <multiselect v-model="udepSelection" v-bind:placeholder="this.$t('message.ecSelectDepart')"
                                     :close-on-select="true" :options="depOptions" :multiple="true" :max="1"
                                     v-bind:select-label="this.$t('message.ecSelect')" v-bind:deselect-label="this.$t('message.ecDeselect')">
                            <span slot="noResult">Oops! No elements found. Consider changing the search query.</span>
                        </multiselect>
                    </div>
                    <div class="col col-lg-4">
<!--                        DEPARTMENT_ID_EXTRA-->
                        <multiselect v-model="extrDepSelections" v-bind:placeholder="this.$t('message.ecSelectDepartExt')"
                                     :close-on-select="false" :options="depOptions" :multiple="true"
                                     v-bind:select-label="this.$t('message.ecSelect')" v-bind:deselect-label="this.$t('message.ecDeselect')">
                            <span slot="noResult">Oops! No elements found. Consider changing the search query.</span>
                        </multiselect>
                    </div>
                </div>
                <div class="row pt-3">
                    <div class="col">
<!--                        ADD_BUTTON-->
                        <b-button @click="checkEntry" class="float-right" variant="primary">{{$t('message.ecAdd')}}</b-button>
                    </div>
                </div>
            </div>
            <h2 class="pt-3">{{$t('message.ecRegExams')}}</h2>
            <div class="pt-3">
                <b-table ref="examsTable" striped hover :fields="fields" :items="submittedExams"></b-table>
            </div>
            <div class="pt-3">
                <!--                SEARCH BAR AND LIST FOR DELETE?-->
                <!--                UPDATE_BUTTON-->
                <!--                        PREVIEW_BUTTON-->
                <b-button @click="doSmth2" class="float-left" variant="info">{{$t('message.ecPreview')}}</b-button>
                <!--                        SEND_BUTTON-->
                <b-button @click="doSmth3" class="ml-3 float-right" variant="primary">{{$t('message.ecCreate')}}</b-button>
            </div>
        </div>
        <ErrorPopup ref="error"/>
    </div>
</template>

<script>
    import Multiselect from 'vue-multiselect';
    import functions from "../functions";
    import datetime from 'vuejs-datetimepicker';
    import ErrorPopup from "../ErrorPopup/ErrorPopup";

    export default {
        components: {
            ErrorPopup,
            Multiselect,
            datetime
        },

        data() {
            return {
                name: '',
                examiner: '',
                startTime: '',
                endTime: '',
                date: '',
                roomOptions: ['a', 'b'],
                roomSelections: [],
                sidSelection: [],
                sidOptions: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
                udepSelection: [],
                extrDepSelections: [],
                depOptions: ['ΠΛΗΡ', 'ΟΕ', 'ΟΔΕ'],
                submittedExams: [],
                fields: [
                    { key: 'name', label: this.$t('message.ecLessonName') },
                    { key: 'examiner', label: this.$t('message.ecExaminer') },
                    { key: 'startTime', label: this.$t('message.ecStart') },
                    { key: 'endTime', label: this.$t('message.ecEnd') },
                    { key: 'date', label: this.$t('message.ecDate') },
                    { key: 'rooms', label: this.$t('message.ecSelectRooms'), formatter: "formatTable" },
                    { key: 'semester', label: this.$t('message.ecSelectSem'), formatter: "formatTable" },
                    { key: 'department', label: this.$t('message.ecSelectDepart'), formatter: "formatTable" },
                    { key: 'extraDepartments', label: this.$t('message.ecSelectDepartExt'), formatter: "formatTable" }
                ],
                exam: {
                    name: '',
                    examiner: '',
                    startTime: '',
                    endTime: '',
                    date: '',
                    rooms: [],
                    semester: 0,
                    department: '',
                    extraDepartments: []
                },
            }
        },

        methods: {
            addExam: function () {
                let newExam = {
                    name : this.name,
                    examiner : this.examiner,
                    startTime : this.startTime,
                    endTime : this.endTime,
                    date : this.date,
                    rooms : this.roomSelections,
                    semester : this.formatTable(this.sidSelection),
                    department : this.formatTable(this.udepSelection) ,
                    extraDepartments : this.extrDepSelections,
                };
                this.submittedExams.push(newExam);
                this.clearForm();
                this.$refs.examsTable.refresh();
            },

            clearForm: function () {
                this.name = '';
                this.examiner = '';
                this.startTime = '';
                this.endTime = '';
                this.date = '';
                this.roomSelections = [];
                this.sidSelection = '';
                this.udepSelection = '';
                this.extrDepSelections = [];
            },

            formatTable: function (table) {
                return table.toString();
            },

            checkEntry: function (e) {
                if (this.name === '') {
                    this.$refs.error.setError(this.$t('message.ecFormError'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.examiner === '') {
                    this.$refs.error.setError(this.$t('message.ecFormError'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (!functions.isValidTime(this.startTime)) {
                    this.$refs.error.setError(this.$t('message.ecFormError'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (!functions.isValidTime(this.endTime)) {
                    this.$refs.error.setError(this.$t('message.ecFormError'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (!functions.isGoodDate(this.date)) {
                    this.$refs.error.setError(this.$t('message.ecFormError'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.roomSelections.length === 0) {
                    this.$refs.error.setError(this.$t('message.ecFormError'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.sidSelection.length === 0) {
                    this.$refs.error.setError(this.$t('message.ecFormError'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                if (this.udepSelection.length === 0) {
                    this.$refs.error.setError(this.$t('message.ecFormError'));
                    this.$refs.error.showAlert();
                    e.preventDefault();
                    return;
                }
                this.addExam();
                e.preventDefault();
            }
        }
    }
</script>

<style lang="scss">
@import "~vue-multiselect/dist/vue-multiselect.min.css";
@import "../scss/ExamsCalendar";
@import "~bootstrap/scss/bootstrap";
</style>