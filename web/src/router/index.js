import Vue from 'vue'
import Router from 'vue-router'
import store from '../store'
import InventoryList from '@/components/InventoryList'
import InventoryDetail from '@/components/InventoryDetail'
import Login from '@/components/Login'

//settings
import SettingList from '@/views/settings/List.vue'
import Category from '@/views/settings/Category.vue'
import Color from '@/views/settings/Color.vue'
import Item from '@/views/settings/Item.vue'
import Outlet from '@/views/settings/Outlet.vue'
// router, routes, auth guard
Vue.use(Router)

const routes = [
  {
    path: '/',
    name: 'inventory.list',
    meta: {
      auth: true
    },
    component: InventoryList,
    children: [
      {
        path: '/inventory/:id',
        name: 'inventory.detail',
        component: InventoryDetail
      }
    ]
  },
  {
    path: '/settings',
    name: 'settings',
    meta: {
      auth: true
    },
    component: SettingList,
    children: [
      {
        path: '/settings/category',
        name: 'settings.category',
        component: Category
      },
      {
        path: '/settings/color',
        name: 'settings.color',
        component: Color
      },
      {
        path: '/settings/item',
        name: 'settings.item',
        component: Item
      },
      {
        path: '/settings/outlet',
        name: 'settings.outlet',
        component: Outlet
      }
    ]
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

const router = new Router({
  mode: 'history',
  routes: routes
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.auth)) {
    // if (!store.getters.authorized) {
    //   next({
    //     path: '/login',
    //   })
    //   return
    // }
  }

  next()
})

router.afterEach((to, from, next) => {
  window.scrollTo(0, 0) // reset position
})

export default router
