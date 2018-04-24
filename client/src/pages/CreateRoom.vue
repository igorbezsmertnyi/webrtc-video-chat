<template>
  <div class="create-room">
    <div class="create-room__overlay" />

    <div class="create-room__title">
      <h1>WebRTC</h1> 
      <h3>VIDEO CHAT</h3>
    </div>

    <div class="create-room__card">
      <button type="button" @click="createRoom()" >
        Create Room
      </button>
    </div>

    <footer class="create-room__footer">
      <a href="https://github.com/igorbezsmertnyi/webrtc-video-chat" target="_blank">
        <div class="create-room__footer__logo">
          <img src="../assets/gh-logo.png" alt="Git Hub Igor Bezsmertnyi" />
        </div>
        <p>Igor Bezsmertnyi <br /> WebRTC video chat</p>
      </a>
    </footer>
  </div>
</template>

<script>
export default {
  name: 'CreateRoom',
  methods: {
    createRoom() {
      this.$store.dispatch('createRoom')
    }
  }
}
</script>

<style lang="stylus" scoped>
  @import '../styles'

  .create-room
    position relative
    width 100%
    height 100vh
    background url('../assets/bg.jpg')
    background-size cover
    font-family  'Helvetica Neue', Helvetica, 'PingFang SC'

    &__overlay
      position absolute
      width 100%
      height 100%
      top 0
      left 0
      background #000
      opacity .4
      z-index 1

    &__title
      position relative
      padding 20px
      color #fff
      text-align center
      z-index 2

      h1
        margin 0 0 20px
        font-size 4em
        font-weight 200

      h3
        font-weight 300
        font-size 1.4em
        letter-spacing 2px

    &__card
      position relative
      max-width 500px
      margin 80px auto
      border-radius 8px
      padding 80px 40px
      display flex
      justify-content center
      z-index 2

      button
        padding 30px 80px
        color #fff
        text-transform uppercase
        font-size 1em
        border-radius 40px
        border 0
        outline 0
        background linear-gradient(to bottom right, $electric-violet, $aqua)
        cursor pointer
        transition background-color ease-in .2s

        &:hover
          background linear-gradient(to bottom right, $electric-violet, ($aqua 80%))

    &__footer
      position absolute
      width 100%
      left 0
      bottom 0
      padding 20px 0
      z-index 2

      a
        display flex
        align-items center
        flex-direction column
        font-family  'Helvetica Neue', Helvetica, 'PingFang SC'

      &__logo 
        width 64px
        height 64px
        margin-bottom 20px

        img
          width 100%
          height 100%

      p
        text-align center
        color #fff
        text-transform uppercase
        line-height 1.4
</style>
