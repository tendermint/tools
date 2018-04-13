<template lang="pug">
router-link.card-simple.li-vote(:to="{ name: 'proposal', params: { proposal: proposal.id }}")
  .icon.icon-no(v-if="vote.value === -1") #[i.fa.fa-times]
  .icon.icon-abstained(v-if="vote.value === 0") #[i.fa.fa-question]
  .icon.icon-yes(v-if="vote.value === 1") #[i.fa.fa-check]
  .text
    .title {{ proposal.title }}
    .date(:title="voteUpdatedAt")
      | {{ voteUpdatedAtAgo }}
</template>

<script>
import { mapGetters } from 'vuex'
import dateUnix from '../scripts/dateUnix'
import dateUnixAgo from '../scripts/dateUnixAgo'
export default {
  name: 'li-vote',
  computed: {
    ...mapGetters(['proposals']),
    voteUpdatedAt () {
      return dateUnix(this.vote.updated_at)
    },
    voteUpdatedAtAgo () {
      return dateUnixAgo(this.vote.updated_at)
    },
    proposal () {
      if (this.proposals.length > 0) {
        return this.proposals.find(p => p.vote_id === this.vote.proposal_id)
      } else {
        return { id: 0, title: 'undefined proposal' }
      }
    }
  },
  props: ['vote']
}
</script>
