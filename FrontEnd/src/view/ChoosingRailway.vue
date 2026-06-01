<script setup>
import { usePassengerStore } from '@/store/passenger';

defineOptions({
	name: "choosing-railway"
})

const store = usePassengerStore()
</script>

<template>
	<div v-for="route in store.routes">
		<h1>{{ route.from_station }} -> {{ route.to_station }}</h1>
		<h2>{{ route.sending }}</h2>
		<h2>{{ route.arrival }}</h2>
		<p>{{ route.distance }}</p>
	</div>
</template>