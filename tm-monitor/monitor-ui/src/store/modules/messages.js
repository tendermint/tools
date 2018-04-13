const state = {
  messages: []
}

const mutations = {
  ADD_MSG (state, message) {
    console.log('creating', JSON.stringify(message))
    message.created_at = Date.now()
    state.messages.push(message)
  },
  RM_MSG (state, message) {
    console.log('removing', JSON.stringify(message))
    state.messages.splice(state.messages.indexOf(message), 1)
  },
  RM_MSGS_BY_PROPOSAL (state, proposal_id) {
    console.log('removing messages with the Proposal ID:', proposal_id)
    state.messages.forEach(function (message) {
      if (message.proposal_id === proposal_id) {
        state.messages.splice(state.messages.indexOf(message), 1)
      }
    })
  }
}

export default {
  state,
  mutations
}
