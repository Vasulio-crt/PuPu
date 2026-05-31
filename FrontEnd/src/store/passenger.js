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
		},
		getTrainFlights() {
			console.log(this.data);
			
			// const host = inject('hostBacked');
			// const params = new URLSearchParams({
			// 	from: this.data.from,
			// 	to: this.data.to,
			// 	date: this.data.date
			// });

			// fetch(`${host}/api/getRoutes?${params.toString()}`, {
			// 	method: 'GET'
			// }).then(res => res.text()).then((data) => console.log('Запрос рейсов прошел успешно', data))	
		}
	}
})