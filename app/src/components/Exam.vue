<template>
    <div class="exams text-center" >
        <p class="exam-text common fade-in" v-if="isExam">
            <i class="fas fa-exclamation-triangle"></i> Πραγματοποιείται εξέταση!
        </p>
    </div>
</template>

<script>
    export default {
        name: "Exam",

        data() {
            return {
                isExam: true
            }
        }
    }
</script>

<style lang="scss">
    @import "../css/Exam.scss";
    @import "../css/common.scss";

</style>