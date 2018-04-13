const state = {
  votes: []
}

const mutations = {
  ADD_VOTE (state, vote) {
    console.log('creating', JSON.stringify(vote))
    vote.updated_at = Date.now()
    state.votes.push(vote)
  },
  RM_VOTE (state, vote) {
    console.log('removing', JSON.stringify(vote))
    state.votes.splice(state.votes.indexOf(vote), 1)
  },
  UPDATE_VOTE (state, vote) {
    console.log('modifying', JSON.stringify(vote))
    let originalVote = state.votes.find(n => n.id === vote.id)
    originalVote.entity_id = vote.entity_id
    originalVote.value = vote.value
    originalVote.updated_at = vote.updated_at
  },
  ADD_BLANK_VOTES (state, proposal_id) {
    var entity_ids = ['john-connor', 'agent-janssen', 'janelle-voight', 'dr-silberman', 'serena-korgan', 'kate-brewster', 'miles-dyson', 'marcus-wright', 'kyle-reese', 'sarah-connor', 'machine-01', 'machine-02', 'machine-03', 'machine-04', 'machine-05', 't-x', 'guardian-01', 't-850', 't-800', 't-1000', 'jae-kwon', 'ethan-buchman', 'dustin-byington', 'peng-zhong']
    for (let i = 0; i < entity_ids.length; i++) {
      let vote = {
        proposal_id: proposal_id,
        entity_id: entity_ids[i],
        value: 0,
        updated_at: Date.now()
      }
      state.votes.push(vote)
    }
  },
  RM_VOTES_BY_PROPOSAL (state, vote_id) {
    console.log('removing votes with the Vote ID:', vote_id)
    state.votes.forEach(function (vote) {
      if (vote.proposal_id === vote_id) {
        state.votes.child(vote.id).remove()
      }
    })
  }
}

export default {
  state,
  mutations
}
