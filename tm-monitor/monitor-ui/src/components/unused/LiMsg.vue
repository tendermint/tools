<template lang="pug">
router-link.card-simple.li-msg(:to="{ name: 'proposal', params: { proposal: proposal.id }}")
  .icon: i.fa.fa-comment
  .text
    .title {{ proposal.title }}
    .date(:title="msgCreatedAt")
      | {{ msgCreatedAtAgo }}
</template>

<script>
import { mapGetters } from 'vuex'
import smoothScroll from 'smooth-scroll'
import dateUnix from '../scripts/dateUnix'
import dateUnixAgo from '../scripts/dateUnixAgo'
export default {
  name: 'li-msg',
  computed: {
    ...mapGetters(['proposals']),
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
