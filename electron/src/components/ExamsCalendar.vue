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
                                     :close-on-select="false" :options="sidOptions" :multiple="false"
                                     v-bind:select-label="this.$t('message.ecSelect')" v-bind:deselect-label="this.$t('message.ecDeselect')">
                            <span slot="noResult">Oops! No elements found. Consider changing the search query.</span>
                        </multiselect>
                    </div>
                    <div class="col col-lg-2">
<!--                        DEPARTMENT_ID_UNIQUE-->
                        <multiselect v-model="udepSelection" v-bind:placeholder="this.$t('message.ecSelectDepart')"
                                     :close-on-select="false" :options="depOptions" :multiple="false"
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
                        <b-button @click="addExam" class="float-right" variant="primary">{{$t('message.ecAdd')}}</b-button>
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
    </div>
</template>

<script>
    import Multiselect from 'vue-multiselect';

    export default {
        components: {
            Multiselect
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
                sidSelection: '',
                sidOptions: [1, 2, 3, 4, 5, 6, 7, 8, 9, 10],
                udepSelection: '',
                extrDepSelections: [],
                depOptions: ['ΠΛΗΡ', 'ΟΕ', 'ΟΔΕ'],
                submittedExams: [],
                fields: [
                    { key: 'name' },
                    { key: 'examiner' },
                    { key: 'startTime' },
                    { key: 'endTime' },
                    { key: 'date' },
                    { key: 'rooms' },
                    { key: 'semester' },
                    { key: 'department' },
                    { key: 'extraDepartments' }
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
                }
            }
        },

        methods: {
            addExam: function () {
                let newExam = Object.create(this.exam);
                newExam.name = this.name;
                newExam.examiner = this.examiner;
                newExam.startTime = this.startTime;
                newExam.endTime = this.endTime;
                newExam.date = this.date;
                newExam.rooms = this.roomSelections;
                newExam.semester = this.sidSelection;
                newExam.department = this.udepSelection;
                newExam.extraDepartments = this.extrDepSelections;
                this.clearForm();
                this.submittedExams.push(newExam);
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
            }
        }
    }
</script>

<style lang="scss">
@import "~vue-multiselect/dist/vue-multiselect.min.css";
@import "../scss/ExamsCalendar";
@import "~bootstrap/scss/bootstrap";
</style>