const state = {
  entityId: 0,
  groupId: 0
}

const mutations = {
  UPDATE_CURRENT_ENTITY_ID (state, entityId) {
    state.entityId = entityId
  },
  UPDATE_CURRENT_GROUP_ID (state, groupId) {
    state.groupId = groupId
  }
}

export default {
  state,
  mutations
}
