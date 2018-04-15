<template>
  <div class="video-stream">
    <video id="videoStream" class="video-stream__video" ref="videoStream" autoplay playsinline />
  </div>
</template>

<script>
export default {
  name: 'VideoSream',
  mounted() {
    this.startStream()
  },
  methods: {
    startStream() {
      const video = this.$refs.videoStream

      video.muted = true

      navigator.getMedia = (
        navigator.getUserMedia || 
        navigator.webkitGetUserMedia || 
        navigator.mozGetUserMedia || 
        navigator.msGetUserMedia
      )

      const handleSuccess = stream => {
        video.srcObject = stream
      
        setTimeout(() => video.play().catch(this.logError), 1000)
      }

      navigator.mediaDevices.getUserMedia({ video: true, audio: true })
        .then(handleSuccess)
        .catch(this.logError)
    },
    logError(err) {
      console.error(`${err.name}: ${err.message}`)
    }
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
