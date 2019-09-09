<template>
    <div class="roomAnnouncements">
        <p class="mt-1 common room-ann-text" v-if="roomAnnouncement != null">
            <i class="far fa-bullhorn"></i> {{roomAnnouncement['msg']}}
        </p>
    </div>
</template>

<script>
    export default {
        name: "RoomAnnouncement",

        data() {
            return {
                roomAnnouncement: null
            }
        },

        created() {
            this.getAnnouncement();
        },

        methods: {
            getAnnouncement: function () {
                setInterval(() => {
                    fetch(this.$root.$data['api'] + this.$root.$data['port'] + "/api/announcement/" + this.$root.$data['room'])
                        .then(response => response.json())
                        .then(roomAnn => {
                            this.roomAnnouncement = roomAnn;
                        })
                }, 5000)
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/RoomAnnouncement.scss";
</style>