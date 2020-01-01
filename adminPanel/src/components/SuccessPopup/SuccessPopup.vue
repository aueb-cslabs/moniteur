<template>
    <div class="success">
        <b-alert
                :show="successCountDown"
                dismissible
                variant="success"
                @dismissed="successCountDown=0"
                @dismiss-count-down="successChanged"
        >
            <p>{{$t('message.calSuccess')}}</p>
            <b-progress
                    variant="dark"
                    :max="dismissSecs"
                    :value="successCountDown"
                    height="4px"
            ></b-progress>
        </b-alert>
    </div>
</template>

<script>
    export default {
        name: "SuccessPopup",

        data() {
            return {
                successCountDown: 0,
                dismissSecs: 5
            }
        },

        methods: {
            successChanged(successCountDown) {
                this.successCountDown = successCountDown;
            },

            showSuccess() {
                this.successCountDown = this.dismissSecs;
            },
        }
    }
</script>

<style lang="scss">
@import "SuccessPopup";
</style>