<template lang="pug">
tr.table-tr-header(:id="id")
  th(v-for="property in sort.properties")
    .label.disabled(v-if="property.disabled === true"): .text {{ property.title }}
    .label(
      v-else
      @click="orderBy(property.value, $event)",
      :class="{ 'active': property.initial, 'asc': property.initial }")
      .text(v-if="property.hidden !== true") {{ property.title }}
      i.fa.fa-sort
      i.fa.fa-sort-desc
      i.fa.fa-sort-asc
</template>

<script>
import $ from 'jquery'
import shortid from 'shortid'
export default {
  name: 'table-tr-header',
  data () {
    return {
      id: 'tth-' + shortid.generate()
    }
  },
  methods: {
    orderBy (property, event) {
      let sortBys = '#' + this.id + ' .label'
      $(sortBys).removeClass('active desc asc')

      let el
      if ($(event.target).hasClass('label')) {
        el = $(event.target)
      } else {
        el = $(event.target).parent()
      }
      // console.log(el)

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

.table-tr-header
  th
    border 1px solid bc

  .label
    cursor pointer
    display flex
    background bc

    .text
      flex 1
      display inline-block
      line-height 2.5*x
      padding-left x*0.5

    i.fa
      width x

      width 2.5*x
      line-height 2.5*x

      text-align center
      color faint

      &.fa-sort-asc
      &.fa-sort-desc
        display none

    &:hover
      background lighten(faint,50%)
      .text
        color link

    &.disabled
      cursor default
      &:hover
        background faint
        .text
          color txt

    &.active
      background #fff
      .text
        color txt
        font-weight bold
      &:after
        color txt

    &.asc, &.desc
      i.fa
        color txt

    &.asc
      i.fa.fa-sort-desc, i.fa.fa-sort
        display none
      i.fa.fa-sort-asc
        display inline-block

    &.desc
      i.fa.fa-sort-desc
        display inline-block
      i.fa.fa-sort-asc, i.fa.fa-sort
        display none
</style>
