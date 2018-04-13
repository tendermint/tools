<template lang="pug">
.panel-votes
  chart-votes(:votes="votes" size="lg" cutout="90")

  // key-values
    key-value
      div(slot="key") Yes
      div(slot="value") 
        | {{ votes.yes }}
        btn(
          theme="cosmos",
          v-if="myVote.value !== 'yes'",
          icon="check"
          @click.native="voteYes")

  // key-value
    div(slot="key") No
    div(slot="value") 
      | {{ votes.no }}
      btn(
        theme="cosmos",
        v-if="myVote.value !== 'no'",
        icon="times"
        @click.native="voteNo")

  // key-value
    div(slot="key") Reject
    div(slot="value") 
      | {{ votes.reject }}
      btn(
        theme="cosmos",
        v-if="myVote.value !== 'reject'",
        icon="times"
        @click.native="voteReject")

  // key-value
    div(slot="key") Abstain
    div(slot="value") 
      | {{ votes.abstain }}
      btn(
        theme="cosmos",
        v-if="myVote.value !== 'abstain'",
        icon="ban"
        @click.native="voteAbstain")

  // card-vote-validator(v-for="vote in yesVotes", :key="vote.id", :vote="vote")
  // card-vote-validator(v-for="vote in noVotes", :key="vote.id", :vote="vote")
  // card-vote-validator(v-for="vote in nonVotes", :key="vote.id", :vote="vote")
</template>

<script>
import CardVoteValidator from './CardVoteValidator'
import ChartVotes from './ChartVotes'
import KeyValue from './NiKeyValue'
import KeyValues from './NiKeyValues'
import VoteProgressBar from './VoteProgressBar'
export default {
  name: 'panel-votes',
  props: ['votes'],
  components: {
    CardVoteValidator,
    ChartVotes,
    KeyValue,
    KeyValues,
    VoteProgressBar
  },
  computed: {
    myVote () {
      if (this.votes && this.votes.length > 0) {
        return this.votes.find(v => v.validator_id === this.me.id)
      } else {
        return {
          validator_id: 0,
          value: 0,
          created_at: ''
        }
      }
    }
  },
  methods: {
    voteYes () {
      this.myVote.value = 'yes'
      this.$store.commit('UPDATE_VOTE', this.myVote)
    },
    voteNo () {
      this.myVote.value = 'no'
      this.$store.commit('UPDATE_VOTE', this.myVote)
    },
    voteReject () {
      this.myVote.value = 'reject'
      this.$store.commit('UPDATE_VOTE', this.myVote)
    },
    voteAbstain () {
      this.myVote.value = 'abstain'
      this.$store.commit('UPDATE_VOTE', this.myVote)
    }
  }
}
</script>
