<template>
  <div class="dock-btn screen">
    <button type="button" @click="makeScreen()">
      <svg-icon src="picture" />
      <canvas ref="screenShow" class="screen__canvas" />
    </button>
  </div>
</template>

<script>
import SvgIcon from '@/components/SvgIcon'

export default {
  name: 'ScreenShot',
  components: {
    SvgIcon
  },
  methods: {
    makeScreen() {
      const canvas = this.$refs.screenShow

      canvas.width = window.innerWidth
      canvas.height = window.innerWidth / 1.4

      const video = document.getElementById('otherPic')
      const context = canvas.getContext('2d')

      context.drawImage(video, 0, 0, canvas.width, canvas.height)

      this.$store.dispatch('setImage', canvas.toDataURL())
    }
  }
}
</script>

<style lang="stylus">
  @import '../../styles'

  .screen
    background rgba(255,255,255,0.8)

    &__canvas
      position absolute
      width 100vw
      height 100vh
      top 0
      left 0
      opacity 0
      z-index -1
</style>
