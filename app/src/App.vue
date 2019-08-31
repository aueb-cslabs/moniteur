<template>
  <div class="moniteur">
    <TitleBar/>
    <Exam/>
    <Schedule/>
  </div>
</template>

<script>
import TitleBar from './components/TitleBar.vue'
import Schedule from "./components/Schedule";
import Exam from "./components/Exam";

export default {
  name: 'app',
  components: {
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
