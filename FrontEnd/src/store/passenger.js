import { defineStore } from "pinia"

export const usePassengerStore = defineStore('passenger', {
	state: () => ({
		/**
		 * @type {Object}
		 * @property {string} from
		 * @property {string} to
		 * @property {string} date
		 * @property {Array} passengers
		*/
		data: {
			from: '',
			to: '',
			date: '',
			passengers: []
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
			console.log(this.data)			
		}
	}
})