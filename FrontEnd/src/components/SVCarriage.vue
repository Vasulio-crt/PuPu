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
	for (let i = 0; i < 18; i += 2) {
		result.push([
			props.seats[i],
			props.seats[i + 1]
		])
	}
	return result
})
</script>

<template>
	<div class="carriage">
		<div class="room2 SV" v-for="chunk in rooms">
			<button :disabled="chunk[0]?.occ" style="border-right: 1px solid var(--color-4)">{{ chunk[0]?.num }}</button>
			<button :disabled="chunk[1]?.occ" style="border-left: 1px solid var(--color-4)">{{ chunk[1]?.num }}</button>
		</div>
	</div>
</template>

<style scoped>
@import url('../assets/carriage.css');
</style>
