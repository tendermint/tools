<template lang="pug">
.group-tag(:class="scale")
  .wrapper
    i.fa.fa-bank
    .title {{ groupId }}
</template>

<script>
export default {
  name: 'group-tag',
  props: ['group-id', 'scale']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.group-tag
  display inline-block
  margin 0.125*x 0.3333*x 0.125*x 0

  .wrapper
    background lighten(mc,25%)
    border-radius 0.125*x
    height 1.75*x
    padding 0 0.5*x
    display flex
    align-items center

  i.fa
    color lighten(mc,66%)
    font-size 0.85*x
    margin-right 0.333*x

  .title
    color lighten(mc,90%)
    font-weight normal
    display inline
  &.large
    margin 0.125*x 0.6666*x 0.125*x 0
    .wrapper
      height 2.25*x
      padding 0 0.666*x
    i.fa
      font-size 1.125*x
      margin-right 0.4*x
    .title
      font-weight 600
</style>
