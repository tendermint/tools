<template lang="pug">
.card-avatar(:class="{ 'avatar-editable': editable }")
  template(v-if="validator.type === 'user'")
    // img(v-if="validator.id == me.id", :src="session.password.profileImageURL")
    .icon.icon-user: i.fa.fa-user
  template(v-if="validator.type === 'machine'")
    .icon.icon-machine
      i.fa.fa-bolt
  template(v-if="validator.type === 'undefined'")
    .icon.icon-undefined
      i.fa.fa-spinner.fa-spin
  .overlay(@click="showAvatarMenu")
    i.fa.fa-cloud-upload
    .label Change avatar
</template>

<script>
import { mapGetters } from 'vuex'
export default {
  name: 'card-avatar',
  computed: {
    ...mapGetters(['me', 'validators']),
    validator () {
      let thisEntity = this.validators.find(e => e.id === this.validatorId)
      if (thisEntity) {
        return thisEntity
      } else {
        return { type: 'undefined' }
      }
    }
  },
  methods: {
    showAvatarMenu () {
      alert('TODO: Show avatar menu')
    }
  },
  props: ['validator-id', 'editable']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

user-hue = 130
user-color = hsl(user-hue,100%,35%)

machine-hue = mhue
machine-color = hsl(machine-hue,100%,45%)

.card-avatar
  img, .icon
    width 2*x
    height 2*x

  .icon
    display flex
    justify-content center
    align-items center

    &.icon-user
      background user-color
      i.fa
        color hsl(user-hue, 60%, 80%)

    &.icon-machine
      background machine-color
      i.fa
        color hsl(machine-hue, 100%, 85%)

    &.icon-undefined
      background bc

    i.fa
      font-size x

  .overlay
    display none

.card-avatar.avatar-editable
  position relative
  cursor pointer

  width 4.5*x
  height 4.5*x
  margin-right x

  img, .icon
    width 100%
    height 100%
    i.fa
      font-size 2.5*x

  .overlay
    border 2px solid hsla(0,0,0,0.75)
    background hsla(0,0,0,0.5)
    position absolute
    top 0
    left 0

    height 100%
    width 100%

    display flex
    flex-flow column
    justify-content center

    text-align center
    transition 0.3s ease-in-out all
    opacity 0

    i.fa
      color #fff
      font-size 1.5*x
      text-shadow #000 0 0.125*x 0.25*x
    .label
      font-size 0.75*x
      color #fff
      line-height 1.25em
      text-shadow #000 0 0.125*x 0.25*x

  &:hover
    .overlay
      opacity 1
</style>
