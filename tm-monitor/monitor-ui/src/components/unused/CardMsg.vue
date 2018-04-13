<template lang="pug">
.card-msg(:id="msgId")
  router-link.avatar(to="{ path: 'entities', params: { entity: message.entity_id }}")
    card-avatar(:entity-id="message.entity_id")
  .content
    .meta
      a.title(to="{ path: 'entities', params: { entity: message.entity_id }}")
        | {{ message.entity_id }}
      .date(:title="msgCreatedAt")
        | commented {{ msgCreatedAtAgo }}
      a.thread(v-if="indexed === true",
        to="{ path: 'proposals', params: { proposal: proposal.id }}")
        | on #[span.strong {{ proposal.title }}]
      .right
        template(v-if="indexed === true")
          // TODO: add #message- + 'message.id
          router-link.permalink(@click="smoothScroll",
          to="{ path: 'proposals', params: { proposal: message.proposal_id }}")
            | #[i.fa.fa-link] Permalink
        template(v-else)
          a.permalink(@click="smoothScroll", :href="msgIdHash")
            | #[i.fa.fa-link] Permalink
        
    .body(:id="editId")
      // TODO: Add debounce
      textarea(v-model="message.body", v-show="editable")
      .markdown(v-html="message.body | marked", v-show="!editable")
    .menu
      a.btn.btn-default.btn-small(@click="rmMessage(message)")
        | #[i.fa.fa-trash] Delete
      template(v-if="indexed === true")
        a.btn.btn-primary.btn-small(@click="edit", v-show="!editable")
          | #[i.fa.fa-pencil] Edit
        a.btn.btn-primary.btn-small(@click="save", v-show="editable")
          | #[i.fa.fa-save] Save
        // TODO: add #messages-panel
        router-link.btn.btn-primary.btn-small(
          to="{ path: 'proposals', params: { proposal: message.proposal_id }}" )
          | #[i.fa.fa-eye] View Thread
      template(v-else)
        a.btn.btn-primary.btn-small(@click="edit", v-show="!editable")
          | #[i.fa.fa-pencil] Edit
        a.btn.btn-primary.btn-small(@click="save", v-show="editable")
          | #[i.fa.fa-save] Save
</template>

<script>
import { mapGetters } from 'vuex'
import smoothScroll from 'smooth-scroll'
import CardAvatar from './CardAvatar'
import dateUnix from '../scripts/dateUnix'
import dateUnixAgo from '../scripts/dateUnixAgo'
export default {
  name: 'card-msg',
  components: {
    CardAvatar
  },
  computed: {
    ...mapGetters(['proposals']),
    msgId () {
      return 'message-' + this.message.id
    },
    msgIdHash () {
      return '#' + this.msgId
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
    }
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
  mounted () {
    this.smoothScrollOnLoad()
  },
  props: ['message', 'indexed']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.card-msg
  margin 0 0 1em
  display flex
  margin 0 0 1em
  align-items flex-start

  a.avatar
    display block
    border 1px solid bc
    padding x*0.125
    background #fff
    margin-right 1em
    .avatar
      display block
      width x*2.625
      height x*2.625

  .meta
    background lighten(dim, 95%)
    padding x*0.75 x x*0.75 - 1px
    border-bottom 1px solid bc
    display flex
    vertical-align text-top

    a.title
      font-weight 600
      color txt
      margin-right x*0.3
      &:hover
        color link

    .date
      color dim
      margin-right x*0.3

    a.thread
      color dim
      span.strong
        font-weight 600
        color txt
      &:hover
        span.strong
          color link

    .right
      flex 1
      a.permalink
        display block
        text-align right
        color dim
        &:hover
          color link

  .content
    flex 1
    border 1px solid bc

    .body
      textarea, .markdown
        max-width 50em
        margin x
      textarea
        display block
        width 100%

    .menu
      border-top 1px dotted bc
      padding x*0.5
      display flex
      justify-content flex-end
      .btn
        margin-left 0.5*x
</style>
