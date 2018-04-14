import Vue from 'vue'
import Router from 'vue-router'
import CreateRoom from '@/pages/CreateRoom'
import VideoStream from '@/components/VideoStream'

Vue.use(Router)

export default new Router({
  mode: 'history',
  routes: [
    {
      path: '/',
      name: 'CreateRoom',
      component: CreateRoom
    }
  ]
})
