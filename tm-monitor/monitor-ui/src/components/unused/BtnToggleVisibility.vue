<template lang="pug">
.btn.btn-large.btn-toggle-visibility(@click="toggleVisibility",
  v-bind:class="{ 'invisible': invisible }")
  span.visible #[i.fa.fa-eye]
  span.invisible #[i.fa.fa-eye-slash]
</template>

<script>
export default {
  name: 'btn-toggle-visibility',
  computed: {
    input () {
      let field = document.querySelector('#' + this.fieldId)
      return field
    }
  },
  data () {
    return {
      invisible: false
    }
  },
  methods: {
    toggleVisibility () {
      if (this.input.type === 'password') {
        this.input.type = 'text'
        this.invisible = false
      } else {
        this.input.type = 'password'
        this.invisible = true
      }
    }
  },
  props: ['field-id']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.btn-toggle-visibility
  width 3*x
  padding-left 0
  padding-right 0

  cursor pointer
  user-select none

  span.invisible
    display none

  &.invisible

    span.visible
      display none

    span.invisible
      display block
</style>
