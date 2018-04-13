<template lang="pug">
page(title='New Validator')
  part
    form-struct(v-on:submit.prevent="propose")
      div(slot="title") #[i.fa.fa-search] New Validator
      div(slot="subtitle") Create a proposal for a new Validator.

      form-group
        label(for="new-validator-title") Validator ID
        field#new-validator-title(theme="cosmos", type="text", v-model="newValidator.id"
          placeholder="Validator Name", required="true",
          pattern=".{3,60}", title="3 to 60 characters")

      form-group
        label(for="new-validator-pub-key") Public Key
        field#new-validator-pub-key(theme="cosmos", type="textarea", v-model="newValidator.pub_key",
          placeholder="Paste public key here", required="true")

      form-group
        label(for="new-validator-info") Notes
        field#new-validator-info(theme="cosmos", type="textarea", v-model="newValidator.info",
          placeholder="Write notes (Optional)")

      div(slot="footer")
        btn(theme="cosmos", icon="refresh", value="Reset", @click.native="resetForm")
        btn(theme="cosmos", value="Propose")
</template>

<script>
import { mapGetters } from 'vuex'
import Btn from '@nylira/vue-button'
import FormStruct from './FormStruct'
import FormGroup from './FormGroup'
import Field from '@nylira/vue-field'
import Page from './NiPage'
import Part from './NiPart'
export default {
  name: 'page-validators-new',
  components: {
    Field,
    Btn,
    FormGroup,
    FormStruct,
    Page,
    Part
  },
  computed: {
    ...mapGetters(['newValidator'])
  },
  data: () => ({
    debug: false,
    validatorTypes: [
      { id: 'user' },
      { id: 'machine' }
    ]
  }),
  methods: {
    propose () {
      this.$store.commit('ADD_VALIDATOR', this.newValidator)
      this.$store.commit('RESET_VALIDATORS_NEW')
      this.$router.push('/validators/')
    },
    resetForm () {
      this.$store.commit('RESET_VALIDATORS_NEW')
    }
  },
  mounted () {
    this.resetForm()
  }
}
</script>
