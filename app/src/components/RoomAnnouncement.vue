<template>
    <div class="roomAnnouncements">
        <p class="mt-1 common room-ann-text" v-if="roomAnnouncement != null">
            <i class="far fa-bullhorn"></i> {{roomAnnouncement['msg']}}
        </p>
    </div>
</template>

<script>
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
                    axios.get(this.$root.$data['api']+"/api/announcement/" + this.$route.params.id)
                         .then(res => this.roomAnnouncement = res.data);
                }, 5000)
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/RoomAnnouncement.scss";
</style>