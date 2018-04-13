<template lang='pug'>
page(title='Delegators')
  list-item(
    :to="'/delegators/' + delegator.id"
    v-for='delegator in filteredDelegators'
    :key='delegator.id'
    icon='person'
    :title='delegator.id')
</template>

<script>
import { mapGetters } from 'vuex'
import { orderBy } from 'lodash'
import smoothScroll from 'smooth-scroll'
import ListItem from './NiListItem'
import Page from './NiPage'
import Part from './NiPart'
export default {
  name: 'page-delegators-index',
  components: {
    ListItem,
    Page,
    Part
  },
  computed: {
    ...mapGetters(['delegators']),
    filteredDelegators () {
      return orderBy(this.delegators, [this.sort.property], [this.sort.order])
    }
  },
  data: () => ({
    sort: {
      property: 'id',
      order: 'asc',
      properties: [
        { id: 1, title: 'Type', value: 'type', hidden: true },
        { id: 2, title: 'ID', value: 'id', initial: true },
        { id: 3, title: 'Public Key', value: 'pub_key', disabled: true },
        { id: 4, title: 'Info', value: 'info' },
        { id: 5, title: 'Created', value: 'created_at' }
      ]
    }
  }),
  methods: {
    smoothScrollOnLoad () {
      if (window.location.hash) {
        // console.log('theres a hash!')
        let hash = smoothScroll.escapeCharacters(window.location.hash)
        let toggle = document.querySelector('a[href*="' + hash + '"]')
        let options = { offset: 65 }
        smoothScroll.animateScroll(window.location.hash, toggle, options)
        history.replaceState('', document.title, window.location.pathname)
      }
    }
  },
  mounted () {
    this.smoothScrollOnLoad()
  }
}
</script>
