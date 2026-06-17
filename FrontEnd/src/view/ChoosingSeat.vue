<script setup>
import KupeCarriage from '@/components/KupeCarriage.vue';
import PlatskartCarriage from '@/components/PlatskartCarriage.vue';
import SVCarriage from '@/components/SVCarriage.vue';
import { useAuthStore } from '@/store/auth';
import { usePassengerStore } from '@/store/passenger';
import { computed, inject, onMounted, reactive, ref } from 'vue';
import { useRoute, useRouter } from 'vue-router';

const vueRoute = useRoute()
const store = usePassengerStore()
const storeAuth = useAuthStore()
const vueRouter = useRouter()
const host = inject('hostBacked')
const error = ref([])
const currentCarriage = reactive({ seats: [], type: '', id: null })
const carriages = ref([])

const passengerCount = computed(() => store.data.passengers.length)
const selectedCount = computed(() => store.selectedSeats.length)
const remainingSeats = computed(() => passengerCount.value - selectedCount.value)

async function getCarriage(id) {
	try {
		const response = await fetch(`${host}/api/getCarriage/${id}`)
		if (response.ok) {
			const seats = await response.json()
			currentCarriage.seats = seats
			currentCarriage.type = getCarriageType(seats.length)
			currentCarriage.id = id
			store.clearSelectedSeats()
		} else {
			console.log("No ok getSeats")
		}
	} catch (e) {
		console.error('Ошибка получения маршрутов:', e)
	}
}

function getCarriageType(seatCount) {
	switch (seatCount) {
		case 54:
			return 'Плацкарт';
		case 36:
			return 'Купе';
		case 18:
			return 'СВ';
		default:
			return 'Плацкарт';
	}
}

async function getTrain() {
	try {
		const response = await fetch(`${host}/api/getTrain/${vueRoute.params.id}`)
		if (response.ok) {
			carriages.value = await response.json()
			if (carriages.value.length > 0) {
				getCarriage(carriages.value[0])
			}
		} else {
			const errorText = await response.text()
			const status = response.status
			error.value.push(status, errorText)
		}
	} catch (e) {
		console.error('Ошибка получения маршрутов:', e)
	}
}

async function bookingSeats(req_data) {
	try {
		const response = await fetch(`${host}/api/bookingSeats`, {
			method: 'POST', body: JSON.stringify(req_data),
			headers: {'Login': storeAuth.user.login, 'Content-Type': 'application/json'}, 
		})
		if (response.ok) {
			vueRouter.push({ name: 'adding-passengers' })
		} else if (response.status === 404) {
			storeAuth.logout()
			vueRouter.push({name: "not-registered"})
		}
	} catch (e) {
		console.error('Ошибка:', e)
	}
}

async function confirmSelection() {
	if (selectedCount.value === passengerCount.value) {
		const req_data = {
			[currentCarriage.id]: store.getSelectedSeats
		}
		try {
			const response = await fetch(`${host}/api/checkingSeats`, {
				method: 'POST', body: JSON.stringify(req_data),
				headers: {'Content-Type': 'application/json'}
			})
			if (response.ok) {
				bookingSeats(req_data)
			} else {
				const errorText = await response.text()
				const status = response.status
				error.value.push(status, errorText)
			}
		} catch (e) {
			console.error('Ошибка:', e)
		}
	}
}

onMounted(() => {
	if (store.data.passengers.length === 0) {
		vueRouter.push({name: "home"})
		return
	}
	if (storeAuth.userExits) {
		vueRouter.push({name: "not-registered"})
		return
	}
	getTrain()
})
</script>

<template>
	<main>
		<div v-if="error.length > 0" class="error-div">
			<h2>{{ error[0] }}</h2>
			<p class="red-text">{{ error[1] }}</p>
			<RouterLink to="/" class="a-link">На главную</RouterLink>
		</div>
		<div v-else class="flex-display-column">
			<div>
				<h1>Выбор вагонов</h1>
				<div class="interactive-elem carriages">
					<button v-for="carriage in carriages" :key="carriage" @click="getCarriage(carriage)" :class="{active: currentCarriage.id === carriage}">Вагон {{ carriage }}</button>
				</div>
				<div class="interactive-elem">
					<p>Количество пассажиров: <strong>{{ passengerCount }}</strong></p>
					<p>Выбрано мест: <strong>{{ selectedCount }}</strong></p>
					<p v-if="remainingSeats > 0" style="color: #ff9800;">Осталось выбрать: <strong>{{ remainingSeats }}</strong></p>
					<strong v-else style="color: green;">Все места выбраны</strong>
				</div>
			</div>
			<div class="div-carriage">
				<h1>{{ currentCarriage.type }}</h1>
				<PlatskartCarriage v-if="currentCarriage.type === 'Плацкарт'" :seats="currentCarriage.seats"/>
				<KupeCarriage v-else-if="currentCarriage.type === 'Купе'" :seats="currentCarriage.seats"/>
				<SVCarriage v-else-if="currentCarriage.type === 'СВ'" :seats="currentCarriage.seats"/>
			</div>
			<div class="button-row">
				<button @click="store.clearSelectedSeats()" :disabled="selectedCount === 0">Очистить выбор</button>
				<button @click="confirmSelection" :disabled="selectedCount !== passengerCount">Подтвердить выбор</button>
			</div>
		</div>
	</main>
</template>

<style scoped>
.carriages {
	max-width: 44rem;
	display: flex;
	overflow-x: auto;
	gap: 8px;
	& > button {
		min-width: 6rem;
		color: var(--color-4);
		border-radius: 1rem;
		padding: 1rem 0.5rem;
		transition: background-color 0.3s;
	}
}

.carriages button.active {
	background: var(--color-2);
}

.button-row {
	display: flex;
	margin-top: 1rem;
	& button {
		padding: 1rem;
		border-radius: 1rem;
		font-size: 18px;
		transition: background-color 0.3s ease-in;
		&:nth-child(1) {
			margin-right: 8px;
			background-color: brown;
		}
	}
}

.div-carriage {
	margin-top: 1rem;
	overflow-x: auto;
	overflow-y: hidden;
	max-width: 100vw;
	& h1 {
		margin-left: 1rem;
	}
	& div {
		width: max-content;
		margin-right: 1rem;
		margin-left: 1rem;
	}
}
</style>