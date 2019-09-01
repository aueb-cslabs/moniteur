<template>
    <div class="roomAnnouncements">
        <p class="mt-4 common room-ann-text" v-if="roomAnnouncement != null">
            <i class="far fa-bullhorn"></i> {{roomAnnouncement}}
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
                    fetch("http://localhost:1323/api/announcement/a")
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
    @import "../css/RoomAnnouncement.scss";
</style>