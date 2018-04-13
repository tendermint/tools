<template>
  <div :class="cssClass">
    <div class="ni-kv-container">
      <div class="ni-kv-key"><slot name="key"></slot></div>
      <div class="ni-kv-value"><slot name="value"></slot></div>
      <div class="ni-kv-full"><slot name="full"></slot></div>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ni-kv',
  computed: {
    cssClass () {
      let value = 'ni-kv'
      if (this.flush) value += ' ni-kv-type-flush'
      if (this.size) value += ` ni-kv-size-${this.size}`
      if (this.type) value += ` ni-kv-type-${this.type}`
      else value += ` ni-kv-type-default`
      if (this.width) value += ` ni-kv-w${this.width}`
      return value
    }
  },
  props: ['flush', 'size', 'type', 'width']
}
</script>

<style lang="stylus">
@require '~@/styles/variables.styl'

.ni-kv
  min-width 0

.ni-kv-key:empty
.ni-kv-value:empty
  display none

/*============================================================================*/
// default style

.ni-kv.ni-kv-type-default
  border-bottom 1px solid bc-faint
  height 2.5rem

  &:last-of-type
    border-bottom none

  .ni-kv-container
    display flex

  .ni-kv-key
  .ni-kv-value
    height 2.5rem - px
    padding 0 0.5rem

    font-label()
    color dim

    display flex
    align-items center

    min-width 0

    > div
      white-space nowrap
      overflow hidden
      text-overflow ellipsis

  .ni-kv-key
    flex 1

    > div
      font-size 0.666rem
      letter-spacing 0.025em

  .ni-kv-value
    flex 2
    text-transform none
    
    > div
      font-size 0.875rem
      color bright

    > div a
      color accent1

@media screen and (min-width: 768px)
  .ni-kv.ni-kv-type-default
    .ni-kv-value
      flex 3

  .ni-kv.ni-kv-type-default.ni-kv-size-lg
    height auto
    border-bottom none
    position relative

    .ni-kv-container
      height auto

      padding 1rem - 2*px 0
      border 2px solid bc-dim
      border-radius 0.25rem
      display flex
    
    .ni-kv-container
      flex 1
      display block

    .ni-kv-key
    .ni-kv-value
      height auto
      font-size 0.75rem
      padding-left 1rem - 2*px
      padding-right 1rem - 2*px

    .ni-kv-key
      border-bottom 1px solid bc-faint
      display flex
      align-items center
      height 2rem

    .ni-kv-value
      > div
        font-size 1.75rem
        line-height 3rem

/*============================================================================*/
// slim style

.ni-kv.ni-kv-type-slim
  border-bottom 1px solid bc-faint
  height 2.5rem

  &:last-of-type
    border-bottom none

  .ni-kv-container
    display flex

  .ni-kv-key
  .ni-kv-value
    height 2.5rem - px
    padding 0 0.5rem

    font-label()
    color dim

    display flex
    align-items center

    min-width 0

    > div
      white-space nowrap
      overflow hidden
      text-overflow ellipsis

  .ni-kv-key
    flex 1

    > div
      font-size 0.666rem
      letter-spacing 0.025em

  .ni-kv-value
    flex 2
    text-transform none
    
    > div
      font-size 0.875rem
      color bright

    > div a
      color accent1

  .ni-kv-key:empty
  .ni-kv-value:empty
  .ni-kv-full:empty
    display none

  .ni-kv-full
    flex 1

@media screen and (min-width: 768px)
  .ni-kv.ni-kv-type-slim + .ni-kv.ni-kv-type-slim
    border-left none

  .ni-kv.ni-kv-type-slim
    height 4rem
    overflow hidden

    &.nv-kv-type-flush > div
      padding 0 0.5rem

    .ni-kv-container
      flex-flow column nowrap

    .ni-kv-key:not(:empty)
      font-label()
      color dim
      padding 0 0.5rem
      height 1.5rem
      display flex
      align-items center

    .ni-kv-value:not(:empty)
      padding 0

    .ni-kv-value:not(:empty) > div
      padding 0 0.5rem

      font-weight 300
      font-size 1.25rem
      line-height 2.5rem
      color bright

      display flex
      align-items center

      a
        color bright
        positon relative
        &:after
          content 'link'
          padding 0 0.5rem
          font-family 'Material Icons'
          color faint
          font-size 1rem
          line-height 1rem
        &:hover
          color link
        &:hover:after
          color link
</style>
