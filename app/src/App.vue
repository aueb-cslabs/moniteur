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
@import "scss/App.scss";
</style>
