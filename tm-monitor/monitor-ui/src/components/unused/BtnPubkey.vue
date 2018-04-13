<template lang="pug">
btn.btn-pubkey(
  theme="cosmos",
  icon="content_copy",
  value="Pubkey",
  size="sm"
  :data-clipboard-text="text")
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

.btn-pubkey
  position relative
</style>
