<template>
  <div class="create-room">
    <el-card class="create-room__card">
      <el-button @click="createRoom()" type="success" round>Create Room</el-button>
    </el-card>
  </div>
</template>

<script>
import Peer from 'simple-peer'

export default {
  name: 'CreateRoom',
  data: () => ({
    peer: null
  }),
  mounted() {
    const peer = new Peer({
      initiator: true,
      trickle: false
    })

    peer.on('signal', e => this.peer = e)
  },
  methods: {
    createRoom() {
      this.$store.dispatch('createRoom', this.peer)
    }
  }
}
</script>

<style lang="stylus" scoped>
  .create-room
    width 100%
    height 100vh
    display flex
    justify-content center
    align-items center

    &__card
      padding 32.36px 48.54px
</style>
