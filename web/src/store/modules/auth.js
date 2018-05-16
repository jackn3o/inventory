import * as types from '../mutation_types'

const state = {
    authorized: false,
    token: null,
}

const getters = {
    authorized(state) {
        if (localStorage.getItem('auth')) {
            const obj_auth = JSON.parse(localStorage.getItem('auth')),
                ts_now = Math.round((new Date()).getTime() / 1000)
            if (ts_now > obj_auth.expiredAt) {
                //todo
            } else {
                state.authorized = true
            }
        }
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
    [types.LOGIN](state, obj_auth) {
        state.authorized = true;
        state.token = obj_auth.token;
        localStorage.setItem('auth', JSON.stringify(obj_auth));
    },

    [types.LOGOUT](state) {
        state.authorized = false;
        state.token = null;
        localStorage.removeItem('auth');

    }
}

const actions = {
    login({ commit }, obj_auth) {
        commit(types.LOGIN, obj_auth)
    },
    logout({ commit }) {
        commit(types.LOGOUT)
    }
}

export default {
    state,
    getters,
    mutations,
    actions,
}