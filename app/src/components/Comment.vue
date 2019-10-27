<template>
    <div class="comments" v-if="comment != null">
        <p class="mt-2 common comments-text fade-in">
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
                    axios.get(this.$root.$data['api'] + "/api/comment")
                        .then(res => this.comment = res.data)
                }, 5000)
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/Comment.scss";
</style>