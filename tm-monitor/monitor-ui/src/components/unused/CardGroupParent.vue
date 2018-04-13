<template lang="pug">
.card-group-parent
  header
    router-link.title(:to="{ name: 'group', params: { group: group.id }}")
      | #[i.fa.fa-bank] {{ group.id }}

  .content(v-if="childGroups.length > 0")
    .child-groups
      card-group-child(v-for="childGroup in childGroups", :key="group.id", :group="childGroup")
</template>

<script>
import { mapGetters } from 'vuex'
import CardGroupChild from './CardGroupChild'
export default {
  name: 'card-group-parent',
  components: {
    CardGroupChild
  },
  computed: {
    ...mapGetters(['groups']),
    childGroups () {
      let self = this
      return this.groups.filter(function (group) {
        return group.parent_id === self.group.id
      })
    }
  },
  props: ['group']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.card-group-parent
  border 1px solid bc
  margin 0.5*x

  header
    height 3*x

    a.title
      display block
      padding 0 x

      font-size 1.125*x
      font-weight 500
      line-height 3*x

      display flex
      align-items center

      color txt

      i.fa
        color faint
        font-size x
        margin-right 0.5*x
        
      &:hover
        color link
        background lighten(mc,95%)
        i.fa
          color lighten(link,50%)

  .content
    border-top 1px solid bc2
    padding 1rem 1rem 0.5rem

  .child-groups
    display flex 
    flex-flow row wrap
</style>
