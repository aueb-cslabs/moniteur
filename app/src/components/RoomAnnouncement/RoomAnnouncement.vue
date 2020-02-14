<template>
    <div class="roomAnnouncements" v-if="roomAnnouncement != null">
        <p class="mt-1 common room-ann-text">
            <i class="fas fa-bullhorn"></i> {{roomAnnouncement['msg']}}
        </p>
    </div>
</template>

<script>
    import axios from 'axios';
    import config from "../../config";

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
                    axios.get(config.api + "/api/announcement/" + config.room)
                         .then(res => this.roomAnnouncement = res.data);
                }, 5000)
            }
        }
    }
</script>

<style lang="scss">
    @import "~bootstrap/scss/bootstrap";
    @import "RoomAnnouncement";
    @import "RoomAnnouncement_phone";
    @import "RoomAnnouncement_tablet";
</style>