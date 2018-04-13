<template lang="pug">
.vote-power-bar
  .power-bar(:style="'width:' + power * 10 + '%'") &nbsp;
  .label &times; {{ voteLabel }}
</template>

<script>
import pluralizeWithCount from '../scripts/pluralizeWithCount'
export default {
  name: 'vote-power-bar',
  computed: {
    voteLabel () {
      return pluralizeWithCount('Vote', this.power)
    }
  },
  props: ['power']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.vote-power-bar
  background darken(mc,30%)
  border 1px solid lighten(mc,10%)
  position relative
  height 1.5*x
  padding 2px

  .power-bar
    width 50%
    background linear-gradient(top, lighten(mc,30%), lighten(mc,10%))
    box-shadow inset 0 1px 0 lighten(mc,45%)
    height 1.5*x - 6px

  .label
    position absolute
    top - 1px
    left 50%
    width 8*x
    margin-left -4*x
    text-align center

    font-size x*0.75
    font-weight 600
    color lighten(mc,75%)
    line-height 1.5*x
    text-shadow hsla(0,0,0,0.5) 0 1px 0
</style>
