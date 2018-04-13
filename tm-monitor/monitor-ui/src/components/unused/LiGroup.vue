<template lang="pug">
router-link.li-group(:to="{ name: 'group', params: { group: group.id }}")
  .icon
    i.fa.fa-bank
  .title {{ group.id }}
</template>

<script>
export default {
  name: 'li-group',
  props: ['group']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

a.li-group
  border-top 1px solid bc
  display flex
  color txt
  line-height 3*x

  .icon
    border-right 1px solid bc2
    width x*3 + 1px

    color faint
    text-align center
    background faint

  .title
    padding 0 0.75*x

  &:hover
    background lighten(mc,95%)
</style>
