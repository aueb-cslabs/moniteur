<template>
    <div class="comments">
        <p class="mt-4 common comments-text" v-if="comment != null">
           {{comment}}
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
                    fetch("http://localhost:1323/api/comment")
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