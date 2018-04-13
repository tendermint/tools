import moment from 'moment'
import shortid from 'shortid'

function Proposal () {
  let id = shortid.generate()
  return {
    vote_id: id,
    active_at: moment().add(30, 'days').valueOf(),
    created_at: '',
    entity_id: '',
    title: `Proposal ${id}`,
    type: 'text',
    flags: {
      new: true,
      read: false,
      passed: false
    },
    data: {
      group_id: '',
      body: ''
    }
  }
}

const state = {
  proposal: JSON.parse(JSON.stringify(new Proposal()))
}

const mutations = {
  RESET_PROPOSAL_TEXT (state) {
    state.proposal = JSON.parse(JSON.stringify(new Proposal()))
  },
  SET_PROPOSAL_TEXT_GROUP_ID (state, group_id) {
    state.proposal.data.group_id = group_id
  }
}

export default {
  state,
  mutations
}
