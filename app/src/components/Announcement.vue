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
        },

        methods: {
            getAnnouncement: function () {
                setInterval(() => {
                    fetch(this.$root.$data['api'] + this.$root.$data['port'] + "/api/announcement")
                        .then(response => response.json())
                        .then(ann => {
                            this.announcement = ann;
                        })
                }, 5000);
            }
        }
    }
</script>

<style lang="scss">
    @import "../scss/Announcement.scss";
    @import "../scss/common.scss";
</style>