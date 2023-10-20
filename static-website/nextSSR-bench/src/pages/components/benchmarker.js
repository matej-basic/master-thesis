
const Benchmarker = ({props}) => {
    const currentTime = new Date();
    const fullTime = currentTime - Date.parse(props.message);
    console.log(props.message);
      return (
          <div>
              <div className="timer-header">Simple NextJS Benchmark with SSR Enabled</div>
              <div className="timer" suppressHydrationWarning={true}>Time from the start of the request to when the page loaded: {fullTime} ms</div>
          </div>
      )
  }
  
  export default Benchmarker;