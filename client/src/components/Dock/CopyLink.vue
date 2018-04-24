<template>
  <div class="dock-btn copy-link">
    <button type="button" @click="copyLink()">
      <svg-icon src="copy" />
      <input ref="link" type="text" :value="link" />
    </button>
  </div>
</template>

<script>
import SvgIcon from '@/components/SvgIcon'

export default {
  name: 'CopyLink',
  data: () => ({
    link: null
  }),
  components: {
    SvgIcon
  },
  beforeMount() {
    this.link = window.location.href
  },
  methods: {
    copyLink() {
      this.$refs.link.value = this.link
      this.$refs.link.select()      
      document.execCommand('copy')

      this.$message({
        showClose: true,
        message: 'Room link copied',
        type: 'success'
      })
    }
  }
}
</script>

<style lang="stylus">
  @import '../../styles'

  .copy-link
    background transparentify(#fff, .8)

    input
      position absolute
      opacity 0
      bottom -100%
</style>
