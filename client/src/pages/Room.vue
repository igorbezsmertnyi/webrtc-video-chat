<template>
  <div>
    <own-pic :stream="ownStream" :streaming="streaming" />
    <other-pic :stream="otherStream" :lostConn="lostConn" />
    <dock />
    <screen-img />
  </div>
</template>

<script>
import OwnPic from '@/components/Stream/OwnPic'
import OtherPic from '@/components/Stream/OtherPic'
import Dock from '@/components/Dock'
import ScreenImg from '@/components/ScreenImg'
import Routes from '@/static/routes'

export default {
  name: 'Room',
  data: () => ({
    id: null,
    ws: null,
    ownStream: null,
    otherStream: null,
    lostConn: false,
    streaming: false
  }),
  components: {
    OwnPic,
    OtherPic,
    Dock,
    ScreenImg
  },
  beforeMount() {
    this.id = Math.random().toString(36).substr(2, 9)
    this.ws = new WebSocket(Routes.socketPath(this.$router.history.current.params.id))
  },
  mounted() {
    this.$store.dispatch('checkRoom')
    this.startStream()
  },
  methods: {
    newCommand(data) {
      if (data.id === this.id || !this.room.created) return

      console.info('command NEW')
      this.connObject()
    },

    connCommand(data) {
      if (data.id === this.id) return

      console.info('command CONN')
      this.currentPeer.signal(data.peer)
    },

    newObject() {
      const obj = JSON.stringify({
        id: this.id,
        command: 'NEW'
      })
      
      this.ws.send(obj)
    },

    connObject() {
      const obj = JSON.stringify({
        id: this.id,
        command: 'CONN',
        peer: this.currentConn
      })

      this.ws.send(obj)
    },

    startStream() {
      navigator.getMedia = (
        navigator.getUserMedia || 
        navigator.webkitGetUserMedia || 
        navigator.mozGetUserMedia || 
        navigator.msGetUserMedia
      )

      const handleSuccess = stream => {
        this.ownStream = stream
        this.$store.dispatch('setPeer', { stream: stream, init: this.room.created })

        this.listenWs()
        this.peerConnected()
      }

      navigator.mediaDevices.getUserMedia({ video: true, audio: true })
        .then(handleSuccess)
        .catch(this.logError)
    },

    listenWs() {
      if (!this.room.created) this.newObject()

      this.currentPeer.on('signal', e => {
        this.$store.dispatch('setConn', e)
        console.info(`SIGNAL ${e.type}`)

        if (!this.room.created) this.connObject()
      })

      this.currentPeer.on('close', () => {
        if (!this.currentPeer.connected) {
          this.lostConn = true

          console.log('peer conncetion cloused')

          this.$store.dispatch('destroyPeer')
          this.$store.dispatch('setPeer', { stream: this.ownStream, init: this.room.created })

          this.listenWs()
          this.peerConnected()
        }
      })

      this.ws.onmessage = e => {
        const wsData = JSON.parse(e.data)

        switch (wsData.command) {
          case 'NEW':
            this.newCommand(wsData)
            break
          case 'CONN':
            this.connCommand(wsData)
            break 
        }
      }
    },

    peerConnected() {
      this.currentPeer.on('connect', () => console.info('peer conncection created'))
      this.currentPeer.on('stream', stream => { 
        this.streaming = true
        this.lostConn = false
        this.otherStream = stream 
      })
    },

    logError(err) {
      console.error(`${err.name}: ${err.message}`)
    }
  },
  computed: {
    room() {
      return this.$store.getters.room
    },

    currentPeer() {
      return this.$store.getters.currentPeer
    },

    currentConn() {
      return this.$store.getters.currentConn
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

  button
    position absolute
    top 0
    z-index 10
</style>
