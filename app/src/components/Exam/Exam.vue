<template>
    <div class="exams text-center" v-if="isExam">
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
                isExam: null
            }
        },

        created() {
            this.registerExamListener();
        },

        methods: {
            registerExamListener: function () {
               this.$parent.$on('exam', exam => {
                   this.isExam = exam;
               });
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