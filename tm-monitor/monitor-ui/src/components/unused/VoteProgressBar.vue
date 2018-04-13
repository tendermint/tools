<template lang="pug">
.vote-progress-bar(v-bind:class="{ 'large': isLarge }", )
  .outer-bar
    .inner-bar.inner-bar-none(v-if="noVotes.length + yesVotes.length === 0") No votes yet.
    .inner-bar.inner-bar-yes(v-if="yesVotes.length",
      :style="'width:' + yesVotes.length / entityCount * 100 + '%'")
      .label 
        i.fa.fa-check
        | {{ Math.round(yesVotes.length / entityCount * 100) }}%
    
    .inner-bar.inner-bar-no(v-if="noVotes.length > 0",
      :style="'width:' + noVotes.length / entityCount * 100 + '%'")
      .label
        i.fa.fa-times
        | {{ Math.round(noVotes.length / entityCount * 100) }}%
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'vote-progress-bar',
  computed: {
    ...mapGetters(['entities', 'votes']),
    entityCount () { return this.entities.length },
    filteredVotes () {
      if (this.votes) {
        return this.votes.filter(v => v.proposal_id === this.proposalId)
      } else {
        return []
      }
    },
    noVotes () {
      return this.filteredVotes.filter(v => v.value === -1)
    },
    yesVotes () {
      return this.filteredVotes.filter(v => v.value === 1)
    }
  },
  props: ['proposal-id', 'is-large']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.vote-progress-bar
  height 33px
  background bc2
  
  .outer-bar, .inner-bar
    height x
    line-height x

  .outer-bar
    background lighten(mc, 10%)
    display flex

  .inner-bar
    text-align center
    border-left 1px solid bc2
    position relative
    transition 0.3s ease-in-out width

    .label
      position absolute
      bottom -1*x
      left 0
      background hsla(0,0,0,0.666)
      background #fff
      color txt
      font-size x*0.6666
      line-height x
      padding 0 x*0.3
      display flex
      i.fa
        margin-right 0.175*x
        line-height x
        &.fa-check
          color btn-green-color
        &.fa-times
          color red

  .inner-bar-none
    color lighten(link,50%)
    font-size 0.66*x
    padding 0 0.75*x

  .inner-bar-yes
    box-shadow inset lighten(link,50%) 0 1px 0
    background linear-gradient(top, lighten(link,10%), darken(link,15%))
  .inner-bar-no
    box-shadow inset lighten(red,50%) 0 1px 0
    background linear-gradient(top, lighten(red,10%), darken(red,15%))

  &.large
    height 60px
    .outer-bar, .inner-bar
      height 2.25*x
      height 2.25*x
    .inner-bar
      .label
        font-size 0.75*x
        line-height x*1.5
        bottom -1.5*x
        i.fa
          line-height x*1.5
    .inner-bar-none
      font-size x
      line-height 2.25*x
</style>
