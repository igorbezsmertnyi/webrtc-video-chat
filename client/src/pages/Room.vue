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
    ws: null,
    peer: null,
    conn: null,
    otherConn: null
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

    this.peer = new Peer({
      initiator: this.room.created,
      trickle: false
    })

    if (!this.room.created) this.ws.onopen = () => this.newObject()

    this.peer.on('signal', e => {
      this.conn = e

      console.info('commang SIGNAL')
      if (!this.room.created) this.connObject()
    })

    this.ws.onmessage = e => {
      switch (JSON.parse(e.data).command) {
        case 'NEW':
          this.newCommand(JSON.parse(e.data))
          break
        case 'CONN':
          this.connCommand(JSON.parse(e.data))
          break 
      }
    }

    this.peer.on('connect', () => console.info(`peer conncection created`))
  },
  methods: {
    newCommand(data) {
      if (data.id === this.id || !this.room.created) return

      console.info('commang NEW')
      this.connObject()
    },

    connCommand(data) {
      if (data.id === this.id) return

      console.info('commang CONN')
      this.peer.signal(data.peer)
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
        peer: this.conn
      })

      this.ws.send(obj)
    }
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

  button
    position absolute
    top 0
    z-index 10
</style>
