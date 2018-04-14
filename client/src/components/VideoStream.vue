<template>
  <div class="video-stream">
    <video id="videoStream" class="video-stream__video" ref="videoStream" autoplay playsinline />
  </div>
</template>

<script>
import WebRTC from '@/utils/webrtc'

export default {
  name: 'VideoSream',
  mounted() {
    WebRTC.start()
  }
}
</script>

<style lang="stylus" scoped>
  .video-stream
    width 100%
    height 100vh

    &__video
      width 100%
      height 100%
      object-fit cover
      transform scaleX(-1)
</style>
