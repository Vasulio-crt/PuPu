<script setup>
import { usePassengerStore } from '@/store/passenger'
import { computed } from 'vue'

const props = defineProps({
	seats: {
		type: Array,
		required: true,
		default: () => []
	}
})

const store = usePassengerStore()

const rooms = computed(() => {
	const result = []
	for (let i = 0; i < 36; i += 4) {
		result.push([
			props.seats[i],
			props.seats[i + 1],
			props.seats[i + 2],
			props.seats[i + 3]
		])
	}
	return result
})

function toggleSeat(seat) {
	if (!seat || seat.occ) return
	
	if (store.isSeatSelected(seat.num)) {
		store.deselectSeat(seat.num)
	} else {
		if (store.canSelectMoreSeats) {
			store.selectSeat(seat)
		}
	}
}

function isSeatSelected(seat) {
	return seat && store.isSeatSelected(seat.num)
}
</script>

<template>
	<div class="carriage">
		<div class="room" v-for="chunk in rooms">
			<button class="i1" :disabled="chunk[0]?.occ" @click="toggleSeat(chunk[0])" :class="{selected: isSeatSelected(chunk[0])}">{{ chunk[0]?.num }}</button>
			<button class="i2" :disabled="chunk[1]?.occ" @click="toggleSeat(chunk[1])" :class="{selected: isSeatSelected(chunk[1])}">{{ chunk[1]?.num }}</button>
			<button class="i3" :disabled="chunk[2]?.occ" @click="toggleSeat(chunk[2])" :class="{selected: isSeatSelected(chunk[2])}">{{ chunk[2]?.num }}</button>
			<button class="i4" :disabled="chunk[3]?.occ" @click="toggleSeat(chunk[3])" :class="{selected: isSeatSelected(chunk[3])}">{{ chunk[3]?.num }}</button>
		</div>
	</div>
</template>

<style scoped>
@import url('../assets/carriage.css');
</style>