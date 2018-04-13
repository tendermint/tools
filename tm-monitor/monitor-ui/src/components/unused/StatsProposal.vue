<template lang="pug">
key-values
  key-value
    div(slot="key") Proposer
    div(slot="value")
      router-link(
        :to="{ path: 'delegators', params: { delegator: proposal.validatorId }}")
        | {{ proposal.validatorId }}
  key-value
    div(slot="key") Created
    div(slot="value") {{ proposalCreatedAt }}
  key-value
    div(slot="key") Type
    div(slot="value") {{ proposal.type }}
      div(v-if="proposal.online_at ===''") N/A
      div(v-else) {{ proposalOnlineAt }}
</template>

<script>
import dateUnixAgo from '../scripts/dateUnixAgo'
import KeyValues from './NiKeyValues'
import KeyValue from './NiKeyValue'
export default {
  name: 'stats-proposal',
  props: ['proposal'],
  components: {
    KeyValues,
    KeyValue
  },
  computed: {
    proposalCreatedAt () { return dateUnixAgo(this.proposal.createdAt) }
  }
}
</script>
