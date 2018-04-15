import Routes from "@/static/routes"
import router from '@/router'

export default {
  state: {
    id: null,
    slug: null,
    peer: null,
    createdAt: null
  },

  getters: {
    room: state => state
  },

  mutations: {
    CREATE_ROOM(state, data) {
      state.id = data.id
      state.slug = data.slug
      state.peer = data.peer
      state.createdAt = data.created_at

      localStorage.setItem('room', JSON.stringify(data))
    },

    SET_ROOM(state, data) {
      state.id = data.id
      state.slug = data.slug
      state.peer = data.peer
      state.createdAt = data.created_at

      localStorage.setItem('room', JSON.stringify(data))
    },

    LEAVE_ROOM(state, id) {
      state.roomId = null
    }
  },

  actions: {
    async createRoom({ commit }, peer) {
      try {
        const res = await fetch(Routes.createRoomPath(), { 
          method: 'POST',
          headers: {
            'content-type': 'application/json'
          },
          body: JSON.stringify({ peer: peer.sdp })
        })
        const room = await res.json()
  
        await commit('CREATE_ROOM', room)
        await router.push({ path: `/room/${room.slug}` })
      } catch (err) {
        console.error(err)
      }
    },

    checkRoom({ commit }) {
      const room = JSON.parse(localStorage.getItem('room'))

      if (room) {
        commit('SET_ROOM', room)
      }
    }
  }
}