<template lang="pug">
transition(name="vote-entity"): router-link.card-vote-entity(
  :to="{ name: 'entity', params: { entity: vote.entity_id }}",
  :class="{ 'entity-me': isMe }")
  card-avatar(:entity-id="vote.entity_id")
  .text {{ vote.entity_id }}
  .date(:title="voteUpdatedAt") {{ voteUpdatedAtAgo }}
</template>

<script>
import CardAvatar from './CardAvatar'
import dateUnix from '../scripts/dateUnix'
import dateUnixAgo from '../scripts/dateUnixAgo'
export default {
  name: 'card-vote-entity',
  components: {
    CardAvatar
  },
  computed: {
    isMe () {
      if (this.vote.entity_id === 1) {
        return true
      }
    },
    voteUpdatedAt () { return dateUnix(this.vote.updated_at) },
    voteUpdatedAtAgo () { return dateUnixAgo(this.vote.updated_at) }
  },
  props: ['vote']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.card-vote-entity
  display block
  height x*2.5
  padding 0.25*x

  display flex
  align-items center

  user-select none

  &:last-of-type
    border-bottom none

  .avatar
    width 2*x
    height 2*x
    margin-right 0.5*x

  .text
    flex 1
    color txt

  .date
    text-align right
    font-size x*0.75
    color faint
    padding-right 0.25*x

  &.entity-me
    transition 0.15s ease all
    .text
      font-weight bold

  &:hover
    background lighten(mc,90%)
    .text
      color link
</style>
