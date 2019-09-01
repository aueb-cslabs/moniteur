<template>
  <div class="moniteur">
    <TitleBar/>
    <Exam/>
    <Announcement/>
    <Schedule/>
    <RoomAnnouncement/>
  </div>
</template>

<script>
import TitleBar from './components/TitleBar.vue'
import Schedule from "./components/Schedule";
import Exam from "./components/Exam";
import Announcement from "./components/Announcement";
import RoomAnnouncement from "./components/RoomAnnouncement";

export default {
  name: 'app',
  components: {
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
