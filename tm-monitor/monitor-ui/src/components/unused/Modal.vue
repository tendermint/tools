<template>
  <div :class="cssClass">
    <div class="ni-modal-container">
      <header class="ni-modal-header">
        <div class="ni-modal-icon" v-if="icon"><i :class="iconCssClass"></i></div>
        <div class="ni-modal-title"><slot name="title"></slot></div>
      </header>
      <main class="ni-modal-main">
        <slot></slot>
      </main>
      <footer class="ni-modal-footer">
        <slot name="footer"></slot>
      </footer>
    </div>
  </div>
</template>

<script>
import Ps from 'perfect-scrollbar'
export default {
  computed: {
    iconCssClass () {
      let value = 'fa fa-'
      if (this.icon) {
        value += this.icon
        if (this.icon === 'refresh') {
          value += ' fa-spin'
        }
      }
      return value
    },
    cssClass () {
      let value = 'ni-modal'
      if (this.size === 'fullscreen' || this.size === 'fs') {
        value += ' ni-modal-fullscreen'
      }
      return value
    }
  },
  methods: {
    close () {
      this.$destroy()
    }
  },
  mounted () {
    Ps.initialize(document.querySelector('.ni-modal-main'))
  },
  props: ['size', 'icon']
}
</script>

<style lang="stylus">
@import '../styles/variables.styl'

.ni-modal
  position fixed
  top 0
  left 0
  z-index 1000

  width 100vw
  height 100vh
  background hsla(0,0,0,0.5)

  display flex
  justify-content center
  align-items center
  backdrop-filter blur(0.5em)

  &.ni-modal-fullscreen
    .ni-modal-container
      max-height none
      max-width 40rem + 6rem
      height 100vh
      width 100vw
      display flex
      flex-flow column nowrap
    .ni-modal-main
      flex 1
      overflow-y scroll

.ni-modal-container
  background c-app-fg
  box-shadow hsla(0,0,0,0.25) 0 0.25rem 1rem

  max-width 30rem
  max-height 40rem
  
.ni-modal-header, .ni-modal-main
  border-bottom 1px solid bc

.ni-modal-header
  display flex
  align-items center
  height 3rem + 0.0625rem

.ni-modal-main, .ni-modal-footer
  padding 1rem

.ni-modal-icon
  padding-left 1rem
  padding-right 0.5rem
  height 3rem
  display flex
  align-items center
  justify-content center

.ni-modal-icon + .ni-modal-title
  padding-left 0

.ni-modal-title
  flex 1
  font-weight 500
  padding 0 1rem

.ni-modal-close
  cursor pointer
  &:hover
    color link

.ni-modal-main
  display flex
  flex-flow column
  justify-content center

  .ps-scrollbar-y-rail
    display none

.ni-modal-main p
  margin-bottom 1rem

.ni-modal-main > .ni-article-body
  margin -1rem

.ni-modal-footer > div
  display flex
  justify-content space-between
</style>
