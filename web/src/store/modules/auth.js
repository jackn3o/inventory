import * as types from '../mutation_types'


const state = {
    authorized: localStorage.getItem('auth') && JSON.parse(localStorage.getItem('auth')).token,
    token: null,
}

const getters = {
    authorized(state) {
        return state.authorized
    },
    token(state) {
        if (state.token == null) {
            if (localStorage.getItem('auth')) {
                state.token = JSON.parse(localStorage.getItem('auth')).token;
            }
        }
        return state.token;
    }
}

const mutations = {
    [types.LOGIN](state, token) {
        state.authorized = true;
        state.token = token;
        localStorage.setItem('auth', JSON.stringify(token));
    },

    [types.LOGOUT](state, token) {
        state.authorized = false;
        state.token = null;
        localStorage.removeItem('auth');
    }
}

const actions = {
    login({ commit }, obj_auth) {
        commit(types.LOGIN, obj_auth)
    },
}

export default {
    state,
    getters,
    mutations,
    actions,
}