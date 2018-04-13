<template lang="pug">
.nav-group
  router-link.title(:to="{ name: 'group', params: { group: group.id }}")
    | #[i.fa.fa-bank] {{ group.id }}

  ul(v-if="childGroups.length > 0")
    li(v-for="cg in childGroups", :key="cg.id")
      router-link(:to="{ name: 'group', params: { group: cg.id }}")
        | #[i.fa.fa-angle-right] {{ cg.id }}
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'nav-group',
  computed: {
    ...mapGetters(['groups']),
    childGroups () {
      return this.groups.filter(g => g.parent_id === this.group.id)
    },
    parentGroupLink () {
      return { name: 'group', params: { group: this.group.id } }
    }
  },
  methods: {
    groupLink (groupId) {
      return { name: 'group', params: { group: groupId } }
    }
  },
  props: ['group']
}
</script>
