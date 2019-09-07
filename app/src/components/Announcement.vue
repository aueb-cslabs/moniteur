<template>
    <div class="announcements">
        <p class="mt-4 common ann-text fade-in" v-if="announcement != null">
            <i class="fas fa-scroll-old"></i> {{announcement['msg']}}
        </p>
    </div>
    
</template>

<script>
    export default {
        name: "Announcement",

        data() {
            return {
                announcement: null
            }
        },

        created() {
            this.getAnnouncement();
            this.removeAnnouncement();
        },

        methods: {
            getAnnouncement: function () {
                setInterval(() => {
                    fetch(this.$root.$data['api']+":27522/api/announcement")
                        .then(response => response.json())
                        .then(ann => {
                            this.announcement = ann;
                        })
                }, 5000);
            },

            removeAnnouncement: function () {
                setInterval(() => {
                    if (this.announcement != null) {
                        let timestamp = Math.round(+new Date()/1000);
                        if (timestamp >= this.announcement['end']) {
                            fetch(this.$root.$data['api']+":27522/api/announcement", {
                                method: 'delete'
                            })
                        }
                    }
                }, 3600000);
            }
        }
    }
</script>

<style lang="scss">
    @import "../css/Announcement.scss";
    @import "../css/common.scss";
</style>