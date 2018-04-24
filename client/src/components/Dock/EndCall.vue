<template>
  <div class="dock-btn end-call">
    <button type="button" @click="destroyPeer()">
      <svg-icon src="phone" />
    </button>
  </div>
</template>

<script>
import SvgIcon from '@/components/SvgIcon'

export default {
  name: 'EndCall',
  components: {
    SvgIcon
  },
  methods: {
    destroyPeer() {
      window.location.href = '/'
    }
  }
}
</script>

<style lang="stylus">
  @import '../../styles'

  .end-call
    background transparentify($amour, .8)

    &:hover
      background transparentify($amour, .6) !important
</style>
