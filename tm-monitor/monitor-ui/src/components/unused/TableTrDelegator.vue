<template lang="pug">
tr.table-tr-validator(:id="validatorId")
  // td(:class="avatarClass"): card-avatar(:validator-id="validator.id")
  td.id(@click="viewValidator") #[i.fa.fa-validator] {{ validator.id }}
  td.public-key
    btn-copy(:text="validator.pub_key")
  //
    td.info
      template(v-if="validator.info === ''") --
      template(v-else) {{ validator.info }}
    td.voting-power(v-if="validator.votingPower") {{ validator.votingPower }}&times;
    td.created(v-else) {{ validatorCreatedAt }}
</template>

<script>
import CardAvatar from './CardAvatar'
import BtnCopy from './BtnCopy'
import dateUnixShort from '../scripts/dateUnixShort'
export default {
  name: 'table-tr-validator',
  components: {
    CardAvatar,
    BtnCopy
  },
  computed: {
    validatorId () { return 'validator-' + this.validator.id },
    avatarClass () { return 'avatar ' + this.validator.type },
    validatorCreatedAt () { return dateUnixShort(this.validator.created_at) }
  },
  methods: {
    viewValidator () {
      this.$router.push('/validators/' + this.validator.id)
    }
  },
  props: ['validator', 'voting-power', 'created-at']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.table-tr-validator
  font-size 0.5rem
  td
    border 1px solid bc
    padding 0.25*x

  td.avatar
    width 2.5*x + 1px
    text-align center
    color #fff
    .avatar
      height x*2

  td.id
    width 25%
    cursor pointer
    &:hover
      background lighten(mc,95%)
      color link

    .avatar
      display inline-block
      display none
      margin-right 0.125*x

      img, i.fa
        width 32px
        height 32px
        vertical-align middle
        background ac

      i.fa
        line-height 32px
        text-align center
        color lighten(mc,75%)

  td.public-key
    width 5%
    text-align center

  td.info
    padding 0.25*x 0.5*x
    color dim
    font-style italic

  td.created
    padding 0.25*x 0.5*x
    color dim
    width 20%
    mono()

  td.voting-power
    width x*4
    text-align center
</style>
