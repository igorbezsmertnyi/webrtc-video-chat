import Vue from 'vue'
import Router from 'vue-router'
import CreateRoom from '@/pages/CreateRoom'
import Room from '@/pages/Room'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'CreateRoom',
      component: CreateRoom
    },
    {
      path: '/room/:id',
      name: 'Room',
      component: Room
    }
  ]
})
