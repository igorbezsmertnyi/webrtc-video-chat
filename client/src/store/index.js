import Vuex from 'vuex'
import 'es6-promise/auto'
import Room from './room'

export default () => new Vuex.Store({
  modules: {
    Room
  }
})