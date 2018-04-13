<template lang="pug">
form-struct(:submit='signProposal')
  form-group(:error='$v.fields.seed.$error')
    label(for='field-seed') Private Seed
    field-group
      field#field-seed(theme='cosmos' type='text' v-model='fields.seed' placeholder='Enter your private seed')
      .ni-field-addon #[i.material-icons lock]
    form-msg(name='Seed' type='required' v-if='!$v.fields.seed.required')
  form-group(:error='$v.fields.hex.$error')
    label(for='field-hex') Hex String
    field-group
      field#field-hex(type='text' theme='cosmos' v-model='fields.hex' placeholder='Enter the hex string')
    form-msg(name='Hex String' type='required' v-if='!$v.fields.hex.required')
  div(slot='footer')
    div
    btn(theme='cosmos' type='submit' icon='edit' value='Sign Proposal')
</template>

<script>
import { minLength, maxLength, required } from 'vuelidate/lib/validators'
import Btn from '@nylira/vue-button'
import FormStruct from './FormStruct'
import FormGroup from './FormGroup'
import FormMsg from '@nylira/vue-form-msg'
import FieldGroup from './FieldGroup'
import Field from '@nylira/vue-field'
export default {
  name: 'wallet',
  components: {
    Btn,
    Field,
    FieldGroup,
    FormMsg,
    FormGroup,
    FormStruct
  },
  data: () => ({
    fields: {
      seed: '',
      hex: ''
    }
  }),
  methods: {
    signProposal () {
      this.$v.$touch()
      if (!this.$v.$error) {
        this.$store.commit('notify', { title: 'Proposal Signed',
          body: 'You have successfuly signed $TODO Proposal.' })
      }
    },
    resetFields () {
      this.fields = {
        seed: '',
        hex: ''
      }
    }
  },
  validations: () => ({
    fields: {
      seed: {
        required,
        minLength (x) { return minLength(8) },
        maxLength (x) { return maxLength(255) }
      },
      hex: {
        required,
        minLength (x) { return minLength(8) },
        maxLength (x) { return maxLength(255) }
      }
    }
  })
}
</script>
