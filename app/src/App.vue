<template>
  <TitleBar/>
</template>

<script>
import TitleBar from './components/TitleBar.vue'

export default {
  name: 'app',
  components: {
    TitleBar
  }
}
</script>

<style lang="scss">
@import "~bootstrap/scss/bootstrap.scss";
</style>
