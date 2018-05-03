<template lang="pug">
page(title='Search')
  part(title='Search')
    form-struct(:submit="onSubmit")
      form-group
        // label(for='search-input') Input
        field#search-input(
          disabled
          theme="cosmos"
          type="text"
          placeholder="(search is coming soon)"
          required
          v-model="query"
          pattern=".{1,60}"
          title="1 to 60 characters")
      div(slot="footer")
        div
        btn(theme="cosmos" icon="search" value="Go")

  part(title='Sample Queries')
    list-item(dt='Address' dd='3HNSiAq7wFDaPsYDcUxNSRMD78qVcYKicw' @click.native="fillField('3HNSiAq7wFDaPsYDcUxNSRMD78qVcYKicw')")
    list-item(dt='Transaction' dd='878797d16182f6ffa5e98cd88576ab973b8d6f9dfe8ff441bd3521a7d8559800' @click.native="fillField('878797d16182f6ffa5e98cd88576ab973b8d6f9dfe8ff441bd3521a7d8559800')")
    list-item(dt='Block #' dd='112436' @click.native="fillField('112436')")
</template>

<script>
import { mapGetters } from 'vuex'
import Btn from '@nylira/vue-button'
import Field from '@nylira/vue-field'
import FormGroup from './NiFormGroup'
import FormStruct from './NiFormStruct'
import ListItem from './NiListItem'
import Page from './NiPage'
import Part from './NiPart'
export default {
  name: 'page-search',
  components: {
    Btn,
    Field,
    FormGroup,
    FormStruct,
    ListItem,
    Page,
    Part
  },
  computed: {
    ...mapGetters(['blockchain'])
  },
  data: () => ({
    query: ''
  }),
  methods: {
    fillField (value) {
      this.query = value
    },
    onSubmit () {
      alert('Search functionality is coming soon. Please have patience!')
    }
  }
}
</script>
