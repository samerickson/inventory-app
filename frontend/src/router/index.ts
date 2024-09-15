import { createRouter, createWebHistory } from 'vue-router'
import HomeView from '../views/HomeView.vue'

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: '/',
      name: 'home',
      component: HomeView,
      children: [
        {
          name: 'allBoxes',
          path: '',
          component: import('../views/allBoxes.vue')
        },
        {
          name: 'boxContents',
          path: ':id',
          component: import('../views/BoxContents.vue')
        }
      ]
    },
  ]
})

export default router
