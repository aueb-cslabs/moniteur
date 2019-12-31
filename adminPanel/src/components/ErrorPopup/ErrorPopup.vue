<template>
    <div class="error">
        <b-alert
                :show="dismissCountDown"
                dismissible
                variant="danger"
                @dismissed="dismissCountDown=0"
                @dismiss-count-down="countDownChanged"
        >
            <p>Error! {{error}}</p>
            <b-progress
                    variant="dark"
                    :max="dismissSecs"
                    :value="dismissCountDown"
                    height="4px"
            ></b-progress>
        </b-alert>
    </div>
</template>

<script>
    export default {
        name: "ErrorPopup",

        data() {
            return {
                dismissSecs: 5,
                dismissCountDown: 0,
                error: null
            }
        },

        methods: {
            countDownChanged(dismissCountDown) {
                this.dismissCountDown = dismissCountDown
            },

            showAlert() {
                this.dismissCountDown = this.dismissSecs
            },

            setError(error) {
                this.error = error;
            }
        }
    }
</script>

<style lang="scss">
@import "~bootstrap/scss/bootstrap";
@import "ErrorPopup";
</style>