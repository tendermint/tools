<template lang="pug">
page(title='Settings')
  sign-in
</template>

<script>
import Btn from '@nylira/vue-button'
import Page from './NiPage'
import Part from './NiPart'
import SignIn from './SignIn'
export default {
  name: 'page-accounts-settings',
  components: {
    Btn,
    Page,
    Part,
    SignIn
  }
}
</script>
