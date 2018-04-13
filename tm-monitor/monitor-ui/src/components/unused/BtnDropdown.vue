<template lang="pug">
.btn-dropdown-wrapper(@click="toggleDropdown")
  .btn.btn-primary
    i.fa.fa-plus
    i.fa.fa-caret-down
  .btn-dropdown-menu(v-if="showDropdown", @click="toggleDropdown")
    ul
      li(v-if="!this.groupId"): a(@click="createGroupProposal")
        | #[i.fa.fa-plus] New Group Proposal
      li: a(@click="createUpdateGroupProposal")
        | #[i.fa.fa-plus] Update Group Proposal
      li: a(@click="createTextProposal")
        | #[i.fa.fa-plus] Text Proposal
      li
        a.cancel #[i.fa.fa-times] Cancel
</template>

<script>
export default {
  name: 'btn-dropdown',
  data () {
    return {
      showDropdown: false
    }
  },
  methods: {
    toggleDropdown: function (event) {
      event.stopPropagation()
      this.showDropdown = !this.showDropdown
    },
    createGroupProposal () {
      this.$router.push('/proposals/new/group')
    },
    createUpdateGroupProposal () {
      if (this.groupId) {
        this.$store.commit('SET_PROPOSAL_UPDATE_GROUP_GROUP_ID', this.groupId)
      }
      this.$router.push('/proposals/new/update-group')
    },
    createTextProposal () {
      if (this.groupId) {
        this.$store.commit('SET_PROPOSAL_TEXT_GROUP_ID', this.groupId)
      }
      this.$router.push('/proposals/new/text')
    }
  },
  beforeDestroy () {
    this.showDropdown = false
  },
  props: ['group-id']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.btn-dropdown-wrapper
  display flex
  position relative

  .btn
    padding-right x*3.5

  i.fa.fa-caret-down
    position absolute
    top 0
    right 0
    display block
    padding 0 x*0.5
    margin-left -2px

    width x*2
    height x*2
    line-height x*2

    border-left 1px solid darken(mc,25%)
    box-shadow inset 0 0 1px lighten(mc,50%)

.btn-dropdown-menu
  height 100vh
  width 100vw
  position fixed
  top 0
  left 0
  background hsla(0,0,0,0.25)
  z-index 101

  display flex
  justify-content center
  align-items center

  ul
    width 30*x
    li
      margin 0 0 0.5*x
      &:last-of-type
        margin-bottom 0
  a
    display block
    padding 0 2*x
    background #fff

    font-size x*2
    line-height x*6
    text-align center

    text-shadow none
    color txt
    text-align left

    box-shadow hsla(0,0,0,0.15) 0 0.25*x 0.5*x
    cursor pointer

    &:hover
      background faint
      color link

    &.cancel
      background hsla(0,0,0,0.75)
      color #fff
      border none
      &:hover
        background hsla(0,0,0,0.85)

    i.fa
      margin-right x*0.666
</style>
