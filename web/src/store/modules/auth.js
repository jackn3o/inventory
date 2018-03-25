import * as types from '../mutation_types'

const state = {
    authorized: localStorage.getItem('auth.token') != null,
    token: null,
}

const getters = {
    authorized(state) {
        return state.authorized
    },
    token(state) {
        if (state.token == null) {
            state.token = JSON.parse(localStorage.getItem('auth.token'));
        }
        return state.token;
    }
}

const mutations = {
    [types.LOGIN](state, token) {
        state.authorized = true;
        state.token = token;
        localStorage.setItem('auth.token', JSON.stringify(token));
    },

    [types.LOGOUT](state, token) {
        state.authorized = false;
        state.token = null;
        localStorage.removeItem('auth.token');
    }
}

const actions = {
}

export default {
    state,
    getters,
    mutations,
    actions,
}