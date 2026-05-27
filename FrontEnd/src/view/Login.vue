<script setup>
import { inject, reactive } from 'vue';

defineOptions({
	name: 'login'
})

const host = inject('hostBacked')

const registration_data = reactive({
	login: '',
	password: '',
	email: '',
	name: '',
	surname: ''
})
const errors = reactive({})

function validate() {
	errors.login = !registration_data.login ? 'Логин обязателен' : ''
	errors.password = !registration_data.password ? 'Пароль обязателен' : (registration_data.password.length < 6 ? 'Пароль должен быть не менее 6 символов' : '')
	errors.email = !registration_data.email ? 'Email обязателен' : (!/^\S+@\S+\.\S+$/.test(registration_data.email) ? 'Неверный формат email' : '')
	errors.name = !registration_data.name ? 'Имя обязательно' : ''
	errors.surname = !registration_data.surname ? 'Фамилия обязательна' : ''

	return Object.values(errors).every(error => !error)
}

function register() {
	if (validate()) {
		fetch(`${host}/api/register`, {
			method: 'POST',
			body: JSON.stringify(registration_data),
			headers: {'Content-Type': 'application/json'}
		}).then(res => res.text()).then((data) => console.log('Регистрация прошла успешно', data))
	}
}
</script>

<template>
<main>
	<div class="size-display-1">
		<div class="text-row">
			<h2>Логин <span class="red-text">*</span></h2>
			<h2>Пароль <span class="red-text">*</span></h2>
		</div>
		<div class="search-row">
			<input v-model="registration_data.login" type="text" placeholder="Логин" class="interactive-elem">
			<input v-model="registration_data.password" type="password" placeholder="Пароль" class="interactive-elem">
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.login }}</p>
			<p class="error-text">{{ errors.password }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<div class="text-row">
			<h2>Email <span class="red-text">*</span></h2>
			<h2>Имя <span class="red-text">*</span></h2>
		</div>
		<div class="search-row">
			<input v-model="registration_data.email" type="text" placeholder="Email" class="interactive-elem">
			<input v-model="registration_data.name" type="text" placeholder="Имя" class="interactive-elem">
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.email }}</p>
			<p class="error-text">{{ errors.name }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<div class="text-row">
			<h2>Фамилия <span class="red-text">*</span></h2>
			<h2>Отчество</h2>
		</div>
		<div class="search-row">
			<input v-model="registration_data.surname" type="text" placeholder="Фамилия" class="interactive-elem">
			<input v-model="registration_data.patronymic" type="text" placeholder="Отчество" class="interactive-elem">
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.surname }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<button @click="register" class="interactive-elem">Зарегистрироваться</button>
	</div>
</main>
</template>

<style scoped>
.text-row {
	display: flex;
	gap: 12px;
	& > h2 {
		width: 100%;
	}
}

.error-row {
	display: flex;
	gap: 12px;
	margin-top: 4px;
}

.error-text {
	color: #E31B1B;
	width: 100%;
	min-height: 1em;
}
</style>