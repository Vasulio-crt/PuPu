import { createRouter, createWebHistory } from 'vue-router'

const router = createRouter({
	history: createWebHistory(import.meta.env.BASE_URL),
	routes: [
		{
			path: '/',
			name: 'home',
			component: () => import('../view/SearchRailway.vue'),
		},
		{
			path: '/railway',
			name: 'choosing-railway',
			component: () => import('../view/ChoosingRailway.vue'),
		},
		{
			path: '/register',
			name: 'register',
			component: () => import('../view/Register.vue'),
		},
		{
			path: '/login',
			name: 'login',
			component: () => import('../view/Login.vue'),
		},
	],
})

export default router
