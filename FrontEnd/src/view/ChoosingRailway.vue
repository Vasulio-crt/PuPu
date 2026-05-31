<script setup>
import { onMounted, ref, watch } from 'vue';
import UserPlusIcon from '@/components/icons/UserPlusIcon.vue';
import XIcon from '@/components/icons/XIcon.vue';
import { usePassengerStore } from '@/store/passenger';
import UserIcon from '@/components/icons/UserIcon.vue';

defineOptions({
	name: 'home'
})

const passengers = ref([])
const isPassengerMenuOpen = ref(false)
const store = usePassengerStore()

let nextId = 0
const passengerTypes = [
	{ name: 'Пассажир', color: 'DarkCyan' },
	{ name: 'Студент', color: 'DarkBlue' },
	{ name: 'Ребенок', color: 'DarkOrchid' },
	{ name: 'Пенсионер', color: 'DarkRed' }
]

function addPassenger(type) {
	if (passengers.value.length === 0) nextId = 0
	passengers.value.push({ id: nextId, ...type })
	store.data.passengers.push({ id: nextId++, name: type.name })
	isPassengerMenuOpen.value = false
}

function removePassenger(passengerId) {
	passengers.value = passengers.value.filter(p => p.id !== passengerId)
	store.deletePassenger(passengerId)
}

function togglePassengerMenu() {
	isPassengerMenuOpen.value = !isPassengerMenuOpen.value
}

onMounted(() => {
	const today = new Date();
	store.data.date = new Date(today.getTime()).toISOString().split('T')[0];
	const savedPassengers = localStorage.getItem('passengers_view')
	const savedStoreData = localStorage.getItem('store_data')

	if (savedPassengers && savedStoreData) {
		passengers.value = JSON.parse(savedPassengers)
		const parsedStoreData = JSON.parse(savedStoreData)
		store.data.from = parsedStoreData.from
		store.data.to = parsedStoreData.to
		store.data.passengers = parsedStoreData.passengers
		if (passengers.value.length > 0) {
			nextId = Math.max(...passengers.value.map(p => p.id)) + 1
		}
	} else {
		addPassenger(passengerTypes[0])
	}
})

watch(passengers, (newValue) => localStorage.setItem('passengers_view', JSON.stringify(newValue)), {deep: true})
watch(store.data, (newValue) => localStorage.setItem('store_data', JSON.stringify(newValue)))
</script>

<template>
	<main>
		<div class="size-display-1">
			<h2>Путь</h2>
			<div class="search-row">
				<input v-model="store.data.from" type="text" placeholder="Откуда" class="interactive-elem">
				<input v-model="store.data.to" type="text" placeholder="Куда" class="interactive-elem">
			</div>
		</div>
		<div class="size-display-1">
			<h2>Дата отправления</h2>
			<div class="search-row">
				<input v-model="store.data.date" type="date" class="interactive-elem">
			</div>
		</div>
		<div class="size-display-1">
			<h2>Пассажиры</h2>
			<div class="passengers">
				<div v-for="passenger in passengers" :key="passenger.id" class="interactive-elem passenger">
					<UserIcon :style="{ backgroundColor: passenger.color }" />
					<p>{{ passenger.name }}</p>
					<XIcon @click="removePassenger(passenger.id)" class="remove-passenger-btn" width="24" height="24" />
				</div>
				<div>
					<div @click="togglePassengerMenu" class="interactive-elem passenger" style="cursor: pointer;">
						<UserPlusIcon class="white-and-black" />
						<p>Добавить пассажира</p>
					</div>
					<div v-if="isPassengerMenuOpen" class="passenger-dropdown">
						<div v-for="pType in passengerTypes" :key="pType.name" @click="addPassenger(pType)" class="dropdown-item">
							<UserIcon :style="{ backgroundColor: pType.color, borderRadius: '40%', padding: '4px' }" />
							<span>{{ pType.name }}</span>
						</div>
					</div>
				</div>
			</div>
		</div>
		<div class="size-display-1">
			<button @click="store.getTrainFlights()" class="interactive-elem">НАЙТИ</button>
		</div>
	</main>
</template>

<style scoped>
.passengers {
	display: grid;
	grid-template-columns: 1fr 1fr;
	gap: 12px;
	width: 100%;
	max-width: 800px;
}

.passenger {
	display: flex;
	align-items: center;
	&:hover:not(:has(.remove-passenger-btn:hover)) {
		border-color: var(--color-3-b-hover);
		background-color: var(--color-3-a-hover);
	}
	& > p {
		margin-left: 12px;
	}
	& > svg {
		padding: 8px;
		border-radius: 1rem;
	}
}

.white-and-black {
	color: var(--color-3);
	background-color: var(--color-4);
}

.passenger-dropdown {
	position: absolute;
	border: 2px solid var(--color-3-b);
	background-color: var(--color-3-a);
	border-radius: 8px;
	z-index: 10;
	width: 100%;
	max-width: 15rem;
	margin-top: 5px;
}

.dropdown-item {
	display: flex;
	align-items: center;
	gap: 10px;
	padding: 10px 15px;
	cursor: pointer;
	&:hover {
		background-color: var(--color-3-b-hover);
	}
}

.remove-passenger-btn {
	margin-left: auto;
	color: red;
	border: 2px solid var(--color-3-b);
	cursor: pointer;
	transition: border-color 0.3s, background-color 0.3s;
	&:hover {
		border-color: var(--color-3-b-hover);
		background-color: var(--color-3-a-hover);
	}
}

@media (max-width: 500px) {
	.passengers {
		display: flex;
		flex-direction: column;
	}
}
</style>