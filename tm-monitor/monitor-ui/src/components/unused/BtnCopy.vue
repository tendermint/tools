<template lang="pug">
btn.btn-copy(theme="cosmos", icon="clipboard", value="Copy", :data-clipboard-text="text")
</template>

<script>
/* eslint-disable no-new */
import Clipboard from 'clipboard'
import Btn from '@nylira/vue-button'
export default {
  name: 'btn-copy',
  components: {
    Btn
  },
  mounted () {
    new Clipboard('.btn-copy')
  },
  props: ['text']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.btn-copy
  position relative

  &:active
    .ni-btn
      background ac
      color #fff
</style>
