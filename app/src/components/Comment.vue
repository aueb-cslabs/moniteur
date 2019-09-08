<template>
    <div class="comments">
        <p class="mt-2 common comments-text fade-in" v-if="comment != null">
           {{comment['msg']}}
        </p>
    </div>
</template>

<script>
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
                    fetch(this.$root.$data['api']+":27522/api/comment")
                        .then(response => response.json())
                        .then(comment => {
                            this.comment = comment;
                        })
                }, 5000)
            }
        }
    }
</script>

<style lang="scss">
    @import "../css/Comment.scss";
</style>