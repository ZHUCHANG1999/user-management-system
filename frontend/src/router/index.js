import { createRouter, createWebHistory } from 'vue-router'

const routes = [
  {
    path: '/',
    name: 'UserList',
    component: () => import('@/views/UserList.vue')
  },
  {
    path: '/users/create',
    name: 'UserCreate',
    component: () => import('@/views/UserCreate.vue')
  },
  {
    path: '/users/:id',
    name: 'UserDetail',
    component: () => import('@/views/UserDetail.vue')
  },
  {
    path: '/users/:id/edit',
    name: 'UserEdit',
    component: () => import('@/views/UserEdit.vue')
  }
]

const router = createRouter({
  history: createWebHistory(),
  routes
})

export default router
