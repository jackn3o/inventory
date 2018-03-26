import Vue from 'vue'
import Router from 'vue-router'
import store from '../store'
import InventoryList from '@/components/InventoryList'
import InventoryDetail from '@/components/InventoryDetail'
import Login from '@/components/Login'

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


const router = new Router({
  mode: 'history',
  routes: routes,
})

router.beforeEach((to, from, next) => {
  if (to.matched.some(record => record.meta.auth)) {
    if (!store.getters.authorized) {
      next({
        path: '/login',
      })
      return
    }
  }

  next()
})

router.afterEach((to, from, next) => {
  window.scrollTo(0, 0) // reset position
})

export default router
