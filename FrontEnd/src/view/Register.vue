<script setup>
import { inject, onMounted, reactive } from 'vue';
import { RouterLink, useRouter } from 'vue-router';
import { useAuthStore } from '@/store/auth';

const host = inject('hostBacked')
const authStore = useAuthStore()
const router = useRouter()

const registration_data = reactive({
	login: '',
	password: '',
	email: '',
	phone: '',
	birth_date: '',
	name: '',
	surname: ''
})
const errors = reactive({})

function validate() {
	errors.login = !registration_data.login ? 'Логин обязателен' : ''
	errors.password = !registration_data.password ? 'Пароль обязателен' : (registration_data.password.length < 6 ? 'Пароль должен быть не менее 6 символов' : '')
	errors.email = !registration_data.email ? 'Email обязателен' : (!/^\S+@\S+\.\S+$/.test(registration_data.email) ? 'Неверный формат email' : '')
	errors.name = !registration_data.name ? 'Имя обязательно' : ''
	errors.phone = !registration_data.phone ? 'Телефон обязателен' : (!/^((\+7|7|8)+([0-9]){10})$/.test(registration_data.phone.replace(/\s/g, '')) ? 'Неверный формат телефона' : '')
	errors.birth_date = !registration_data.birth_date ? 'Дата рождения обязательна' : (!/^\d{4}-\d{2}-\d{2}$/.test(registration_data.birth_date) ? 'Неверный формат даты (ГГГГ-ММ-ДД)' : '')
	errors.surname = !registration_data.surname ? 'Фамилия обязательна' : ''

	return Object.values(errors).every(error => !error)
}

async function register() {
	if (validate()) {
		try {
			registration_data.phone = registration_data.phone.replace(/\s+/g, '')
			const response = await fetch(`${host}/api/register`, {
				method: 'POST',
				body: JSON.stringify(registration_data),
				headers: {'Content-Type': 'application/json'}
			})
			if (response.ok) {
				authStore.setAuthData(registration_data)
				router.push('/')
			} else {
				const errorText = await response.text()
				errors.login = errorText
			}
		} catch (error) {
			console.error('Ошибка регистрации:', error)
			errors.login = 'Ошибка сервера при регистрации.'
		}
	}
}

onMounted(() => {
	const date = new Date()
	date.setFullYear(date.getFullYear() - 18)
	registration_data.birth_date = new Date(date).toISOString().split('T')[0]
})
</script>

<template>
<main>
	<div class="size-display-1">
		<div class="search-row-column">
			<div class="input-group">
				<label for="Login">Логин <span class="red-text">*</span></label>
				<input v-model="registration_data.login" type="text" placeholder="Логин" id="Login" class="interactive-elem">
			</div>
			<div class="input-group">
				<label for="Password">Пароль <span class="red-text">*</span></label>
				<input v-model="registration_data.password" type="password" placeholder="Пароль" id="Password" class="interactive-elem">
			</div>
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.login }}</p>
			<p class="error-text">{{ errors.password }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<div class="search-row-column">
			<div class="input-group">
				<label for="Email">Email <span class="red-text">*</span></label>
				<input v-model="registration_data.email" type="text" id="Email" placeholder="Email" class="interactive-elem">
			</div>
			<div class="input-group">
				<label for="Phone">Телефон <span class="red-text">*</span></label>
				<input v-model="registration_data.phone" type="text" id="Phone" placeholder="8 999 111 44 66" class="interactive-elem">
			</div>
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.email }}</p>
			<p class="error-text">{{ errors.phone }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<div class="search-row-column">
			<div class="input-group">
				<label for="BirthDate">Дата рождения <span class="red-text">*</span></label>
				<input v-model="registration_data.birth_date" type="date" id="BirthDate" placeholder="Дата рождения" class="interactive-elem">
			</div>
			<div class="input-group">
				<label for="Name">Имя <span class="red-text">*</span></label>
				<input v-model="registration_data.name" type="text" id="Name" placeholder="Имя" class="interactive-elem">
			</div>
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.birth_date }}</p>
			<p class="error-text">{{ errors.name }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<div class="search-row-column">
			<div class="input-group">
				<label for="Surname">Фамилия <span class="red-text">*</span></label>
				<input v-model="registration_data.surname" type="text" id="Surname" placeholder="Фамилия" class="interactive-elem">
			</div>
			<div class="input-group">
				<label for="Patronymic">Отчество</label>
				<input v-model="registration_data.patronymic" type="text" id="Patronymic" placeholder="Отчество" class="interactive-elem">
			</div>
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.surname }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<button @click="register" class="interactive-elem">Зарегистрироваться</button>
	</div>
	<div class="already-there">
		<p>Уже есть аккаунт?</p>
		<RouterLink to="/login">Войти</RouterLink>
	</div>
</main>
</template>