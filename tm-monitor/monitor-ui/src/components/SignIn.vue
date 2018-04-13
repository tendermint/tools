<template lang="pug">
form-struct(:submit='signIn')
  form-group(:error='$v.fields.seed.$error')
    label(for='field-seed') Cosmos Seed
    field-group
      field#field-seed(theme='cosmos' type='text' v-model='fields.seed' placeholder='Enter your 12-word seed')
      .ni-field-addon #[i.material-icons lock]
    form-msg(name='Seed' type='required' v-if='!$v.fields.seed.required')
  div(slot='footer')
    div
    btn(theme='cosmos' type='submit' value='Sign In')
</template>

<script>
import { minLength, maxLength, required } from 'vuelidate/lib/validators'
import Btn from '@nylira/vue-button'
import FormStruct from './NiFormStruct'
import FormGroup from './NiFormGroup'
import FormMsg from '@nylira/vue-form-msg'
import FieldGroup from './NiFieldGroup'
import Field from '@nylira/vue-field'
export default {
  name: 'sign-in',
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
      seed: ''
    }
  }),
  methods: {
    signIn () {
      this.$v.$touch()
      if (!this.$v.$error) {
        this.$store.commit('notifySignIn')
      }
    },
    resetFields () {
      this.fields = {
        seed: ''
      }
    }
  },
  validations: () => ({
    fields: {
      seed: {
        required,
        minLength (x) { return minLength(8) },
        maxLength (x) { return maxLength(255) }
      }
    }
  })
}
</script>
