<template lang="pug">
page(title='Blockchain')
  // blockchain-select-modal
  // tool-bar
    router-link(to="/search" exact): i.material-icons search
    a(@click='toggleBlockchainSelect')
      i.material-icons(v-if='!config.blockchainSelect') filter_list
      i.material-icons(v-else='') close

  part(title='Blockchain')
    list-item(dt='Network' :dd='bc.status.node_info.network')
    list-item(dt='Tendermint Version' :dd='bc.status.node_info.version')
    list-item(dt='Peer Nodes' :dd='validators.length')

  part(title='Current Block')
    list-item(dt='Block Height' :dd='num.prettyInt(bc.status.latest_block_height)'
      :to="{ name: 'block', params: { block: bc.status.latest_block_height} }")
    list-item(dt='Latest Block Time' :dd='readableDate(bc.status.latest_block_time)')
    list-item(dt='Latest Block Hash' :dd='bc.status.latest_block_hash')

  part(title='Connected Node')
    // list-item(label='Node IP' :input='bc.connectedIp')
    list-item(dt='Node IP' dd='206.189.22.179')
    list-item(dt='Node Moniker' :dd='bc.status.node_info.moniker')

  // part(title='Nodes')
    // list-item(dt='Current Rate' :dd="currentRate + ' bytes/s'")
    // list-item(dt='Average Rate' :dd="averageRate + ' bytes/s'")
</template>

<script>
import moment from 'moment'
import num from '../scripts/num'
import { mapGetters } from 'vuex'
import BlockchainSelectModal from './BlockchainSelectModal'
import ListItem from './NiListItem'
import Page from './NiPage'
import Part from './NiPart'
import ToolBar from './NiToolBar'
export default {
  name: 'page-index',
  components: {
    BlockchainSelectModal,
    ListItem,
    Page,
    Part,
    ToolBar
  },
  computed: {
    ...mapGetters(['blockchain', 'config', 'validators']),
    bc () { return this.blockchain },
    avgTxThroughput () {
      return Math.round(this.bc.network.avg_tx_throughput * 1000) / 1000
    }
  },
  data: () => ({
    moment: moment,
    num: num
  }),
  methods: {
    readableDate (ms) {
      return moment(ms).format('YYYY-MM-DD h:mm:ss A')
    },
    toggleBlockchainSelect () {
      this.$store.commit('SET_CONFIG_BLOCKCHAIN_SELECT', !this.config.blockchainSelect)
    }
  },
  mounted () {
    console.log(this.blockchain)
  }
}
</script>
