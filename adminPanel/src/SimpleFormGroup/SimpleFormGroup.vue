<template>
    <div>
        <div class="form-group">
            <label for="input">{{this.$props.label}}</label>
            <input type="text"
                   class="form-control"
                   id="input"
                   :value="modelData"
                   :placeholder="this.$props.label"
                   @input="$emit('input', $event.target.value)"
                   required>
        </div>
    </div>
</template>

<script>
    export default {
        name: "SimpleFormGroup",

        props: {
            label: String,
            modelData: String
        }
    }
</script>

<style scoped lang="scss">
@import "SimpleFormGroup";
</style>