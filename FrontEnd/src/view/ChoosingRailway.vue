<script setup>
import { useAuthStore } from '@/store/auth';
import { usePassengerStore } from '@/store/passenger';
import { inject, onMounted, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const store = usePassengerStore()
const storeAuth = useAuthStore()
const vueRouter = useRouter()
const vueRoute = useRoute()
const routes = ref([])
const error = ref([])
const host = inject('hostBacked')

function stringToDate(date) {
	const todayString = new Date(date).toISOString().split('T')[0]
	return todayString	
}

function dayName(date) {
	date = new Date(date)
	let weekday = date.toLocaleString("ru", { weekday: 'long'})
	weekday = weekday.charAt(0).toUpperCase() + weekday.slice(1)
	const numDayMonth = date.toLocaleString("ru", { day: '2-digit', month: 'long' })
	return `${weekday} - ${numDayMonth}`
}

function formationTime(date) {
	date = new Date(date)
	const hour = String(date.getHours()).padStart(2, '0')
	const minute = String(date.getMinutes()).padStart(2, '0')
	return `${hour}:${minute}`
}

function calculatingPrice(distance) {
	// (distance * 5) * Пассажир 
	let result = (distance * 5) * 1
	return result
}

async function getRoute() {
	try {
		const url = new URL(`${host}/api/getRoutes`)
		url.searchParams.set('from', vueRoute.query.from)
		url.searchParams.set('to', vueRoute.query.to)
		url.searchParams.set('date', vueRoute.query.date)
		const response = await fetch(url)
		if (response.ok) {
			routes.value = await response.json()
		} else {
			const errorText = await response.text()
			const status = response.status
			error.value.push(status, errorText)
		}
	} catch (e) {
		console.error('Ошибка получения маршрутов:', e)
	}
}

let lastDate = null
function checkDate(route) {
	if (lastDate === null) {
		lastDate = stringToDate(route.sending)
		return true
	}
	if (lastDate === stringToDate(route.sending)) {
		return false
	}
	lastDate = stringToDate(route.sending)
	return true
}

function buyTicket(route) {
	if (store.data.passengers.length === 0) {
		vueRouter.push({name: "home"})
		return
	}
	if (storeAuth.userExits) {
		vueRouter.push({name: "not-registered"})
		return
	}
	route.price = calculatingPrice(route.distance)
	store.route = route
	vueRouter.push({name: "choosing-seats", params: {id: route.id}})
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
			<RouterLink to="/" class="a-link">На главную</RouterLink>
		</div>
		<div class="horizontal-scroll-container">
			<div v-for="route in routes">
				<h1 v-if="checkDate(route)">{{ dayName(route.sending) }}</h1>
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
						<span>Цена: {{ calculatingPrice(route.distance) }}</span>
						<button @click="buyTicket(route)" class="interactive-elem">Выбрать</button>
					</div>
				</div>
			</div>
		</div>
	</main>
</template>

<style scoped>
main {
	margin-top: 1rem;
}

h1 {
	margin-top: 1rem;
}

.size-display-1 {
	margin-top: 8px;
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

.horizontal-scroll-container {
	overflow-x: auto;
	overflow-y: hidden;
	max-width: 100vw;
}

button {
	width: 12rem;
	border-radius: 1rem;
	padding: 0.8rem;
	color: var(--color-4);
	font-size: 18px;
	transition: border-color 0.3s, background-color 0.3s;
}

span {
	font-size: 18px;
	font-weight: bold;
	border: 2px solid var(--color-3-b);
	border-radius: 1rem;
	padding: 0.8rem;
}
</style>