<template>
    <div class="exams text-center" v-if="isExam">
        <p>
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