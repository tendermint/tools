import shortid from 'shortid'
function Validator () {
  return {
    id: shortid.generate(),
    pub_key: '',
    info: '',
    type: '',
    created_at: 0,
    online_at: 0
  }
}

const state = {
  validator: JSON.parse(JSON.stringify(new Validator()))
}

const mutations = {
  RESET_VALIDATORS_NEW (state) {
    state.validator = JSON.parse(JSON.stringify(new Validator()))
  }
}

export default {
  state,
  mutations
}
