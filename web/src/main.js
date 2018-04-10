// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import Router from 'vue-router'
import router from './router'
import Vuetify from 'vuetify'
import store from './store'
import moment from "moment";
import VueMomentJS from "vue-momentjs";
import lodash from 'lodash'
import VueLodash from 'vue-lodash'
import axios from 'axios'
import VueAxios from 'vue-axios'

import 'vuetify/dist/vuetify.min.css'

Vue.use(Vuetify)
Vue.use(VueMomentJS, moment);
Vue.use(VueLodash, lodash)
Vue.use(VueAxios, axios)
Vue.config.productionTip = false


/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
