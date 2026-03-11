import { createRouter, createWebHistory } from 'vue-router'

// 检查是否登录
const isAuthenticated = () => {
  return !!localStorage.getItem('token')
}

const routes = [
  {
    path: '/login',
    name: 'Login',
    component: () => import('@/views/Login.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/register',
    name: 'Register',
    component: () => import('@/views/Register.vue'),
    meta: { requiresAuth: false }
  },
  {
    path: '/',
    name: 'UserList',
    component: () => import('@/views/UserList.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/users/create',
    name: 'UserCreate',
    component: () => import('@/views/UserCreate.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/users/:id',
    name: 'UserDetail',
    component: () => import('@/views/UserDetail.vue'),
    meta: { requiresAuth: true }
  },
  {
    path: '/users/:id/edit',
    name: 'UserEdit',
    component: () => import('@/views/UserEdit.vue'),
    meta: { requiresAuth: true }
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

// 路由守卫
router.beforeEach((to, from, next) => {
  const requiresAuth = to.meta.requiresAuth !== false
  
  if (requiresAuth && !isAuthenticated()) {
    next('/login')
  } else if ((to.path === '/login' || to.path === '/register') && isAuthenticated()) {
    next('/')
  } else {
    next()
  }
})

export default router
