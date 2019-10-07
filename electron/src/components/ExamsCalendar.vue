<template>
    <div>
        <div class="pt-5">
            <h2>{{$t("message.ecGeneral")}}</h2>
            <div class="=pt-3">
                <h5>{{$t('message.ecSubmit')}}</h5>
                <div class="row">
                    <div class="col col-lg-3">
<!--                        NAME-->
                        <b-form-input v-model="name" v-bind:placeholder="this.$t('message.ecLessonName')"></b-form-input>
                    </div>
                    <div class="col col-lg-2">
<!--                        ROOMS-->
                        <multiselect class="form-control" v-model="roomSelections" v-bind:placeholder="this.$t('message.ecSelectRooms')"
                                     :close-on-select="true" :options="roomOptions" :multiple="true">

                        </multiselect>
                    </div>
                    <div class="col col-lg-2">
<!--                        DATE-->
                        <b-form-input v-model="date" v-bind:placeholder="this.$t('message.ecDate')"></b-form-input>
                    </div>
                    <div class="col col-lg-2">
<!--                        START_ΤΙΜΕ-->
                        <b-form-input v-model="startTime" v-bind:placeholder="this.$t('message.ecStart')"></b-form-input>
                    </div>
                    <div class="col col-lg-2">
<!--                        END_TIME-->
                        <b-form-input v-model="endTime" v-bind:placeholder="this.$t('message.ecEnd')"></b-form-input>
                    </div>
                    <div class="col">
<!--                        SEMESTER_ID-->
                    </div>
                    <div class="col">
<!--                        DEPARTMENT_ID-->
                    </div>
                    <div class="col">
<!--                        ADD_BUTTON-->
<!--                        PREVIEW_BUTTON-->
<!--                        SEND_BUTTON-->
                    </div>
                </div>
            </div>
            <div class="pt-3">
<!--                SEARCH BAR AND LIST FOR DELETE?-->
<!--                UPDATE_BUTTON-->
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
                startTime: '',
                endTime: '',
                date: '',
                roomOptions: ['a', 'b'],
                roomSelections: []
            }
        }
    }
</script>

<style lang="scss">
@import "../scss/ExamsCalendar";
@import "~bootstrap/scss/bootstrap";
</style>