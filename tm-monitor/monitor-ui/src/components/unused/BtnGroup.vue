<template>
  <div class="ni-btn-group">
    <slot></slot>
  </div>
</template>

<script>
export default {
  name: 'ni-btn-group'
}
</script>

<style lang="stylus">
.ni-btn-group
  margin-bottom 0.5rem
  max-width 40rem

  display flex
  .ni-btn-wrapper
    flex 1

  .ni-btn-wrapper + .ni-btn-wrapper
    .ni-btn
      border-left none

@media screen and (min-width: 768px)
  .ni-btn-group
    margin-bottom 0.75rem
</style>
