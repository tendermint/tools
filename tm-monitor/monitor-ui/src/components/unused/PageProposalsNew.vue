<template lang="pug">
page(title='New Proposal')
  part(title='Proposal Types')
    key-values
      key-value(@click.native="boo")
        div(slot="key") #[i.fa.fa-wrench] Adjust Parameters
        div(slot="value") Adjust parameters of this blockchain. 
      key-value(@click.native="boo")
        div(slot="key") #[i.fa.fa-file-text-o] Amend Constitution
        div(slot="value") Edit the Constitution with an amendment.
      key-value(@click.native="boo")
        div(slot="key") #[i.fa.fa-plus] Create Atoms 
        div(slot="value") Create new atoms and give them ownership.
      key-value(@click.native="boo")
        div(slot="key") #[i.fa.fa-pencil] Plain Text
        div(slot="value") Propose in plain text. (Markdown is supported.)
      key-value(@click.native="boo")
        div(slot="key") #[i.fa.fa-cloud-upload] Upgrade Code
        div(slot="value") Update the blockchain to a new version.
</template>

<script>
import $ from 'jquery'
import moment from 'moment'
import { mapGetters } from 'vuex'
import Btn from '@nylira/vue-button'
import FieldDate from './FieldDate'
import KeyValue from './NiKeyValue'
import KeyValues from './NiKeyValues'
import Page from './NiPage'
import Part from './NiPart'
export default {
  name: 'page-power',
  components: {
    Btn,
    FieldDate,
    KeyValue,
    KeyValues,
    Page,
    Part
  },
  computed: {
    ...mapGetters(['blockchain'])
  },
  data: () => ({
    showPowerMenu: true,
    showStopForm: false,
    showUpdateForm: false,
    stopDate: moment().add(7, 'days').valueOf(),
    updateDate: moment().add(7, 'days').valueOf(),
    jsonData: {valid: false, msg: '', content: ''}
  }),
  methods: {
    disableForms () {
      this.showPowerMenu = true
      this.showStopForm = false
      this.showUpdateForm = false
    },
    enableForm (form) {
      this.showPowerMenu = false

      if (form === 'stop') {
        $('#power-update').removeClass('active')
        this.showUpdateForm = false
        $('#power-stop').addClass('active')
        this.showStopForm = true
      } else if (form === 'update') {
        $('#power-stop').removeClass('active')
        this.showStopForm = false
        $('#power-update').addClass('active')
        this.showUpdateForm = true
      }
    },
    onSubmitStop () {
      alert('Stop proposal sent.')
    },
    uploadJson (event) {
      let fileInput = $('#blockchain-update-upload')[0]
      let jsonFile = fileInput.files[0]
      console.log('file uploaded:', jsonFile)

      let reader = new FileReader()
      reader.onload = (() => {
        return (evt) => {
          try {
            let jsonString = JSON.parse(evt.target.result)
            this.jsonData = {
              valid: true,
              msg: 'JSON parsed successfully',
              content: JSON.stringify(jsonString, null, '  ')
            }
          } catch (err) {
            this.jsonData = { valid: false, msg: err }
          }
        }
      })(jsonFile)
      reader.readAsText(jsonFile)
    },
    onSubmitUpdate (event) {
      try {
        alert('Update proposal sent.')
      } catch (err) {
        this.jsonData.valid = false
        this.jsonData.msg = err
      }
    }
  }
}
</script>

<style lang="stylus">
@require '../styles/variables.styl'

.power-menu
  margin-top 1rem
</style>
