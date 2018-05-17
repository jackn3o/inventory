// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import Vuetify from 'vuetify'
import Router from 'vue-router'
import moment from 'moment'
import VueMomentJS from 'vue-momentjs'
import axios from 'axios'
import VueAxios from 'vue-axios'
import lodash from 'lodash'
import VueLodash from 'vue-lodash'
import App from './App'
import axiosConfig from './config.js'
import router from './router'
import store from './store'
import 'vuetify/dist/vuetify.min.css'

Vue.use(Vuetify)
Vue.use(VueMomentJS, moment)
Vue.use(VueLodash, lodash)
Vue.config.productionTip = false
const instance = axios.create(axiosConfig)
instance.interceptors.request.use(config => {
  // set default headers
  config.headers['X-Client'] = 'inventory app'

  // if (config.method == 'post') {
  //   config.headers['Content-Type'] = 'text/plain'
  //   config.data = JSON.stringify(config.data)
  // } else {
  //   config.headers['Content-Type'] = 'application/json'
  // }
  // // only if content type not specified
  // if (config.headers['Content-Type'] === '') {
  //   config.headers['Content-Type'] = 'application/json'
  // }
  if (store.getters.authorized) {
    let token = store.getters.token
    config.headers['Authorization'] = `Bearer ${token}`
  } else {
  }

  return config
})

instance.interceptors.response.use(
  response => {
    // store.dispatch('setProgress', 100)
    return Promise.resolve(response)
  },
  error => {
    let obj_response = error.response.data,
      errorMsg = 'Something wrong, please try again later.'
    if (obj_response != null && obj_response.meta.messages != null) {
      errorMsg = obj_response.meta.messages
    }

    // toast no show for validation
    if (obj_response != null && !obj_response.meta.messages.validations) {
      store.dispatch('addToast', { type: 'error', message: errorMsg })
    }

    // store.dispatch('setProgress', 100)
    // redirect to login if unauthorized response
    // if (error.response.status === 401) {
    //   store.commit(types.LOGOUT)
    //   let fromRouteName = router.currentRoute.name,
    //     loginRouteName = 'authentication.login'

    //   if (fromRouteName != loginRouteName) {
    //     router.push({
    //       name: loginRouteName
    //     })

    //     store.dispatch('showToast', 'session expired')
    //   }
    // } else {
    return Promise.reject(error)
    // }
  }
)
Vue.use(VueAxios, instance)

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
