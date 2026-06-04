<script setup>
import { inject, onMounted, ref } from 'vue';
import { useRoute } from 'vue-router';

const vueRoute = useRoute()
const host = inject('hostBacked')
const error = ref([])
const seats = ref([])
const carriages = ref([])


async function getTrain() {
	try {
		const response = await fetch(`${host}/api/getTrain/${vueRoute.params.id}`)
		if (response.ok) {
			carriages.value = await response.json()
		} else {
			const errorText = await response.text()
			const status = response.status
			error.value.push(status, errorText)
		}
	} catch (e) {
		console.error('Ошибка получения маршрутов:', e)
	}
}

async function getSeats(id) {
	try {
		const response = await fetch(`${host}/api/getCarriage/${id}`)
		if (response.ok) {
			seats.value = await response.json()
		} else {
			console.log("No ok getSeats")
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
		<div class="size-display-1">
			<div class="interactive-elem carriages">
				<button v-for="carriage in carriages" @click="getSeats(carriage)">Вагон {{ carriage }}</button>
			</div>
			<div>
				{{ seats }}
			</div>
		</div>
	</main>
</template>

<style scoped>
.carriages {
	display: flex;
	overflow-x: auto;
	gap: 8px;
	& > button {
		min-width: 6rem;
		color: var(--color-4);
		border: 1px solid var(--color-3-b);
		border-radius: 1rem;
		padding: 1rem 0.5rem;
		transition: background-color 0.3s;
	}
}
</style>