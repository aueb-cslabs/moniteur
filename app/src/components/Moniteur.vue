<template>
    <div>
        <div class="parent-flex">
            <TitleBar/>
            <div class="flex-container">
                <Exam/>
                <Announcement/>
                <Schedule/>
                <RoomAnnouncement/>
                <Comment/>
                <footer class="text-center footer">
                    {{ $t("message.touchscreen")}}
                </footer>
            </div>
        </div>
    </div>
</template>

<script>
    import config from '../config/config.js';

    import TitleBar from './TitleBar/TitleBar.vue'
    import Schedule from "./Schedule/Schedule";
    import Exam from "./Exam/Exam";
    import Announcement from "./Announcement/Announcement";
    import RoomAnnouncement from "./RoomAnnouncement/RoomAnnouncement";
    import Comment from "./Comment/Comment";

    export default {

        created() {
            let r = document.documentElement;
            r.style.setProperty("--background", config.background_color);
        },

        components: {
            Comment,
            RoomAnnouncement,
            Announcement,
            TitleBar,
            Schedule,
            Exam
        }
    }
</script>

<style lang="scss">
    @import "~bootstrap/scss/bootstrap.scss";
    @import "~@fortawesome/fontawesome-free/css/all.css";
    @import "Moniteur";
</style>
