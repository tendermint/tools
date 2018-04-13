<template lang="pug">
.panel-sort(:id="id")
  .sort-by(
    v-for="property in sort.properties",
    @click="orderBy(property.value, $event)",
    :class="{ 'active': property.initial, 'asc': property.initial }")
    .text {{ property.title }}
</template>

<script>
import $ from 'jquery'
import shortid from 'shortid'
export default {
  name: 'panel-sort',
  data: () => ({
    id: shortid.generate()
  }),
  methods: {
    orderBy (property, event) {
      let sortBys = '#' + this.id + ' .sort-by'
      $(sortBys).removeClass('active desc asc')
      let el = $(event.target).parent()

      if (this.sort.property === property) {
        if (this.sort.order === 'asc') {
          this.sort.order = 'desc'
        } else {
          this.sort.order = 'asc'
        }
      } else {
        this.sort.property = property
      }
      if (this.sort.order === 'asc') {
        $(el).addClass('asc')
      } else {
        $(el).addClass('desc')
      }
      $(el).addClass('active')
    }
  },
  props: ['sort']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.panel-sort
  display flex
  border 1px solid bc
  border-left none
  margin 0 -2*x
  padding 0 2*x

  faint-bg()

.sort-by
  flex 1
  border-left 1px solid bc
  display flex
  cursor pointer
  user-select none
  box-shadow inset #fff 0 1px 0

  position relative
  &:last-of-type
    border-right 1px solid bc

  padding 0 x*0.75

  .text
    line-height x*2 - 2px
    color txt

  .text
    flex 1

  &:after
    display block
    height x*2 - 2px
    font-family FontAwesome
    color dim
    line-height x*2 - 2px
  &.asc:after
    content '\f0d8'
  &.desc:after
    content '\f0d7'

  &:hover
    background lighten(faint,50%)
    .text
      color link

  &.active
    background #fff
    .text
      color txt
      font-weight bold
    &:after
      color txt
</style>
