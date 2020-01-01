<template>
    <form>
        <div class="form-group">
            <label for="start">{{$t("message.calSemesterStart")}}</label>
            <input type="date"
                   class="form-control"
                   id="start"
                   :value="start"
                   v-bind:placeholder="this.$t('message.adminAnnDateForm')"
                   @input="update('Start', $event.target.value)"
                   required>
        </div>
        <div class="form-group">
            <label for="end">{{$t("message.calSemesterEnd")}}</label>
            <input type="date"
                   class="form-control"
                   id="end"
                   :value="end"
                   v-bind:placeholder="this.$t('message.adminAnnDateForm')"
                   @input="update('End', $event.target.value)"
                   required>
        </div>
    </form>
</template>

<script>
    export default {
        name: "PeriodDefinition",

        props: {
            start: String,
            end: String,
            period: Object
        },

        methods: {
            update(key, value) {
                this.$emit('input', {...this.$props.period, [key]: value })
            },
        }
    }
</script>

<style scoped lang="scss">
@import "PeriodDefinition";
</style>