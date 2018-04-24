<template>
  <div 
    v-show="show" 
    class="screen-img"
    :class="{
      'screen-img--animate': animate
    }"
  >
    <div class="screen-img__frame">
      <img :src="img" @click="download()" />
    </div>
  </div>
</template>

<script>
export default {
  name: 'ScreenImg',
  data: () => ({
    img: null,
    show: false,
    animate: false
  }),
  watch: {
    currentImg(val) { 
      if (!val) return

      this.img = val 
      this.show = true

      setTimeout(() => this.animate = true, 400)
    }
  },
  methods: {
    download() {
      const download = document.createElement('a')
      download.setAttribute('href', this.img)
      download.setAttribute('download', 'image.png')
      download.click()
      download.remove()

      this.img = null
      this.show = false
      this.animate = false

      this.$store.dispatch('removeImage')
    }
  },
  computed: {
    currentImg() {
      return this.$store.getters.currentScreenShot
    }
  }
}
</script>

<style lang="stylus">
  @import '../styles'

  .screen-img
    position absolute
    width calc(100vw - 40px)
    height calc(100vh - 40px)
    left 0
    bottom 0
    padding 20px
    z-index 10
    transition transform linear .4s

    &--animate
      transform scale(0.2)
      transform-origin left bottom

    &__frame
      width calc(100% - 16px)
      height calc(100% - 16px)
      background #fff
      padding 8px
      border-radius 8px
      overflow hidden
      box-shadow -4px 7px 8px 0px rgba(0, 0, 0, 0.25)

      img 
        width 100%
        height 100%
        object-fit cover
</style>
