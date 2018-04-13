<template lang="pug">
#app-main: slot
</template>

<script>
export default {
  name: 'app-main'
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

#app-main
  flex 1

  header.section-header
    display flex

    .text
      flex 1
      line-height x*2.25

      h2
        font-size 1.5em
        line-height x*2.25

        display flex

        .segment, .segment a
          display flex
          align-items center
          color txt
          font-weight normal

        .segment a
          cursor pointer
          &:hover
            color link

        .segment
          &:before
            display block
            width 1.5*x

            content '\f105'
            font-family FontAwesome
            text-align center
            color subtle
            font-weight normal
            font-size 1.333*x

          &:first-of-type
            font-weight bold
            a
              font-weight bold
            &:before
              margin-right 0.5*x
              font-weight 200
            &.segment-blockchain:before
              content '\f009'
            &.segment-groups:before
              content '\f19c'
            &.segment-entities:before
              content '\f007'
            &.segment-proposals:before
              content '\f0ca'
            &.segment-proposal:before
              content '\f10d'
            &.segment-permissions:before
              content '\f023'

      p
        color dim

    nav
      height x * 2
      margin-top 2px

      > *
        margin-left 0.5em

      input
        width 18em

      p
        line-height x*2
</style>
