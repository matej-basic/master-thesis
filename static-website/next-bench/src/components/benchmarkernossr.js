
const BenchmarkerNoSSR = () => {
  const currentTime = new Date();
  const timer = currentTime - window.performance.timing.requestStart;
    return (
      <>
        <div className="timer-header">
            Simple NextJS Benchmark with SSR disabled
        </div>
      <div className='timer' id='timer'>Time from the start of the request to when the page loaded: {timer} ms</div>
      </>
    )
}

export default BenchmarkerNoSSR;