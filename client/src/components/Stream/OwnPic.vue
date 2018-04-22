<template>
  <div 
    class="own-pic"
    :class="{
      'own-pic--streaming': streaming,
      'own-pic--resized': resized
    }"
    ref="picContainer"
    @dblclick="resizePic()"
  > 
    <div 
      class="own-pic__drag-area" 
      @mousedown="mouseDownMove()" 
    />
    <video id="ownPic" class="own-pic__video" ref="ownPic" autoplay playsinline />
  </div>
</template>

<script>
export default {
  name: 'OwnPic',
  props: ['stream', 'streaming'],
  data: () => ({
    picW: 323.6,
    picH: 242.7,
    resized: false
  }),
  watch: {
    stream(val) { this.startStream(val) }
  },
  mounted() {
    window.addEventListener('mouseup', this.mouseUp, false)
  },
  methods: {
    mouseDownMove() {
      window.addEventListener('mousemove', this.mouseMove, true)
    },

    mouseUp() {
      window.removeEventListener('mousemove', this.mouseMove, true)
    },

    mouseMove(e) {
      const container = this.$refs.picContainer

      container.style.top = `${e.clientY - (this.picH / 2)}px`
      container.style.left = `${e.clientX - (this.picW / 2)}px`
    },

    resizePic() {
      this.resized = !this.resized

      this.picW = this.resized ? this.picW * 1.5 : this.picW / 1.5
      this.picH = this.resized ? this.picH * 1.5 : this.picH / 1.5
    },

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
    width 100%
    height 100%
    top 0
    right 0
    z-index 2
    resize both

    &--streaming
      width 323.6px
      height 242.7px
      padding 12px
    
    &--resized
      width calc(323.6px * 1.5)
      height calc(242.7px * 1.5)
      padding 12px

    &__video
      position relative
      width 100%
      height 100%
      object-fit cover
      border-radius 4px
      box-shadow -4px 7px 8px 0px rgba(0, 0, 0, 0.25)
      transform scaleX(-1)
      z-index 1

    &__drag-area
      position absolute
      width 40%
      height 40%
      top 0
      left 0
      right 0
      bottom 0
      margin auto
      z-index 2
      cursor move
</style>
