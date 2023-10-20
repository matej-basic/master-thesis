<script>
	const time = new Date();
	const requestTime = window.performance.timing.requestStart;
	const fullTime = time - requestTime;
</script>

<main>
	<div class='timer-header'>
		Simple Svelte Benchmark
	</div>
	<div class='timer'>Time from the start of the request to when the page loaded: {fullTime} ms</div>
</main>