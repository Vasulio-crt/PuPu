<script setup>
import { usePassengerStore } from '@/store/passenger';
import { useAuthStore } from '@/store/auth';
import { ref, computed, onMounted, inject } from 'vue';
import { useRouter } from 'vue-router';
import PassengerForm from '@/components/PassengerForm.vue';

const host = inject('hostBacked')
const store = usePassengerStore()
const storeAuth = useAuthStore()
const vueRouter = useRouter()

const passengersData = ref([])

const initializePassengersData = () => {
	passengersData.value = store.data.passengers.map((passenger, index) => ({
		passengerId: passenger.id,
		passengerType: passenger.name,
		seatNumber: store.selectedSeats[index]?.num || null,
		surname: '',
		name: '',
		patronymic: '',
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
		p.surname.trim() && 
		p.name.trim() && 
		p.birthDate && 
		p.passportSeries.trim() && 
		p.passportNumber.trim()
	)
})

const handleSubmit = async () => {
	if (isFormValid.value) {
		console.log('Данные пассажиров:', passengersData.value)
		const req_data = []
		for (const passenger of passengersData.value) {
			req_data.push({
				route_id: store.route.id,
				price: store.route.price,
				passenger_type: passenger.passengerType,
				seat_number: passenger.seatNumber,
				carriage_id: store.carriageID,
				name: passenger.name,
				surname: passenger.surname,
				patronymic: passenger.patronymic,
				birth_date: passenger.birthDate,
				passport_series: Number(passenger.passportSeries),
				passport_number: Number(passenger.passportNumber)
			})
		}
		try {
			const response = await fetch(`${host}/api/buyTicket`, {
				method: 'POST', body: JSON.stringify(req_data),
				headers: {'Login': storeAuth.user.login, 'Content-Type': 'application/json'}, 
			})
			if (response.ok) {
				vueRouter.push({ name: 'show-tickets' })
			} else if (response.status === 404) {
				storeAuth.logout()
				vueRouter.push({name: "not-registered"})
			}
		} catch (e) {
			console.error('Ошибка:', e)
		}
	} else {
		alert('Пожалуйста, заполните все обязательные поля')
	}
}

onMounted(() => {
	if (store.data.passengers.length === 0) {
		vueRouter.push({ name: 'home' })
		return
	}

	if (store.selectedSeats.length !== store.data.passengers.length) {
		vueRouter.push({ name: 'seat-selection' })
		return
	}

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
				<button @click="handleSubmit" :disabled="!isFormValid" class="interactive-elem">Получить билеты</button>
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