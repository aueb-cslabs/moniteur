<template>
    <div class="exams text-center" >
        <p class="exam-text common fade-in" v-if="isExam">
            <i class="fas fa-exclamation-triangle"></i> Πραγματοποιείται εξέταση!
        </p>
    </div>
</template>

<script>
    import {EventBus} from "./EventBus";

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
               EventBus.$on('exam', exam => {
                   this.isExam = exam;
               });
            }
        }
    }
</script>

<style lang="scss">
    @import "../css/Exam.scss";
    @import "../css/common.scss";

</style>