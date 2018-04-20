import Peer from 'simple-peer'

export default {
  state: {
    peer: {
      sdp: null,
      type: null
    }
  },

  getters: {
    currentPeer: state => state.peer
  },

  mutations: {
    SET_PEER(state, data) {
      state.peer.sdp = data.sdp
      state.peer.type = data.type
    }
  },

  actions: {
    setPeer({ commit }, data) {
      commit('SET_PEER', data)
    },

    createPeer({ commit }, data) {
      const peer = new Peer({
        initiator: true,
        trickle: false
      })
  
      peer.on('signal', e => commit('SET_PEER', e))
    },

    connectPeer({ commit }, data) {
      console.log('connect')
      this.state.peer.signal(data)
    },

    send({ commit, state }, data) {
      console.log(this, data, state.peer)
    }
  }
}