import request from 'superagent'
let url = 'http://116.62.62.39:46657'

const state = {
  status: {
    listen_addr: '',
    sync_info: {
      latest_block_height: 0,
      latest_block_hash: ''
    }
  }
}

const mutations = {
  getStatus (state) {
    request.get(url + '/status').end((err, res) => {
      if (err) { console.log('err', err) }
      state.status = res.body.result
    })
  }
}

export default {
  state,
  mutations
}
