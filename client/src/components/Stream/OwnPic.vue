<template>
  <div class="own-pic">
    <video id="ownPic" class="own-pic__video" ref="ownPic" autoplay playsinline />
  </div>
</template>

<script>
export default {
  name: 'OwnPic',
  props: ['stream'],
  watch: {
    stream(val) { this.startStream(val) }
  },
  methods: {
    startStream(stream) {
      const video = this.$refs.ownPic

      video.muted = true      
      video.srcObject = stream
      setTimeout(() => video.play().catch(this.logError), 1000)
    },
    logError(err) {
      console.error(`${err.name}: ${err.message}`)
    }
  }
}
</script>

<style lang="stylus" scoped>
  .own-pic
    position absolute
    width 200px
    height 150px
    top 0
    right 0
    padding 12px
    z-index 2

    &__video
      width 100%
      height 100%
      object-fit cover
      border-radius 4px
      box-shadow -4px 7px 8px 0px rgba(0, 0, 0, 0.25)
      transform scaleX(-1)
</style>
