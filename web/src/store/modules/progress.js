import * as types from '../mutation_types'

const state = {
  progress: 0
}

const getters = {
  progress: state => state.progress
}

const mutations = {
  [types.SET_PROGRESS](state, payload) {
    state.progress = payload
  }
}

const actions = {
    setProgress({ commit }, nb_progress) {
        
      commit(types.SET_PROGRESS, nb_progress)
      if (nb_progress === 100) {
        setTimeout(() => {
          commit(types.SET_PROGRESS, 0)
        }, 500)
      }
    }
  }

export default {
  state,
  getters,
  actions,
  mutations
}
