<template>
  <div class="create-room">
    <el-card class="create-room__card">
      <el-button type="success" round>Create Room</el-button>
    </el-card>
  </div>
</template>

<script>
export default {
  name: 'CreateRoom',

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
