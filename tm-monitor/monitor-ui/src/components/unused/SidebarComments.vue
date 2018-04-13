<template lang="pug">
.sidebar
  .sidebar-container
    .sidebar-section-header
      .title {{ commentsTitle }}
      .menu
        a(@click="toggleCommentsEditable")
          template(v-if="commentsEditable") Cancel
          template(v-else) Edit
    .sidebar-section-form
      form#message-panel-new-message(v-on:submit.prevent="publish")
        textarea.textarea-small(
          placeholder="Write your comment", required,
          v-model="newMessageBody", minlength="3", maxlength="10000")
        button.btn.btn-primary.btn-small(type="submit")
          | #[i.fa.fa-comment] Post
    .sidebar-section-content
      card-msg-small(
        v-for="message in filteredMessages",
        :key="message.id",
        :message="message",
        :can-be-edited="commentsEditable")
</template>

<script>
import $ from 'jquery'
import moment from 'moment'
import { mapGetters } from 'vuex'
import pluralizeWithCount from '../scripts/pluralizeWithCount'
import CardAvatar from './CardAvatar'
import CardMsgSmall from './CardMsgSmall'
import { orderBy } from 'lodash'
export default {
  name: 'sidebar-comments',
  components: {
    CardAvatar,
    CardMsgSmall
  },
  computed: {
    ...mapGetters(['me', 'messages']),
    filteredMessages () {
      if (this.messages) {
        return orderBy(
          this.messages.filter(m => m.proposal_id === this.proposalId),
          ['created_at'], ['desc'])
      } else {
        return []
      }
    },
    commentsTitle () {
      return pluralizeWithCount('Comment', this.filteredMessages.length)
    }
  },
  data () {
    return {
      newMessageBody: '',
      commentsEditable: false
    }
  },
  methods: {
    publish () {
      let newMsg = {
        proposal_id: this.proposalId,
        entity_id: this.me.id,
        body: this.newMessageBody,
        visible: true,
        created_at: moment().valueOf()
      }
      this.$store.commit('ADD_MSG', newMsg)
      this.newMessageBody = ''
    },
    disableScrollPropagation () {
      $('.sidebar-section-content').on('DOMMouseScroll mousewheel', function (ev) {
        var $this = $(this)
        var scrollTop = this.scrollTop
        var scrollHeight = this.scrollHeight
        var height = $this.height()
        var delta = (ev.type === 'DOMMouseScroll'
            ? ev.originalEvent.detail * -40
            : ev.originalEvent.wheelDelta)
        var up = delta > 0

        var prevent = function () {
          ev.stopPropagation()
          ev.preventDefault()
          ev.returnValue = false
          return false
        }

        if (!up && -delta > scrollHeight - height - scrollTop) {
          // Scrolling down, but this will take us past the bottom.
          $this.scrollTop(scrollHeight)
          return prevent()
        } else if (up && delta > scrollTop) {
          // Scrolling up, but this will take us past the top.
          $this.scrollTop(0)
          return prevent()
        }
      })
    },
    toggleCommentsEditable () {
      this.commentsEditable = !this.commentsEditable
    },
    watchEnterKey () {
      $('textarea.textarea-small').keydown(function (event) {
        if (event.keyCode === 13) {
          console.log('publishing!')
          this.publish()
          return false
        }
      })
    }
  },
  props: ['proposal-id'],
  mounted () {
    this.disableScrollPropagation()
    this.watchEnterKey()
  },
  beforeDestroy () {
    // disable comment editing on leaving the page
    this.commentsEditable = false
  }
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.sidebar-container
  display flex
  flex-flow column nowrap
  margin-top 1rem

.sidebar-section-header
  border-bottom 1px solid bc-dim

  line-height 2*x
  color txt
  font-size 0.75*x
  font-weight 600

  display flex
  justify-content space-between

  .menu
    a
      cursor pointer
      &:hover
        text-decoration underline

.sidebar-section-form
  padding 0.5*x 0
  border-bottom 1px solid bc

  form
    display flex
    flex-flow column nowrap
    textarea.textarea-small
      flex 1
      background transparent
      border-color bc-dim
      height 6rem

.sidebar-section-content
  flex 1
  overflow-y scroll
  background bc2
  padding 0.25*x
</style>
