<template>
    <div class="announcements">
        <p class="mt-4 common ann-text fade-in" v-if="announcement != null">
            <i class="fas fa-scroll"></i> {{announcement['msg']}}
        </p>
    </div>
    
</template>

<script>
    import axios from "axios"
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
                    axios.get(this.$root.$data['api'] + "/api/announcement")
                        .then(response => this.announcement = response.data)
                }, 5000);
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/Announcement.scss";
    @import "../scss/common.scss";
</style>