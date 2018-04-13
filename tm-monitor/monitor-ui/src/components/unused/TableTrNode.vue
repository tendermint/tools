<template lang="pug">
tr.table-tr-node(:id="validatorId")
  td.status
    i.fa.fa-circle.color-online(v-if="node.status.online")
    i.fa.fa-circle.color-offline(v-else)
  td.id {{ node.config.validator.id }}
  td.td-mono {{node.config.p2p_addr}}
  td.td-mono {{node.config.rpc_addr}}
  td.td-mono.latency
    template(v-if="node.status.latency") {{ Math.round(node.status.latency)}} ms
    template(v-else) N/A
  td.td-mono
    template(v-if="node.status.block_height")
      router-link(:to="{ name: 'block', params: { block: node.status.block_height }}")
        | \#{{ node.status.block_height }}
    template(v-else) N/A
  td.public-key: btn-copy(:text="node.config.validator.pub_key[1]")
</template>

<script>
import BtnCopy from './BtnCopy'
export default {
  name: 'table-tr-node',
  components: {
    BtnCopy
  },
  computed: {
    validatorId () {
      return 'node-' + this.node.config.validator.id
    }
  },
  props: ['node']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.table-tr-node
  td
    border 1px solid bc
    padding 0.25*x 0.5*x

  td.td-mono
    mono()
    text-align right

  td.addr
    text-align right

  td.status
    width 2.5*x
    text-align center

  td.public-key
    width 10*x
    .btn
      width 100%
</style>
