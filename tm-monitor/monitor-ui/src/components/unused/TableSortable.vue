<template lang="pug">
table.table-sortable
  thead
    tr(is="table-tr-header", :sort="sort")
  tbody
    tr(is="table-tr-validator", v-for="validator in validators", :validator="validator")
</template>

<script>
import TableTrHeader from './TableTrHeader'
import TableTrValidator from './TableTrValidator'
export default {
  name: 'table-sortable',
  components: {
    TableTrHeader,
    TableTrValidator
  },
  props: ['sort', 'validators']
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.table-sortable
  width 100%
  border-collapse collapse
</style>
