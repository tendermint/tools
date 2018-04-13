<template lang="pug">
.nav-validator
  router-link.title(:to="{ name: 'entity', params: { entity: validator.id }}")
    | #[i.fa.fa-bank] {{ validator.id }}
</template>

<script>
export default {
  name: 'nav-validator',
  props: ['validator']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

</style>

