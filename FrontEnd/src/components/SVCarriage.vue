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
	for (let i = 0; i < 18; i += 2) {
		result.push([
			props.seats[i],
			props.seats[i + 1]
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
		<div class="room2 SV" v-for="chunk in rooms">
			<button :disabled="chunk[0]?.occ" @click="toggleSeat(chunk[0])" :class="{selected: isSeatSelected(chunk[0])}" style="border-right: 1px solid var(--color-4)">{{ chunk[0]?.num }}</button>
			<button :disabled="chunk[1]?.occ" @click="toggleSeat(chunk[1])" :class="{selected: isSeatSelected(chunk[1])}" style="border-left: 1px solid var(--color-4)">{{ chunk[1]?.num }}</button>
		</div>
	</div>
</template>

<style scoped>
@import url('../assets/carriage.css');
</style>
