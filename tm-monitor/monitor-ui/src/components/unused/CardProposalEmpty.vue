<template lang="pug">
.card-proposal-empty
  .title There are no proposals for this group.
  .body Create one by clicking on the #[i.fa.fa-plus] button located in the top right of this page.
</template>

<script>
export default {
  name: 'card-proposal-empty'
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.card-proposal-empty
  margin-top 2*x
  border 1px solid bc
  background faint
  padding 4*x
  display flex

  flex-flow column nowrap
  justify-content center

  border 1px solid lighten(mc,80%)
  background lighten(mc,95%)

  .title, .body
    max-width 30*x

  .title
    font-size x*1.5
    margin 0 0 0.5*x
    font-weight 600
    color lighten(mc,20%)

  .body
    font-weight 200
    font-size 1.25*x
    color dim
</style>
