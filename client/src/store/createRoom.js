export default {
  state: {
    roomId: null
  },

  mutations: {
    CREATE_ROOM(state, id) {
      state.roomId = id
    },

    LEAVE_ROOM(state, id) {
      state.roomId = null
    }
  },

  actions: {
    async createRoom({ commit }) {

    }
  }
}