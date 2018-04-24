export default {
  state: {
    img: null
  },

  getters: {
    currentScreenShot: state => state.img
  },

  mutations: {
    SET_IMAGE(state, data) {
      state.img = data
    },

    REMOVE_IMAGE(state) {
      state.img = null
    }
  },

  actions: {
    setImage({ commit }, data) {
      commit('SET_IMAGE', data)
    },

    removeImage({ commit }) {
      commit('REMOVE_IMAGE')
    }
  }
}