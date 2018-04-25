<template>
  <div class="other-pic">
    <lost-connection v-show="lostConn" />

    <video id="otherPic" class="other-pic__video" ref="otherPic" autoplay playsinline />
  </div>
</template>

<script>
import LostConnection from '@/components/LostConnection'

export default {
  name: 'OtherPic',
  props: ['stream', 'lostConn'],
  components: {
    LostConnection
  },
  watch: {
    stream(val) { this.startStream(val) }
  },
  methods: {
    startStream(stream) {
      const video = this.$refs.otherPic

      video.srcObject = stream
      setTimeout(() => video.play().catch(this.logError), 1000)
    },
    logError(err) {
      console.error(`${err.name}: ${err.message}`)
    }
  }
}
</script>

<style lang="stylus">
  @import '../../styles'

  .other-pic
    position relative
    width 100%
    height 100vh
    z-index 1

    &__video
      width 100%
      height 100%
      object-fit cover
</style>
