<script setup>
import KupeCarriage from '@/components/KupeCarriage.vue';
import PlatskartCarriage from '@/components/PlatskartCarriage.vue';
import SVCarriage from '@/components/SVCarriage.vue';
import { inject, onMounted, reactive, ref } from 'vue';
import { useRoute } from 'vue-router';

const vueRoute = useRoute()
const host = inject('hostBacked')
const error = ref([])
const currentCarriage = reactive({ seats: [], type: '' })
const carriages = ref([])


async function getCarriage(id) {
	try {
		const response = await fetch(`${host}/api/getCarriage/${id}`)
		if (response.ok) {
			const seats = await response.json()
			currentCarriage.seats = seats
			currentCarriage.type = getCarriageType(seats.length)
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

onMounted(() => {
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
					<button v-for="carriage in carriages" :key="carriage" @click="getCarriage(carriage)">Вагон {{ carriage }}</button>
				</div>
			</div>
			<div class="div-carriage">
				<h1>{{ currentCarriage.type }}</h1>
				<PlatskartCarriage v-if="currentCarriage.type === 'Плацкарт'" :seats="currentCarriage.seats"/>
				<KupeCarriage v-else-if="currentCarriage.type === 'Купе'" :seats="currentCarriage.seats"/>
				<SVCarriage v-else-if="currentCarriage.type === 'СВ'" :seats="currentCarriage.seats"/>
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