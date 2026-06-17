<script setup>
import { usePassengerStore } from '@/store/passenger';
import { ref, computed, onMounted } from 'vue';
import { useRouter } from 'vue-router';
import PassengerForm from '@/components/PassengerForm.vue';

const store = usePassengerStore()
const router = useRouter()

const passengersData = ref([])

const initializePassengersData = () => {
	passengersData.value = store.data.passengers.map((passenger, index) => ({
		passengerId: passenger.id,
		passengerType: passenger.name,
		seatNumber: store.selectedSeats[index]?.num || null,
		lastName: '',
		firstName: '',
		middleName: '',
		birthDate: '',
		passportSeries: '',
		passportNumber: ''
	}))
}

const handleUpdate = (passengerId, field, value) => {
	const passenger = passengersData.value.find(p => p.passengerId === passengerId)
	if (passenger) {
		passenger[field] = value
	}
}

const getPassengerData = (passengerId) => {
	return passengersData.value.find(p => p.passengerId === passengerId) || {}
}

const isFormValid = computed(() => {
	return passengersData.value.every(p => 
		p.lastName.trim() && 
		p.firstName.trim() && 
		p.birthDate && 
		p.passportSeries.trim() && 
		p.passportNumber.trim()
	)
})

const handleSubmit = () => {
	if (isFormValid.value) {
		console.log('Данные пассажиров:', passengersData.value)
		
		// Здесь можно отправить данные на сервер
		// await sendPassengerData(passengersData.value)
		
		// Переход на следующую страницу
		//router.push({ name: 'confirmation' })
	} else {
		alert('Пожалуйста, заполните все обязательные поля')
	}
}

onMounted(() => {
	if (store.data.passengers.length === 0) {
		router.push({ name: 'home' })
		return
	}

	if (store.selectedSeats.length !== store.data.passengers.length) {
		router.push({ name: 'seat-selection' })
		return
	}
	
	// Инициализируем данные пассажиров
	initializePassengersData()
})
</script>

<template>
	<main>
		<div class="container">
			<h1 style="margin-bottom: 0.5rem;">Данные для пассажиров</h1>
			<div class="display-elem interactive-elem">
				<div class="info-item">
					<span class="label">Маршрут:</span>
					<span class="value">{{ store.data.from }} > {{ store.data.to }}</span>
				</div>
				<div class="info-item">
					<span class="label">Дата:</span>
					<span class="value">{{ store.data.date }}</span>
				</div>
				<div class="info-item">
					<span class="label">Места:</span>
					<span class="value">{{ store.getSelectedSeats.join(', ') }}</span>
				</div>
				<div class="info-item">
					<span class="label">Пассажиров:</span>
					<span class="value">{{ store.data.passengers.length }}</span>
				</div>
			</div>
			<PassengerForm v-for="(passenger, index) in store.data.passengers"
				:key="passenger.id"
				:passenger="passenger"
				:passenger-data="getPassengerData(passenger.id)"
				:seat="store.selectedSeats[index]?.num"
				:index="index"
				@update="handleUpdate"
			/>
			<div class="size-display-1" style="margin-bottom: 1rem;">
				<button @click="handleSubmit" :disabled="!isFormValid" class="interactive-elem">Получить билет</button>
			</div>
		</div>
	</main>
</template>

<style scoped>
.container {
	max-width: 1200px;
}

.size-display-1 {
	width: 100%;
}

.display-elem {
	margin-bottom: 0.5rem;
	display: flex;
	gap: 30px;
	flex-wrap: wrap;
}

.info-item {
	display: flex;
	flex-direction: column;
	gap: 5px;
}

.info-item .label {
	color: #888;
	font-size: 14px;
}

.info-item .value {
	color: #ffffff;
	font-weight: 600;
	font-size: 16px;
}
</style>