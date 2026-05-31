import { defineStore } from 'pinia'
import { computed, reactive, ref } from 'vue'

export const useAuthStore = defineStore('auth', () => {
	const user = reactive({
		name: localStorage.getItem('user.name') || '',
		surname: localStorage.getItem('user.surname') || '',
	})
	const isLoggedIn = ref(!!localStorage.getItem('user.name'))


	const fullName = computed(() => `${user.name} ${user.surname}`)

	function setAuthData(name, surname) {
		user.name = name
		user.surname = surname
		localStorage.setItem('user.name', name)
		localStorage.setItem('user.surname', surname)
		isLoggedIn.value = true
	}

	function logout() {
		user.name = ''
		user.surname = ''
		localStorage.removeItem('user.name')
		localStorage.removeItem('user.surname')
		isLoggedIn.value = false
	}

	return {
		user,
		isLoggedIn,
		fullName,
		setAuthData,
		logout,
	}
})
