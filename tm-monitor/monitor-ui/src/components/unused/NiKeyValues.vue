<template>
  <div :class="cssClass">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'ni-kvs',
  computed: {
    cssClass () {
      let value = 'ni-kvs'
      if (this.size) value += ` ni-kvs-size-${this.size}`
      if (this.type) value += ` ni-kvs-type-${this.type}`
      return value
    }
  },
  props: ['size', 'type']
}
</script>

<style lang="stylus">
@require '~@/styles/variables.styl'

.ni-kvs
  margin 1rem 0
  display flex
  flex-flow column nowrap

  border 2px solid bc-dim
  border-radius 0.25rem

  &.ni-kvs-type-slim
    border-radius 0.25rem
    border 1px solid bc-faint
    display flex
    flex-flow column nowrap

.ni-kvs.ni-kvs-type-slim
  border 2px solid bc-dim

@media screen and (min-width: 768px)
  .ni-kvs.ni-kvs-type-slim
    flex-flow row wrap

    .ni-kv
      flex 1

      &.ni-kv-w2
        flex 2

  .ni-kvs.ni-kvs-size-lg
    border none
    flex-flow row wrap
    padding 0.5rem
    margin -0.5rem -1rem

    .ni-kv.ni-kv-size-lg
      flex 0 0 33.3333%
      .ni-kv-container
        margin 0.5rem
</style>
