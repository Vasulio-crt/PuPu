<script setup>
import { usePassengerStore } from '@/store/passenger';
import { inject, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

defineOptions({
	name: "choosing-railway"
})

const store = usePassengerStore()
const route = useRoute()
const error = ref([])
const host = inject('hostBacked')

function dayName(date) {
	date = new Date(date)
	let weekday = date.toLocaleString("ru", { weekday: 'long'})
	weekday = weekday.charAt(0).toUpperCase() + weekday.slice(1)
	return `${weekday} ${store.getToday()}`
}

function formationTime(date) {
	date = new Date(date)
	const hour = String(date.getHours()).padStart(2, '0')
	const minute = String(date.getMinutes()).padStart(2, '0')
	return `${hour}:${minute}`
}

function calculatingPrice() {
	return 10101
}

async function getRoute() {
	try {
		const url = new URL(`${host}/api/getRoutes`)
		url.searchParams.set('from', route.query.from)
		url.searchParams.set('to', route.query.to)
		url.searchParams.set('date', route.query.date)
		const response = await fetch(url)
		if (response.ok) {
			const routes = await response.json()
			store.routes = routes
		} else {
			const errorText = await response.text()
			const status = response.status
			error.value.push(status, errorText)
			console.error('Ошибка от сервера:', errorText)
		}
	} catch (e) {
		console.error('Ошибка получения маршрутов:', e)
	}
}

onMounted(() => {
	getRoute()
})
</script>

<template>
	<main>
		<div v-if="error.length > 0" class="error-div">
			<h2>{{ error[0] }}</h2>
			<p class="red-text">{{ error[1] }}</p>
			<RouterLink to="/">На главную</RouterLink>
		</div>
		<div v-for="route in store.routes">
			<h1>{{ dayName(route.sending) }}</h1>
			<div class="size-display-1 interactive-elem">
				<div class="display-2">
					<h2>{{ formationTime(route.sending) }}</h2>
					<div class="route-line">
						<p>{{ route.distance }} км</p>
					</div>
					<h2>{{ formationTime(route.arrival) }}</h2>
				</div>
				<div class="display-2">
					<h3>{{ route.from_station }}</h3>
					<h3>{{ route.to_station }}</h3>
				</div>
				<div class="display-2">
					<span>Цена: {{ calculatingPrice() }}</span>
					<button class="interactive-elem">Купить</button>
				</div>
			</div>
			
		</div>
	</main>
</template>

<style scoped>
.error-div {
	display: flex;
	flex-direction: column;
	align-items: center;
	border: 2px solid var(--color-3-b);
	border-radius: 24px;
	padding: 1rem 2rem;
	font-size: 1.3em;
}

a {
	margin-top: 8px;
	text-decoration: none;
	color: var(--color-2);
	transition: color 0.3s ease;
	&:hover {
		color: var(--color-1);
	}
}

main {
	margin-top: 1rem;
}

h1 {
	margin-top: 1rem;
}

.size-display-1 {
	padding: 1rem 2rem;
}

.display-2 {
	display: flex;
	justify-content: space-between;
	align-items: center;
	&:nth-child(3) {
		margin-top: 0.5rem;
	}
}

.route-line {
	display: flex;
	flex-direction: column;
	align-items: center;
	flex-grow: 1;
	border-bottom: 2px dashed var(--color-3-b);
	margin: 0 1rem;
}

button {
	width: 12rem;
	border: 2px solid var(--color-3-b);
	border-radius: 1rem;
	padding: 0.8rem;
	color: var(--color-4);
	font-size: 18px;
	transition: border-color 0.3s, background-color 0.3s;
}

span {
	font-size: 20px;
	font-weight: bold;
	border: 2px solid var(--color-3-b);
	border-radius: 1rem;
	padding: 0.8rem;
}
</style>