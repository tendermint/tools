<template lang="pug">
page(:title='proposal.title')
  div
    btn(type="link" theme="cosmos" to="/proposals" exact value="All")
    btn(theme="cosmos" @click.native="deleteProposal(proposal)" icon="trash")
    btn(theme="cosmos" icon="check" value="Approved" v-if="proposal.flags && proposal.flags.passed")

  part
    stats-proposal(:proposal="proposal")

  part(title='Votes')
    panel-votes(:votes="proposal.votes")

  part(v-if="proposal.type === 'text'" title='Text')
    proposal-text(:content="proposal.data.text")

  part(title='Comments')
    sidebar-comments(:proposal-id="proposal.id")
</template>

<script>
import { mapGetters } from 'vuex'
import dateUnix from '../scripts/dateUnix'
import Btn from '@nylira/vue-button'
import BtnCopy from './BtnCopy'
import CardMsg from './CardMsg'
import Page from './NiPage'
import Part from './NiPart'
import ProposalText from './ProposalText'
import PanelVotes from './PanelVotes'
import StatsProposal from './StatsProposal'
import SidebarComments from './SidebarComments'
export default {
  name: 'page-proposal',
  components: {
    Btn,
    BtnCopy,
    CardMsg,
    Page,
    Part,
    ProposalText,
    PanelVotes,
    SidebarComments,
    StatsProposal
  },
  computed: {
    ...mapGetters(['proposals']),
    proposal () {
      if (this.proposals && this.proposals.length > 0) {
        return this.proposals.find(p => p.id === this.$route.params.proposal)
      } else {
        return this.emptyProposal
      }
    },
    proposalActiveAt () {
      return dateUnix(this.proposal.active_at)
    }
  },
  created () {
    this.refreshProposal()
    this.$watch('$route.params', this.refreshProposal)
  },
  data: () => ({
    emptyProposal: {
      id: '',
      active_at: '',
      created_at: '',
      entity_id: '',
      title: 'Loading',
      type: '',
      flags: {
        passed: false
      },
      data: {
        body: '',
        old_members: '',
        new_members: ''
      },
      vote_id: 0
    }
  }),
  methods: {
    refreshProposal () {
      if (this.rawProposal) {
        this.proposal = this.rawProposal
      } else {
        this.proposal = this.emptyProposal
      }
    },
    deleteProposal (proposal) {
      this.$router.push('/proposals')
      this.$store.commit('RM_PROPOSAL', proposal)
      this.$store.commit('RM_VOTES_BY_PROPOSAL', proposal.vote_id)
      this.$store.commit('RM_MSGS_BY_PROPOSAL', proposal.id)
    }
  }
}
</script>
