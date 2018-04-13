<template lang="pug">
page(:title='validator.name')
  div
    btn(type="link", theme="cosmos", to="/validators", exact, value="All")
    btn-copy(:text="validator.pub_key")

  part-validator(:node="validator")

  part(title='votesTitle')
    li-vote(
      v-if="filteredVotes"
      v-for="vote in filteredVotes"
      :key="vote.id"
      :vote="vote")

  part(title='msgsTitle')
    li-message(
      v-if="filteredMsgs"
      v-for="msg in filteredMsgs"
      :key="msg.id"
      :message="msg")
</template>

<script>
import { mapGetters } from 'vuex'
import { orderBy } from 'lodash'
import pluralizeWithCount from '../scripts/pluralizeWithCount'
import Btn from '@nylira/vue-button'
import BtnCopy from './BtnCopy'
import LiMessage from './LiMsg'
import LiVote from './LiVote'
import Page from './NiPage'
import Part from './NiPart'
import PartValidator from './PartValidator'
export default {
  name: 'page-validator-index',
  components: {
    Btn,
    BtnCopy,
    LiMessage,
    LiVote,
    Page,
    Part,
    PartValidator
  },
  computed: {
    ...mapGetters(['blockchain', 'current', 'messages', 'proposals', 'votes']),
    validator () {
      if (this.blockchain) {
        return this.blockchain.nodes.find(n => n.name === this.$route.params.validator)
      } else {
        return {}
      }
    },
    filteredMsgs () {
      if (this.messages && this.messages.length > 0) {
        return orderBy(
          this.messages.filter(m => m.validator_id === this.current.validatorId),
          ['created_at'], ['desc'])
      } else {
        return []
      }
    },
    filteredProposals () {
      if (this.proposals && this.proposals.length > 0) {
        return orderBy(
          this.proposals.filter(p => p.validator_id === this.current.validatorId),
          ['created_at'], ['desc'])
      } else {
        return []
      }
    },
    filteredVotes () {
      if (this.votes && this.votes.length > 0) {
        return orderBy(
          this.votes.filter(v => v.validator_id === this.current.validatorId),
          ['created_at'], ['desc'])
      } else {
        return []
      }
    },
    msgsTitle () {
      if (this.filteredMsgs && this.filteredMsgs.length > 0) {
        return pluralizeWithCount('Message', this.filteredMsgs.length)
      } else {
        return '0 Messages'
      }
    },
    votesTitle () {
      if (this.filteredVotes && this.filteredVotes.length > 0) {
        return pluralizeWithCount('Vote', this.filteredVotes.length)
      } else {
        return '0 Votes'
      }
    }
  },
  data () {
    return {
      emptyDelegator: {
        id: 'Loading',
        pub_key: '',
        info: '',
        type: 'user',
        created_at: 0,
        online_at: 0
      }
    }
  }
}
</script>
