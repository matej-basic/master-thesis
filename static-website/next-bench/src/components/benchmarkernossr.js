export const BenchmarkerNoSSR = () => {
  return (
  <div>
    <div className="timer-header">
            Simple NextJS Benchmark with SSR disabled
    </div>
    <div className="timer flex-container">
      Time from the start of the request to when the page loaded: 
      <div id='timer' className="flex-child"></div>
      <div className="flex-child">ms</div>
    </div>
    <script>
        {`
      const currentTime = new Date();
      const fullTime = currentTime - window.performance.timing.requestStart;
      document.getElementById(\`timer\`).innerHTML = fullTime;
        `}
      </script>
  </div>
  )
}