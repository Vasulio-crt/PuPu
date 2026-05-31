<script setup>
import { inject, reactive } from 'vue';
import { useAuthStore } from '@/store/auth';
import { useRouter } from 'vue-router';

defineOptions({
	name: 'login'
})

const host = inject('hostBacked')
const authStore = useAuthStore()
const router = useRouter()

const login_data = reactive({
	login: '',
	password: '',
})
const errors = reactive({})

function validate() {
	errors.login = !login_data.login ? 'Логин обязателен' : ''
	errors.password = !login_data.password ? 'Пароль обязателен' : ''
	return Object.values(errors).every(error => !error)
}

async function login() {
	if (validate()) {
		try {
			const response = await fetch(`${host}/api/login`, {
				method: 'POST',
				body: JSON.stringify(login_data),
				headers: {'Content-Type': 'application/json'}
			})

			if (response.ok) {
				const data = await response.text()
				const [name, surname] = data.split(' ')
				authStore.setAuthData(name, surname)
				router.push('/')
			} else {
				errors.password = 'Неверный логин или пароль'
			}
		} catch (error) {
			console.error('Ошибка входа:', error)
			errors.password = 'Ошибка сервера при попытке входа.'
		}
	}
}
</script>

<template>
<main>
	<div class="size-display-1">
		<div class="input-group">
			<label for="Login">Логин</label>
			<input v-model="login_data.login" type="text" placeholder="Логин" id="Login" class="interactive-elem">
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.login }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<div class="input-group">
			<label for="Password">Пароль</label>
			<input v-model="login_data.password" type="password" placeholder="Пароль" id="Password" class="interactive-elem">
		</div>
		<div class="error-row">
			<p class="error-text">{{ errors.password }}</p>
		</div>
	</div>
	<div class="size-display-1">
		<button @click="login" class="interactive-elem">Войти</button>
	</div>
	<div class="already-there">
		<p>Нет аккаунта?</p>
		<RouterLink to="/register">Зарегистрироваться</RouterLink>
	</div>
</main>
</template>