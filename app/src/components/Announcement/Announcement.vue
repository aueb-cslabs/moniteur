<template>
    <div class="announcements" v-if="announcement != null">
        <p class="mt-4 common ann-text fade-in">
            <i class="fas fa-scroll"></i> {{announcement['msg']}}
        </p>
    </div>
    
</template>

<script>
    import axios from 'axios';
    const config = require('electron').remote.getGlobal('config');

    export default {
        name: "Announcement",

        data() {
            return {
                announcement: null
            }
        },

        created() {
            this.getAnnouncement();
        },

        methods: {
            getAnnouncement: function () {
                setInterval(() => {
                    axios.get(config.api + "/api/announcement")
                        .then(response => this.announcement = response.data)
                }, 5000);
            }
        }
    }
</script>

<style lang="scss">
    @import "Announcement";
    @import "../common";
</style>