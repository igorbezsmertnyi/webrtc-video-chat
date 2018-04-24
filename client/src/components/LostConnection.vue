<template>
  <div class="lost-conn">
    <svg-icon src="network" />
    <p>
      Oops, connection lost. Please wait we try to reconnect you
    </p>
  </div>
</template>

<script>
import SvgIcon from '@/components/SvgIcon'

export default {
  name: 'LostConnection',
  components: {
    SvgIcon
  }
}
</script>

<style lang="stylus">
  @import '../styles'

  .lost-conn
    position absolute
    width 100%
    height 100%
    top 0
    left 0
    display flex
    justify-content center
    align-items center
    background $imperial-primer
    flex-flow column
    z-index 3

    svg 
      width 80px
      height 80px

    p
      padding 0 14px
      color #fff
</style>
