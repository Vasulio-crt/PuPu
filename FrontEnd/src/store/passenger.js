import { defineStore } from "pinia";

export const usePassengerStore = defineStore('passenger', {
	state: () => ({
		data: {
			from: '',
			to: '',
			date: '',
			passengers: []
		},
		route: {},
		selectedSeats: []
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
		},
		selectSeat(seat) {
			if (this.selectedSeats.length < this.data.passengers.length) {
				this.selectedSeats.push(seat)
			}
		},
		deselectSeat(seatNum) {
			const index = this.selectedSeats.findIndex(s => s.num === seatNum)
			if (index !== -1) {
				this.selectedSeats.splice(index, 1)
			}
		},
		clearSelectedSeats() {
			this.selectedSeats = []
		},
		isSeatSelected(seatNum) {
			return this.selectedSeats.some(s => s.num === seatNum)
		}
	},
	getters: {
		canSelectMoreSeats: (state) => {
			return state.selectedSeats.length < state.data.passengers.length
		},
		remainingSeatsToSelect: (state) => {
			return state.data.passengers.length - state.selectedSeats.length
		},
		getSelectedSeats: (state) => {
			return state.selectedSeats.map(item => item.num)
		}
	}
})