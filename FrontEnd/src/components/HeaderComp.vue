<script setup>
import { RouterLink, useRoute } from 'vue-router';

const route = useRoute()
</script>

<template>
<header>
	<RouterLink to="/" class="no-a"><h1>PuPu</h1></RouterLink>
	<RouterLink to="/login" v-if="route.name !== 'login'" class="light-inter-elem">Вход</RouterLink>
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
</style>