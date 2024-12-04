import {createRouter, createWebHistory} from 'vue-router';
import HomeView from '../views/HomeView.vue';

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
					component: import('../views/AllBoxes.vue'),
				},
				{
					name: 'boxContents',
					path: 'box/:id',
					component: import('../views/BoxContents.vue'),
				},
			],
		},
	],
});

export default router;
