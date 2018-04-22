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
    id: null,
    ws: null
  }),
  components: {
    VideoStream
  },
  beforeMount() {
    this.id = Math.random().toString(36).substr(2, 9)
    this.ws = new WebSocket(Routes.socketPath(this.$router.history.current.params.id))
  },
  mounted() {
    this.$store.dispatch('checkRoom')
    this.$store.dispatch('setPeer', this.room.created)

    if (!this.room.created) this.ws.onopen = () => this.newObject()

    this.currentPeer.on('signal', e => {
      this.$store.dispatch('setConn', e)
      console.info('command SIGNAL')

      if (!this.room.created) this.connObject()
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

    this.currentPeer.on('connect', () => console.info(`peer conncection created`))
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

    &__flag
      color red

  .stream
    position relative
    z-index 1

  button
    position absolute
    top 0
    z-index 10
</style>
