<template lang="pug">
page(:title='delegator.id')
  tool-bar
    router-link(to="/delegators" exact): i.material-icons arrow_back
    anchor-copy(:value="delegator.pub_key" icon="content_copy")

  part(title='Profile')
    list-item(dt="Vote Power" dd="12,243 ATOM" to="/vote-power")
    list-item(dt="Vote History" dd="35 Votes" to="/votes")
    list-item(dt="Proposals" dd="5" to="/proposals")
    list-item(dt="Slashes" dd="2" to="/slashes")

  part(title='Staking')
    list-item(dt="Validators" dd="3" to="/validators")
    list-item(dt="Earning Rate" dd="520.1 ATOM/Day")
    list-item(dt="Total Earnings" dd="2,428 ATOM")

  // part(:title='votesTitle')
    li-vote(
      v-if="filteredVotes"
      v-for="vote in filteredVotes"
      :key="vote.id"
      :vote="vote")

  // part(:title='msgsTitle')
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
// import StatsDelegator from './StatsDelegator'
// import LiMessage from './LiMsg'
// import LiVote from './LiVote'
import ListItem from './NiListItem'
import ToolBar from './NiToolBar'
import Page from './NiPage'
import Part from './NiPart'
import AnchorCopy from './AnchorCopy'
export default {
  name: 'page-delegator-index',
  components: {
    AnchorCopy,
    // StatsDelegator,
    // LiMessage,
    // LiVote,
    ListItem,
    Page,
    Part,
    ToolBar
  },
  computed: {
    ...mapGetters(['current', 'delegators', 'messages', 'proposals', 'votes']),
    delegator () {
      if (this.delegators) {
        return this.delegators.find(d => d.id === this.$route.params.delegator)
      } else {
        return this.emptyDelegator
      }
    },
    filteredMsgs () {
      if (this.messages && this.messages.length > 0) {
        return orderBy(
          this.messages.filter(m => m.delegator_id === this.current.delegatorId),
          ['created_at'], ['desc'])
      } else {
        return []
      }
    },
    filteredProposals () {
      if (this.proposals && this.proposals.length > 0) {
        return orderBy(
          this.proposals.filter(p => p.delegator_id === this.current.delegatorId),
          ['created_at'], ['desc'])
      } else {
        return []
      }
    },
    filteredVotes () {
      if (this.votes && this.votes.length > 0) {
        return orderBy(
          this.votes.filter(v => v.delegator_id === this.current.delegatorId),
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

<style lang="stylus">
@require '../styles/variables.styl'
</style>
