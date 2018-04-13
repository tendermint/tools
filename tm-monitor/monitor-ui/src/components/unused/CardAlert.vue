<template lang="pug">
router-link.card-alert(to="{ path: 'proposals', params: { proposal: proposal.id }}")
  .wrap
    .icon
      i.fa.fa-exclamation-circle
    .title
      span.type
        | Proposed #{' '}
        template(v-if="proposal.type === 'Validator Set'") {{ proposal.type }}
        template(v-else) {{ proposal.type | getApplicationType }}
        | #{' '} Launch:
      span.title {{ proposal.title }}
      span.date.light {{ proposal.created_at | dateUnixAgo }}
    .author
      span.light {{ proposal.entity_id }}
      img(:src=" proposal.entity_id | getEntityAvatar ")
  .btn: i.fa.fa-search-plus
</template>

<script>
export default {
  name: 'card-alert',
  props: ['proposal']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.card-alert
  display block
  margin 0 0 0.5em
  display flex

  &:hover
    .wrap, .icon
      border-color link
    .icon i.fa
      color link

  &:last-of-type
    margin-bottom 0 
  .wrap
    flex 1
    border 1px solid bc
    margin-right x*0.5
    display flex

    line-height x*2 - 2px

  .icon
    width x*2
    text-align center
    border-right 1px dotted bc
    i.fa
      color alert-color

  .title
    flex 1
    padding 0 x*0.5

    span.type, span.title
      color txt

    span.title
      font-weight bold

  .author
    display flex
    img
      display block
      width 28px
      height 28px
      margin 1px
      margin-left x*0.5
</style>
