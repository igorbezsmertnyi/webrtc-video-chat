<template>
  <div 
    class="dock-btn mute-stream"
    :class="{
      'dock-btn--active': active
    }"
  >
    <button type="button" @click="muteStrem()">
      <svg-icon v-show="active" src="microphone-mute" />
      <svg-icon v-show="!active" src="microphone" />
    </button>
  </div>
</template>

<script>
import SvgIcon from '@/components/SvgIcon'

export default {
  name: 'MuteStream',
  data: () => ({
    active: false
  }),
  components: {
    SvgIcon
  },
  methods: {
    muteStrem() {
      this.active = !this.active
      this.currentPeer.stream.getAudioTracks()[0].enabled = !this.active
    }
  },
  computed: {
    currentPeer() {
      return this.$store.getters.currentPeer
    }
  }
}
</script>

<style lang="stylus">
  @import '../../styles'

  .mute-stream
    background transparentify(#fff, .8)
</style>
