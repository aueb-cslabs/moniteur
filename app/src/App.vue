<template>
  <div class="moniteur">
    <TitleBar/>
    <Exam/>
    <Announcement/>
    <Schedule/>
    <RoomAnnouncement/>
    <Comment/>
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
@import "css/App.scss";
@import "fontawesome-pro-5.6.0-web/css/all.min.css";
</style>
