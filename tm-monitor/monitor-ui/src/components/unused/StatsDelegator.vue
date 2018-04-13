<template lang="pug">
key-values
  key-value
    div(slot="key") Online
    div(slot="value")
      div(v-if="delegator.online_at ===''") N/A
      div(v-else) {{ delegatorOnlineAt }}
  key-value
    div(slot="key") Created
    div(slot="value") {{ delegatorCreatedAt }}
  key-value
    div(slot="key") Info
    div(slot="value") {{ delegator.info }}
</template>

<script>
import dateUnixAgo from '../scripts/dateUnixAgo'
import BtnCopy from './BtnCopy'
import KeyValues from './NiKeyValues'
import KeyValue from './NiKeyValue'
export default {
  name: 'stats-delegator',
  components: {
    BtnCopy,
    KeyValues,
    KeyValue
  },
  computed: {
    delegatorOnlineAt () { return dateUnixAgo(this.delegator.online_at) },
    delegatorCreatedAt () { return dateUnixAgo(this.delegator.created_at) }
  },
  props: ['delegator']
}
</script>
