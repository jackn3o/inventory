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
const source = axios.CancelToken.source()
const instance = axios.create(axiosConfig)
instance.interceptors.request.use(config => {
  config.headers['Content-Type'] = 'application/json'

  return config
})

instance.interceptors.response.use(
  response => {
    // store.dispatch('setProgress', 100)
    return Promise.resolve(response)
  },
  error => {
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
