<template lang="pug">
.ni-card-delegator(:id="delegatorId")
  .ni-cd-container
    .avatar
    .text
      router-link.id(:to="{ name: 'delegator', params: { delegator: delegator.id }}")
        | {{ delegator.id }}
      btn-pubkey(:text="delegator.pub_key")
</template>

<script>
import CardAvatar from './CardAvatar'
import BtnPubkey from './BtnPubkey'
import dateUnixShort from '../scripts/dateUnixShort'
export default {
  name: 'card-delegator',
  components: {
    CardAvatar,
    BtnPubkey
  },
  computed: {
    delegatorId () { return 'delegator-' + this.delegator.id },
    avatarClass () { return 'avatar ' + this.delegator.type },
    delegatorCreatedAt () { return dateUnixShort(this.delegator.created_at) }
  },
  methods: {
    viewDelegator () {
      this.$router.push('/delegators/' + this.delegator.id)
    }
  },
  props: ['delegator']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.ni-cd-container
  display flex
  height 3rem
  align-items center
  dots()

  .text
    flex 1
    display flex
  .id
    flex 1
    cursor pointer
    font-weight 400
    color bright
    white-space nowrap
    overflow hidden

@media screen and (min-width: 768px)
  .ni-card-delegator
    padding 0.5rem

  .ni-cd-container
    height auto
    flex-flow row nowrap
    align-items center
    justify-content center

    padding 0.5rem - px
    border 1px solid bc-faint
    border-radius 0.25rem
    &:before
    &:after
      display none

    .avatar
      flex 0 0 3rem
      margin 0 1rem
      width 3rem
      height 3rem
      border-radius 5rem
      border 1px solid bc
      background alpha(mc, 25%)

    .text
      display block
      flex 1
      height 4rem
      .id
        display block
</style>
