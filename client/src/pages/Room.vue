<template>
  <div>
    <video-stream class="stream" />
  </div>
</template>

<script>
import VideoStream from '@/components/VideoStream'
import Peer from 'simple-peer'
import Routes from '@/static/routes'

export default {
  name: 'Room',
  data: () => ({
    ws: null,
    peer: null
  }),
  components: {
    VideoStream
  },
  beforeMount() {
    this.ws = new WebSocket(Routes.socketPath(this.$router.history.current.params.id))
  },
  mounted() {
    this.$store.dispatch('checkRoom')

    const peer = new Peer({
      initiator: true,
      trickle: false
    })

    console.log(this.ws)

    peer.on('signal', e => {
      this.peer = e

      console.log(e)

      this.ws.send(JSON.stringify(e))
    })
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

    &__flag
      color red

  .stream
    position relative
    z-index 1
</style>
