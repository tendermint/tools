<template lang="pug">
router-link.card-group-child(:to="{ name: 'group', params: { group: group.id }}")
  .icon: i.fa.fa-bank
  .title {{ group.id }}
</template>

<script>
export default {
  name: 'card-group-child',
  props: ['group']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.card-group-child
  display block
border 1px solid darken(mc,20%)

  color txt
  line-height 2*x
  margin-right 0.5rem
  margin-bottom 0.5rem

  display flex

  box-shadow inset lighten(mc,20%) 0 1px 0
  background linear-gradient(top, lighten(mc,10%), darken(mc,10%))
  min-width 15rem

  .icon
    border-right 1px solid darken(mc,20%)
    padding 0 0.5*x
    i.fa
      color lighten(mc,25%)

  .title
    border-left 1px solid lighten(mc,10%)
    padding 0 x 0 x*0.5
    color lighten(mc,80%)

  &:hover
    background linear-gradient(top, lighten(mc,15%), darken(mc,5%))
</style>
