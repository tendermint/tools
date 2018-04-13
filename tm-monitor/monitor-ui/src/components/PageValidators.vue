<template lang="pug">
page(title='Peer Nodes')
  tab-bar
    router-link(to="/validators" exact) Online ({{ online }})
    router-link(to="/validators/offline" exact) Offline (0)
  // tool-bar
    a(@click='toggleSearch'): i.material-icons search
    a(@click='toggleFilter'): i.material-icons filter_list
  list-item(
    v-for="v in values"
    :key="v.node_info.moniker"
    :title="v.node_info.moniker"
    :subtitle="getIp(v)"
    icon='storage'
    :to="`/validators/${urlsafeIp(getIp(v))}`")
</template>

<script>
import { mapGetters } from 'vuex'
import { orderBy } from 'lodash'
import ListItem from './NiListItem'
import Page from './NiPage'
import TabBar from './NiTabBar'
import ToolBar from './NiToolBar'
export default {
  name: 'page-validators',
  components: {
    ListItem,
    Page,
    TabBar,
    ToolBar
  },
  data: () => ({
    todoAtoms: '13.37M ATOM'
  }),
  computed: {
    ...mapGetters(['validators']),
    values () {
      return orderBy(this.validators, ['node_info.moniker', 'desc'])
    },
    online () { return this.validators.length }
  },
  methods: {
    toggleFilter () {
      this.$store.commit('notify', { title: 'Filtering...', body: 'TODO' })
    },
    toggleSearch () {
      this.$store.commit('notify', { title: 'Searching...', body: 'TODO' })
    },
    urlsafeIp (ip) {
      return ip.split('.').join('-')
    },
    getIp (validator) {
      return validator.node_info.listen_addr.split(':')[0]
    }
  }
}
</script>
