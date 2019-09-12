<template>
    <div class="comments">
        <p class="mt-2 common comments-text fade-in" v-if="comment != null">
           {{comment['msg']}}
        </p>
    </div>
</template>

<script>
    import axios from "axios";

    export default {
        name: "Comment",

        data() {
            return {
                comment: null
            }
        },

        created() {
            this.getComment();
        },

        methods: {
            getComment: function () {
                setInterval(() => {
                    axios.get(this.$root.$data['api'] + ":27522/api/comment")
                        .then(res => this.comment = res.data)
                }, 5000)
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/Comment.scss";
</style>