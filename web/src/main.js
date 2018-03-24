// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import Vuetify from 'vuetify'
import 'vuetify/dist/vuetify.min.css'
import store from './store'
import moment from "moment";
import VueMomentJS from "vue-momentjs";
import lodash from 'lodash'
import VueLodash from 'vue-lodash'

 
Vue.use(Vuetify)
Vue.use(VueMomentJS, moment);
Vue.use(VueLodash, lodash)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  store,
  components: { App },
  template: '<App/>'
})
