<template>
    <div>
        <div class="parent-flex">
            <TitleBar/>
            <div class="flex-container">
                <Exam/>
                <Announcement/>
                <Schedule/>
                <div class="bottom align-items-center">
                    <RoomAnnouncement/>
                    <Comment/>
                </div>
                <footer class="mt-5 text-center">
                    {{ $t("message.touchscreen")}}
                </footer>
            </div>
        </div>
    </div>
</template>

<script>
import TitleBar from './TitleBar.vue'
import Schedule from "./Schedule";
import Exam from "./Exam";
import Announcement from "./Announcement";
import RoomAnnouncement from "./RoomAnnouncement";
import Comment from "./Comment";

export default {

    created() {
        let r = document.documentElement;
        r.style.setProperty("--background", this.$root.$data['background_color']);
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
    @import "../scss/Moniteur.scss";
</style>
