import * as types from '../mutation_types'

const state = {
    queue: [],
    timeout: 2000,
    toast: null
}

const getters = {
    toastQueue: state => state.queue,
    toastTimeout: state => state.timeout,
    toast: state => state.toast
}

const mutations = {
    [types.PUSH_TOAST](state, payload) {
        state.queue.push(payload)
    },

    [types.SHIFT_TOAST](state) {
        if (state.toast) {
            return
        }

        if (state.queue.length > 0) {
            state.toast = state.queue.shift()

            setTimeout(() => {
                state.toast = null
            }, state.timeout - 20)
        }
    },

    [types.RESET_TOAST](state) {
        state.toast = null
    }
}

const actions = {
    addToast({ commit }, toast) {
        if (typeof toast === 'string') {

            let str_type = ''
            if (toast.toLowerCase().includes('success')) {
                str_type = 'success'
            }

            if (toast.toLowerCase().includes('denied')
                || toast.toLowerCase().includes('error')
                || toast.toLowerCase().includes('wrong')) {
                str_type = 'error'
            }

            toast = {
                type: str_type,
                message: toast
            }
        }

        toast._id = Date.now()
        commit(types.PUSH_TOAST, toast)
    },
    // dismiss current active toast
    dismissToast({ commit }) {
        commit(types.RESET_TOAST)
    },

    // should called once
    // start calling toast
    startToaster({ commit }) {
        setInterval(() => {
            commit(types.SHIFT_TOAST)
        }, 500)
    }
}

export default {
    state,
    getters,
    actions,
    mutations
}