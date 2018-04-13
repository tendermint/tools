<template lang="pug">
transition(name="msg-small"): .card-msg-small(:id="msgId")
  router-link.avatar(
    v-if="message.entity_id",
    :to="{ name: 'entity', params: { entity: message.entity_id }}")
    card-avatar(:entity-id="message.entity_id", :is-small="true")
  .content
    .meta
      router-link.title(
        v-if="message.entity_id",
        :to="{ name: 'entity', params: { entity: message.entity_id }}")
        | {{ message.entity_id }}
      span.date(:title="msgCreatedAt")
        | {{ msgCreatedAtAgo }}
        
    .body(:id="editId")
      // TODO: debounce
      textarea.textarea-small(v-model="message.body", v-show="editable")
      .body(v-html="message.body", v-show="!editable")
    .menu(v-if="canBeEdited")
      a(@click="rmMessage(message)", title="Delete"): i.fa.fa-trash
      a(@click="edit", v-show="!editable", title="Edit"): i.fa.fa-pencil
      a(@click="save", v-show="editable", title="Save"): i.fa.fa-save
</template>

<script>
import { mapGetters } from 'vuex'
import smoothScroll from 'smooth-scroll'
import dateUnix from '../scripts/dateUnix'
import dateUnixAgo from '../scripts/dateUnixAgo'
import CardAvatar from './CardAvatar'
export default {
  name: 'card-msg-small',
  components: {
    CardAvatar
  },
  computed: {
    msgId () {
      return 'message-' + this.message.id
    },
    msgCreatedAt () {
      return dateUnix(this.message.created_at)
    },
    msgCreatedAtAgo () {
      return dateUnixAgo(this.message.created_at)
    },
    proposal () {
      if (this.proposals && this.proposals.length > 0) {
        return this.proposals.find(p => p.id === this.message.proposal_id)
      } else {
        return { id: 0, title: 'Loading' }
      }
    },
    ...mapGetters(['proposals'])
  },
  data () {
    return {
      editId: 'message-edit-' + this.message.id,
      editable: false
    }
  },
  methods: {
    rmMessage (msg) {
      this.$store.commit('RM_MSG', msg)
    },
    edit () {
      this.editable = true
    },
    save () {
      this.editable = false
    },
    smoothScroll (event) {
      event.preventDefault()
      smoothScroll.animateScroll(
        event.target.hash,
        event.target,
        { updateURL: false, offset: 65 }
      )
    },
    smoothScrollOnLoad () {
      if (window.location.hash) {
        // console.log('theres a hash!')
        let hash = smoothScroll.escapeCharacters(window.location.hash)
        let toggle = document.querySelector('a[href*="' + hash + '"]')
        let options = { offset: 65 }
        smoothScroll.animateScroll(window.location.hash, toggle, options)
        history.replaceState('', document.title, window.location.pathname)
      }
    }
  },
  props: ['message', 'can-be-edited'],
  mounted () {
    this.smoothScrollOnLoad()
  }
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.card-msg-small
  display flex
  font-size 0.75*x
  padding 0.375*x

  position relative
  background #fff
  margin 0 0 1px

  &:last-of-type
    border-bottom none

  a.avatar
    display block
    margin-right 0.375*x
    .avatar
      display block

  .meta
    a.title
      font-weight 600
      color txt
      margin-right x*0.3
      &:hover
        color link

    .date
      color dim
      margin-right x*0.3

  .content
    flex 1

    .body
      textarea, .body
        max-width 50em
      textarea
        display block
        width 100%
        margin 0 0 1.5*x

    .menu
      position absolute
      bottom 0
      right 0
      display flex
      a
        display block
        color dim
        cursor pointer
        background bc2

        line-height 1.5*x
        width 1.5*x
        text-align center
        
        &:hover
          color #fff
          background link
</style>
