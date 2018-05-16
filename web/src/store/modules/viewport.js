import * as types from '../mutation_types'


const state = {
    viewportSize: {
        'xs': false,
        'sm': false,
        'md': false,
        'lg': false,
        'xl': false
    }
}

const getters = {
    currentViewportSize(state) {
        for (let str_key in state.viewportSize) {
            if (state.viewportSize[str_key]) {
                return str_key
            }
        }
    }
}

const mutations = {
    [types.SET_VIEWPORT_SIZE](state, nb_viewportWidth) {
        // xs < 600px
        // sm 600px > < 960px*
        // md 960px > < 1264px*
        // lg 1264px > < 1904px*
        // xl > 1904px
        state.viewportSize = {
            xs: nb_viewportWidth < 600,
            sm: nb_viewportWidth > 600 && nb_viewportWidth <= 960,
            md: nb_viewportWidth > 960 && nb_viewportWidth <= 1264,
            lg: nb_viewportWidth > 1264 && nb_viewportWidth <= 1904,
            xl: nb_viewportWidth > 1904
        }
    }
}

const actions = {
    setViewportSize({ commit }, nb_viewportWidth) {
        commit(types.SET_VIEWPORT_SIZE, nb_viewportWidth)
    }
}

export default {
    state,
    getters,
    actions,
    mutations
}