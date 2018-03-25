import Vue from 'vue'
import Router from 'vue-router'
import InventoryList from '@/components/InventoryList'
import InventoryDetail from '@/components/InventoryDetail'
import Login from '@/components/Login'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'inventory.list',
      meta: {
        auth: true
      },
      component: InventoryList,
      children: [{
        path: '/inventory/:id',
        name: 'inventory.detail',
        component: InventoryDetail,
      }]
    },
    {
      path: '/login',
      name: 'login',
      meta: {
        auth: false
      },
      component: Login
    }
  ]
})
