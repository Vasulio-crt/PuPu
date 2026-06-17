<script setup>
import { useAuthStore } from '@/store/auth'
import { inject, onMounted, ref } from 'vue'
import { useRouter } from 'vue-router'

const host = inject('hostBacked')
const storeAuth = useAuthStore()
const vueRouter = useRouter()
const Tickets = ref([])

async function getTickets() {
	try {
		const response = await fetch(`${host}/api/getTicket`, {
			headers: {'Login': storeAuth.user.login}, 
		})
		if (response.ok) {
			Tickets.value = await response.json()
		} else if (response.status === 404) {
			storeAuth.logout()
			vueRouter.push({name: "not-registered"})
		}
	} catch (e) {
		console.error('Ошибка:', e)
	}
}

const formatDateTime = (dateString) => {
	const date = new Date(dateString)
	return date.toLocaleString('ru-RU', { 
		day: '2-digit', 
		month: '2-digit', 
		year: 'numeric',
		hour: '2-digit',
		minute: '2-digit'
	})
}

const formatDate = (dateString) => {
	const date = new Date(dateString)
	return date.toLocaleDateString('ru-RU')
}

onMounted(() => {
	getTickets()
})
</script>

<template>
	<main>
		<div class="container">
			<h1>Мои билеты</h1>

			<div v-if="Tickets.length === 0" class="empty">
				<p>У вас пока нет билетов</p>
			</div>

			<div v-else class="tickets">
				<div v-for="ticket in Tickets" :key="ticket.ticket_id" class="ticket">
					
					<div class="ticket-top">
						<h2>Билет №{{ ticket.ticket_id }}</h2>
						<div class="price">{{ ticket.price }} ₽</div>
					</div>

					<div class="route">
						<div class="point">
							<strong>{{ ticket.station_from }}</strong>
							<div>{{ formatDateTime(ticket.sending) }}</div>
						</div>
						<div class="arrow">-></div>
						<div class="point">
							<strong>{{ ticket.station_to }}</strong>
							<div>{{ formatDateTime(ticket.arrival) }}</div>
						</div>
					</div>

					<div class="details">
						<div>Вагон: <strong>{{ ticket.carriage_id }}</strong></div>
						<div>Место: <strong>{{ ticket.seat_number }}</strong></div>
						<div>{{ ticket.passenger_type }}</div>
					</div>

					<div class="passenger">
						<div class="name">{{ ticket.surname }} {{ ticket.name }} {{ ticket.patronymic }}</div>
						<div>Дата рождения: {{ formatDate(ticket.birth_date) }}</div>
						<div>Паспорт: {{ ticket.pass_series }} {{ ticket.pass_number }}</div>
					</div>

				</div>
			</div>
		</div>
	</main>
</template>

<style scoped>
.container {
	max-width: 750px;
}

h1 {
	margin-bottom: 1rem;
}

.empty {
	background: #1e1e1e;
	padding: 40px;
	text-align: center;
	border-radius: 8px;
	color: #888;
}

.tickets {
	display: flex;
	flex-wrap: wrap;
	gap: 20px;
}

.ticket {
	background: #1e1e1e;
	padding: 20px;
	border-radius: 8px;
	border: 1px solid #333;
}

.ticket-top {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: 20px;
	padding-bottom: 15px;
	border-bottom: 1px solid #333;
}

.ticket-top h2 {
	color: #fff;
	font-size: 18px;
	margin: 0;
}

.price {
	color: #4caf50;
	font-size: 24px;
	font-weight: bold;
}

.route {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: 20px;
	gap: 15px;
}

.point {
	flex: 1;
	color: #b0b0b0;
}

.point strong {
	color: #fff;
	font-size: 18px;
	display: block;
	margin-bottom: 5px;
}

.arrow {
	color: #4caf50;
	font-size: 24px;
}

.details {
	display: flex;
	gap: 20px;
	margin-bottom: 15px;
	padding: 15px;
	background: #2a2a2a;
	border-radius: 6px;
	color: #b0b0b0;
}

.details strong {
	color: #fff;
}

.passenger {
	padding-top: 15px;
	border-top: 1px solid #333;
	color: #b0b0b0;
}

.passenger .name {
	color: #fff;
	font-size: 16px;
	font-weight: bold;
	margin-bottom: 8px;
}

@media (max-width: 600px) {
	.route {
		flex-direction: column;
		text-align: center;
	}
	
	.details {
		flex-direction: column;
		gap: 10px;
	}
	
	.ticket-top {
		flex-direction: column;
		gap: 10px;
	}
}
</style>