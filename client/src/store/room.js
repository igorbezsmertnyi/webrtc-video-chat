import Routes from "@/static/routes"
import router from '@/router'

export default {
  state: {
    id: null,
    slug: null,
    createdAt: null,
    created: false
  },

  getters: {
    room: state => state
  },

  mutations: {
    CREATE_ROOM(state, data) {
      state.id = data.id
      state.slug = data.slug
      state.createdAt = data.created_at
      state.created = true

      localStorage.setItem('room', JSON.stringify(state))
    },

    SET_ROOM(state, data) {
      state.id = data.id
      state.slug = data.slug
      state.createdAt = data.created_at

      localStorage.setItem('room', JSON.stringify(state))
    },

    LEAVE_ROOM(state, id) {
      state.roomId = null
    }
  },

  actions: {
    async createRoom({ commit }) {
      try {
        const res = await fetch(Routes.createRoomPath(), { 
          method: 'POST',
          headers: {
            'content-type': 'application/json'
          }
        })
        const room = await res.json()
  
        await commit('CREATE_ROOM', room)
        await router.push({ path: `/room/${room.slug}` })
      } catch (err) {
        console.error(err)
      }
    },

    async getRoom({ commit }, slug) {
      try {
        const res = await fetch(Routes.getRoomPath(slug), {
          method: 'GET',
          headers: {
            'content-type': 'application/json'
          }
        })
        const room = await res.json()
        await commit('SET_ROOM', room)
      } catch (err) {
        console.error(err)
      }
    },

    checkRoom({ commit }) {
      const room = JSON.parse(localStorage.getItem('room'))

      if (room && (room.slug == router.history.current.params.id) && room.created) {
        commit('CREATE_ROOM', room)
      } else {
        this.dispatch('getRoom', router.history.current.params.id)
      }
    }
  }
}
