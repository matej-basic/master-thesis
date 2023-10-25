<script setup>
    const time = new Date();
    const requestStart = window.performance.timing.requestStart;
    const fullTime = time - requestStart;
</script>

<template>
    <div class="timer-header">Simple Vue.JS benchmark</div>
    <div class="timer">Time from the start of the request to when the page loaded: {{ fullTime }} ms</div>
</template>
<style scoped>
.timer-header {
    font-size: 48px;
    text-align: center;
    margin-bottom: 30px;
}

.timer {
    margin-left: 25px;
}
</style>