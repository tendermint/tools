<template lang="pug">
#app
  app-header
  #app-content
    #app-main: router-view
    app-footer
  notifications(:notifications='notifications' theme='cosmos')
</template>

<script>
import requestInterval from 'request-interval'
import { mapGetters } from 'vuex'
import AppFooter from './components/AppFooter'
import AppHeader from './components/AppHeader'
import Notifications from '@nylira/vue-notifications'
import store from './store/index'
export default {
  name: 'app',
  components: {
    AppHeader,
    AppFooter,
    Notifications
  },
  computed: {
    ...mapGetters(['notifications'])
  },
  mounted () {
    requestInterval(1000, () => this.$store.commit('getStatus'))
    requestInterval(1000, () => this.$store.commit('getValidators'))
  },
  store
}
</script>

<style lang="stylus" src="./styles/app.styl"></style>
