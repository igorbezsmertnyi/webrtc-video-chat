import Vuex from 'vuex'
import 'es6-promise/auto'

import Room from './room'
import Peer from './peer'

export default () => new Vuex.Store({
  modules: {
    Room,
    Peer
  }
})