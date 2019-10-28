<template>
    <div class="roomAnnouncements" v-if="roomAnnouncement != null">
        <p class="mt-1 common room-ann-text">
            <i class="fas fa-bullhorn"></i> {{roomAnnouncement['msg']}}
        </p>
    </div>
</template>

<script>
    import config from '../../config/config.js'
    import axios from 'axios';

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
                    axios.get(config.api + "/api/announcement/" + this.$route.params.id)
                         .then(res => this.roomAnnouncement = res.data);
                }, 5000)
            }
        }
    }
</script>

<style lang="scss">
    @import "RoomAnnouncement";
</style>