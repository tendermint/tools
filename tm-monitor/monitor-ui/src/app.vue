<template lang="pug">
#app
  app-header
  #app-content
    #app-main: router-view
    app-footer
</template>

<script>
import requestInterval from 'request-interval'
import AppFooter from './components/AppFooter'
import AppHeader from './components/AppHeader'
import store from './store/index'
export default {
  name: 'app',
  components: {
    AppHeader,
    AppFooter
  },
  mounted () {
    requestInterval(1000, () => this.$store.commit('getStatus'))
    requestInterval(1000, () => this.$store.commit('getValidators'))
  },
  store
}
</script>

<style lang="stylus" src="./styles/app.styl"></style>
