import { defineStore } from "pinia";
import { inject } from "vue";

export const usePassengerStore = defineStore('passenger', {
	state: () => ({
		data: {
			from: '',
			to: '',
			date: '',
			passengers: []
		},
		user: {
			Name: '',
			Surname: ''
		}
	}),
	actions: {
		deletePassenger(id) {
			const index = this.data.passengers.findIndex(p => p.id === id)
			if (index !== -1) {
				this.data.passengers.splice(index, 1)
			}
		}
	}
})