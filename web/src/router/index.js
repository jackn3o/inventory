import Vue from 'vue'
import Router from 'vue-router'
import InventoryList from '@/components/InventoryList'
import InventoryDetail from '@/components/InventoryDetail'
Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'inventory:list',
      component: InventoryList,
      children: [{
        path: '/inventory/:id',
        name: 'inventory:detail',
        component: InventoryDetail,
      }]
    }
  ]
})
