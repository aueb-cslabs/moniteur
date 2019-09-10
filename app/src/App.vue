<template>
  <div class="moniteur">
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
</template>

<script>
import TitleBar from './components/TitleBar.vue'
import Schedule from "./components/Schedule";
import Exam from "./components/Exam";
import Announcement from "./components/Announcement";
import RoomAnnouncement from "./components/RoomAnnouncement";
import Comment from "./components/Comment";

export default {
  name: 'app',

  created() {
    let r = document.documentElement;

    r.style.setProperty("--background", this.$root.$data['BACKGROUND']);
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
@import "scss/App.scss";
@import "fontawesome-pro-5.6.0-web/css/all.min.css";
</style>
