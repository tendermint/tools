<template lang="pug">
router-link.card-entity-link(:to="{ name: 'entity', params: { entity: entityId }}")
  .avatar-wrapper: card-avatar(:entity-id="entityId")
  .id {{ entityId }}
</template>

<script>
import CardAvatar from './CardAvatar'
export default {
  name: 'card-entity-link',
  components: {
    CardAvatar
  },
  props: ['entity-id']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.card-entity-link
  flex 1
  display flex
  padding 0.25*x
  color txt

  .avatar-wrapper
    margin-right 0.5*x

  .id
    line-height 2*x
    flex 1

  &:hover
    color link
</style>
