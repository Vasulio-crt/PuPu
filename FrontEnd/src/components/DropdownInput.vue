<script setup>
import { ref, onMounted, onUnmounted, computed } from 'vue'

//! Этот компонент был полностью написан ИИ.

const props = defineProps({
	modelValue: {
		type: String,
		default: ''
	},
	placeholder: {
		type: String,
		default: 'Введите текст'
	},
	options: {
		type: Array,
		default: () => []
	},
	filterKey: {
		type: String,
		default: 'name'
	},
	displayKey: {
		type: String,
		default: 'name'
	},
	minChars: {
		type: Number,
		default: 0
	}
})

const emit = defineEmits(['update:modelValue', 'select'])

const isOpen = ref(false)
const inputRef = ref(null)
const dropdownRef = ref(null)

const filteredOptions = computed(() => {
	if (!props.modelValue || props.modelValue.length < props.minChars) {
		return props.options
	}
	
	const searchTerm = props.modelValue.toLowerCase()
	return props.options.filter(option => {
		const value = typeof option === 'string' 
			? option 
			: option[props.filterKey]
		return value.toLowerCase().includes(searchTerm)
	})
})

function handleInput(event) {
	emit('update:modelValue', event.target.value)
	isOpen.value = true
}

function selectOption(option) {
	const value = typeof option === 'string' ? option : option[props.displayKey]
	emit('update:modelValue', value)
	emit('select', option)
	isOpen.value = false
}

function handleFocus() {
	isOpen.value = true
}

function handleClickOutside(event) {
	if (
		dropdownRef.value && 
		!dropdownRef.value.contains(event.target) &&
		inputRef.value &&
		!inputRef.value.contains(event.target)
	) {
		isOpen.value = false
	}
}

onMounted(() => {
	document.addEventListener('click', handleClickOutside)
})

onUnmounted(() => {
	document.removeEventListener('click', handleClickOutside)
})
</script>

<template>
	<div class="dropdown-container">
		<input ref="inputRef"
			:value="modelValue"
			@input="handleInput"
			@focus="handleFocus"
			type="text"
			:placeholder="placeholder"
			class="interactive-elem dropdown-input"
			autocomplete="off"/>
		
		<div v-if="isOpen && filteredOptions.length > 0" 
			ref="dropdownRef"
			class="dropdown-menu">
			<div v-for="(option, index) in filteredOptions"
				:key="index"
				@click="selectOption(option)"
				class="dropdown-menu-item">
				<slot name="option" :option="option">
					{{ typeof option === 'string' ? option : option[displayKey] }}
				</slot>
			</div>
		</div>
		
		<div v-if="isOpen && filteredOptions.length === 0 && modelValue.length >= minChars"
			ref="dropdownRef"
			class="dropdown-menu">
			<div class="dropdown-menu-item-empty">
				<slot name="empty">
					Ничего не найдено
				</slot>
			</div>
		</div>
	</div>
</template>

<style scoped>
.dropdown-container {
	position: relative;
	width: 100%;
}

.dropdown-input {
	width: 100%;
}

.dropdown-menu {
	position: absolute;
	top: calc(100% + 4px);
	left: 0;
	right: 0;
	max-height: 300px;
	overflow-y: auto;
	border: 2px solid var(--color-3-b);
	background-color: var(--color-3-a);
	border-radius: 8px;
	z-index: 1000;
	box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.dropdown-menu-item {
	padding: 12px 16px;
	cursor: pointer;
	transition: background-color 0.2s ease;
}

.dropdown-menu-item:hover {
	background-color: var(--color-3-b-hover);
}

.dropdown-menu-item-empty {
	padding: 12px 16px;
	color: var(--color-text-secondary, #999);
	text-align: center;
}
</style>