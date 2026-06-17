import { defineStore } from 'pinia'
import { computed, reactive, ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
	const defaultUser = {
		name: '',
		surname: '',
		patronymic: '',
		birth_date: '',
		email: '',
		phone: ''
	}
	const user = reactive(JSON.parse(localStorage.getItem('user')) || defaultUser)
	const isLoggedIn = ref(!!localStorage.getItem('user'))


	const fullName = computed(() => `${user.name} ${user.surname}`)
	const userNotExits = computed(() => {
		if (user.name === '' || user.surname === '') {
			return true
		} else {
			return false
		}
	})

	function setAuthData(userData) {
		Object.assign(user, userData)
		localStorage.setItem('user', JSON.stringify(userData))
		isLoggedIn.value = true
	}

	function logout() {
		Object.assign(user, defaultUser)
		localStorage.removeItem('user')
		isLoggedIn.value = false
	}

	return {
		user,
		isLoggedIn,
		fullName,
		userExits: userNotExits,
		setAuthData,
		logout,
	}
})
