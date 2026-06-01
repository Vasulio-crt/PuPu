import { defineStore } from "pinia";

export const usePassengerStore = defineStore('passenger', {
	state: () => ({
		data: {
			from: '',
			to: '',
			date: '',
			passengers: []
		},
		routes: []
	}),
	actions: {
		deletePassenger(id) {
			const index = this.data.passengers.findIndex(p => p.id === id)
			if (index !== -1) {
				this.data.passengers.splice(index, 1)
			}
		},
		getToday() {
			const todayString = new Date().toISOString().split('T')[0]
			return todayString
		}
	}
})