import Vuex from 'vuex'
import 'es6-promise/auto'

import Room from './room'
import Peer from './peer'
import ScreenShot from './screenShot'

export default () => new Vuex.Store({
  modules: {
    Room,
    Peer,
    ScreenShot
  }
})