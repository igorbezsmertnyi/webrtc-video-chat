<template>
  <div>
    <code class="peer-conn" v-html="room.peer" />
    <video-stream class="stream" />
  </div>
</template>

<script>
import VideoStream from '@/components/VideoStream'

export default {
  name: 'Room',
  components: {
    VideoStream
  },
  mounted() {
    this.$store.dispatch('checkRoom')
  },
  computed: {
    room() {
      return this.$store.getters.room
    }
  }
}
</script>

<style lang="stylus" scoped>
  .peer-conn
    position absolute
    max-width 800px
    padding 10px
    background rgba(0,0,0, .75)
    color green
    z-index 2

  .stream
    position relative
    z-index 1
</style>
