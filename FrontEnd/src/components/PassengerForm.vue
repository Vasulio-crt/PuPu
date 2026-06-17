<script setup>
const props = defineProps({
	passenger: {
		type: Object,
		required: true
	},
	passengerData: {
		type: Object,
		required: true
	},
	seat: {
		type: Number,
		default: null
	},
	index: {
		type: Number,
		required: true
	}
})

const emit = defineEmits(['update'])

const updateField = (field, value) => {
	emit('update', props.passenger.id, field, value)
}
</script>

<template>
	<div class="interactive-elem">
		<div class="card-header">
			<h3>{{ passenger.name }} {{ index + 1 }}</h3>
			<span class="seat-badge" v-if="seat">
				Место: {{ seat }}
			</span>
		</div>

		<div class="size-display-1">
			<div class="search-row-column">
				<div class="input-group">
					<label>Фамилия <span class="red-text">*</span></label>
					<input
						type="text"
						:value="passengerData.lastName"
						@input="updateField('lastName', $event.target.value)"
						placeholder="Иванов"
						required
						class="interactive-elem"
					>
				</div>
				<div class="input-group">
					<label>Имя <span class="red-text">*</span></label>
					<input
						type="text"
						:value="passengerData.firstName"
						@input="updateField('firstName', $event.target.value)"
						placeholder="Иван"
						required
						class="interactive-elem"
					>
				</div>
			</div>

			<div class="search-row-column">
				<div class="input-group">
					<label>Отчество</label>
					<input 
						type="text" 
						:value="passengerData.middleName"
						@input="updateField('middleName', $event.target.value)"
						placeholder="Иванович"
						class="interactive-elem"
					>
				</div>
			</div>

			<div class="search-row-column">
				<div class="input-group">
					<label>Дата рождения <span class="red-text">*</span></label>
					<input 
						type="date" 
						:value="passengerData.birthDate"
						@input="updateField('birthDate', $event.target.value)"
						:max="new Date().toISOString().split('T')[0]"
						required
						class="interactive-elem"
					>
				</div>
			</div>

			<div class="search-row-column">
				<div class="input-group">
					<label>Серия паспорта <span class="red-text">*</span></label>
					<input 
						type="text" 
						:value="passengerData.passportSeries"
						@input="updateField('passportSeries', $event.target.value)"
						placeholder="1234"
						maxlength="4"
						pattern="[0-9]{4}"
						required
						class="interactive-elem"
					>
				</div>
				<div class="input-group">
					<label>Номер паспорта <span class="red-text">*</span></label>
					<input 
						type="text" 
						:value="passengerData.passportNumber"
						@input="updateField('passportNumber', $event.target.value)"
						placeholder="123456"
						maxlength="6"
						pattern="[0-9]{6}"
						required
						class="interactive-elem"
					>
				</div>
			</div>
		</div>
	</div>
</template>

<style scoped>
.card-header {
	display: flex;
	justify-content: space-between;
	align-items: center;
	margin-bottom: 1rem;
	padding-bottom: 0.7rem;
	border-bottom: 2px solid var(--color-3-b);
}

.seat-badge {
	background: var(--color-2);
	padding: 6px 12px;
	border-radius: 20px;
	font-size: 14px;
	font-weight: 500;
}

.size-display-1 {
	padding-top: 0;
	gap: 10px;
}

@media (max-width: 768px) {
	.card-header {
		flex-direction: column;
		align-items: flex-start;
		gap: 10px;
	}
}
</style>