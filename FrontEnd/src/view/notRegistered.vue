<script setup>
defineOptions({
	name: "not-registered"
})
</script>

<template>
	<main>
		<div class="error-div">
			<h1>401</h1>
			<h2>Вы не зарегистрированы</h2>
			<RouterLink to="/register" class="a-link">Зарегистрироваться</RouterLink>
		</div>
	</main>
</template>