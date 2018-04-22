<template>
  <svg v-html="svg" />
</template>

<script>
export default {
  name: 'SvgIcon',
  props: ['src'],
  data: () => ({
    svg: null
  }),
  beforeMount() {
    const img = require(`./../assets/${this.src}.svg`)
    this.svg = Buffer.from(img.substr(img.indexOf(',') + 1, img.length), 'base64').toString()
  }
}
</script>
