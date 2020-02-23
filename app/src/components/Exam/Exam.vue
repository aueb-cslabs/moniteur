<template>
    <div class="exams text-center" v-if="isExam || overrideExam">
        <p class="exam-text common fade-in" >
            <i class="fas fa-exclamation-triangle"></i> {{ $t("message.examMsg") }}
        </p>
    </div>
</template>

<script>
    export default {
        name: "Exam",

        data() {
            return {
                isExam: null,
                overrideExam: null
            }
        },

        created() {
            this.registerExamListener();
            this.registerOverrideListener();
        },

        methods: {
            registerExamListener: function () {
               this.$parent.$on('exam', exam => {
                   this.isExam = exam;
               });
            },

            registerOverrideListener: function () {
                this.$parent.$on('overrideExam', exam => this.overrideExam = exam)
            }
        }
    }
</script>

<style lang="scss">
    @import "~bootstrap/scss/bootstrap";
    @import "Exam";
    @import "Exam_phone";
    @import "Exam_tablet";
    @import "../common";
</style>