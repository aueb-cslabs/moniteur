<template>
    <div class="exams text-center" >
        <p class="exam-text" v-if="isExam">
            <i class="fas fa-feather"></i> Πραγματοποιείται εξέταση!
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

</style>