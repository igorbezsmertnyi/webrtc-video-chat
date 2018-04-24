<template>
  <div class="dock">
    <el-tooltip id="tooltip1" effect="dark" content="Copy room link" placement="top">
      <copy-link />
    </el-tooltip>

    <el-tooltip id="tooltip2" effect="dark" content="Mute audio" placement="top">
      <mute-stream />
    </el-tooltip>

    <el-tooltip id="tooltip3" effect="dark" content="End call" placement="top">
      <end-call />
    </el-tooltip>

    <el-tooltip id="tooltip4" effect="dark" content="Make screenshot" placement="top">
      <screen-shot />
    </el-tooltip>
  </div>
</template>

<script>
import EndCall from './EndCall'
import MuteStream from './MuteStream'
import CopyLink from './CopyLink'
import ScreenShot from './ScreenShot'

export default {
  name: 'Dock',
  components: {
    EndCall,
    MuteStream,
    CopyLink,
    ScreenShot
  }
}
</script>

<style lang="stylus">
  .dock
    position absolute
    width 246px
    bottom 12px
    left 0
    right 0
    margin auto
    padding 8px
    display flex
    flex-flow row nowrap
    background rgba(225, 255, 255, .5)
    border-radius 8px
    z-index 4

</style>
