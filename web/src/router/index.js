import Vue from 'vue'
import Router from 'vue-router'
import store from '../store'
import Login from '@/components/Login'

//inventory
import InventoryList from '@/views/inventory/List'
import InventoryDetail from '@/views//inventory/Detail'
//settings
import SettingList from '@/views/settings/List.vue'
import Category from '@/views/settings/Category.vue'
import Color from '@/views/settings/Color.vue'
import Cost from '@/views/settings/Cost.vue'
import Item from '@/views/settings/Item.vue'
import Outlet from '@/views/settings/Outlet.vue'
// router, routes, auth guard
Vue.use(Router)

const routes = [{
    path: '/',
    name: 'inventory.list',
    meta: {
      auth: true
    },
    component: InventoryList,
    children: [{
      path: '/inventory/:id',
      name: 'inventory.detail',
      component: InventoryDetail
    }]
  },
  {
    path: '/settings',
    name: 'settings',
    meta: {
      auth: true
    },
    component: SettingList,
    children: [{
        path: '/settings/category',
        name: 'settings.category',
        meta: {
          title: 'Category'
        },
        component: Category
      },
      {
        path: '/settings/color',
        name: 'settings.color',
        meta: {
          title: 'Color'
        },
        component: Color
      },
      {
        path: '/settings/cost',
        name: 'settings.cost',
        meta: {
          title: 'Cost'
        },
        component: Cost
      },
      {
        path: '/settings/item',
        name: 'settings.item',
        meta: {
          title: 'Item'
        },
        component: Item
      },
      {
        path: '/settings/outlet',
        name: 'settings.outlet',
        meta: {
          title: 'Outlet'
        },
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
