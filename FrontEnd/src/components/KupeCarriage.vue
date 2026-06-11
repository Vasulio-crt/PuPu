<script setup>
import { computed } from 'vue'

const props = defineProps({
	seats: {
		type: Array,
		required: true,
		default: () => []
	}
})

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
</script>

<template>
	<div class="carriage">
		<div class="room" v-for="chunk in rooms">
			<button class="i1" :disabled="chunk[0]?.occ">{{ chunk[0]?.num }}</button>
			<button class="i2" :disabled="chunk[1]?.occ">{{ chunk[1]?.num }}</button>
			<button class="i3" :disabled="chunk[2]?.occ">{{ chunk[2]?.num }}</button>
			<button class="i4" :disabled="chunk[3]?.occ">{{ chunk[3]?.num }}</button>
		</div>
	</div>
</template>

<style scoped>
@import url('../assets/carriage.css');
</style>