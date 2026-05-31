<script setup>
import { RouterLink } from 'vue-router';
import { useAuthStore } from '@/store/auth';
import { useRoute } from 'vue-router'

const authStore = useAuthStore()
const route = useRoute()
</script>

<template>
<header>
	<RouterLink to="/" class="no-a"><h1>PuPu</h1></RouterLink>
	<div v-if="authStore.isLoggedIn">
		<span>{{ authStore.fullName }}</span>
		<RouterLink @click="authStore.logout()" to="/login" class="light-inter-elem">Выйти</RouterLink>
	</div>
	<div v-else-if="route.path === '/login' ? false : true">
		<RouterLink to="/login" class="light-inter-elem">Войти</RouterLink>
	</div>
</header>
</template>

<style scoped>
header {
	color: white;
	display: flex;
	justify-content: space-between;
	align-items: center;
	padding: 0.5rem 0;
	border-bottom: 2px solid var(--color-3-b);
}

span {
	margin-right: 8px;
	font-weight: 500;
}
</style>