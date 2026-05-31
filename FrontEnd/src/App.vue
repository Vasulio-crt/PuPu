<script setup>
import { provide } from 'vue';
import HeaderComp from './components/HeaderComp.vue';

const { protocol, hostname } = window.location;
const isLocalOrIp = hostname === 'localhost' || /^(?:[0-9]{1,3}\.){3}[0-9]{1,3}$/.test(hostname);
const hostBack = `${protocol}//${hostname}${isLocalOrIp ? ':8080' : ''}`;
provide('hostBacked', hostBack)
</script>

<template>
<HeaderComp />
<RouterView />
</template>
