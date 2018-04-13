<template lang="pug">
page(title='New Proposal')
  tool-bar
    router-link(to="/proposals" exact): i.material-icons arrow_back
  part(title='Choose a Type')
    list-item(title='Adjust Parameters' to="/proposals/new/adjust-parameters")
    list-item(title='Amend Constitution' to="/proposals/new/amend-constitution")
    list-item(title='Create Atoms' to="/proposals/new/create-atoms")
    list-item(title='Plain Text' to="/proposals/new/plain-text")
    list-item(title='Upgrade Code' to="/proposals/new/upgrade-code")
</template>

<script>
import ListItem from './NiListItem'
import Page from './NiPage'
import Part from './NiPart'
import ToolBar from './NiToolBar'
export default {
  name: 'page-proposals-new',
  components: {
    ListItem,
    Page,
    Part,
    ToolBar
  }
}
</script>
