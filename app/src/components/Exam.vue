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
                isExam: ''
            }
        },

        created() {
            this.registerExamListener();
        },

        methods: {
            registerExamListener: function () {
               this.$on('exam', exam => {
                   this.isExam = exam;
               });
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/Exam.scss";
    @import "../scss/common.scss";

</style>