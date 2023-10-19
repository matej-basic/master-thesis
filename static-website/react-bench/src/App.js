import './App.css';

function App() {
  const currentTime = new Date();
  const timer = currentTime - window.performance.timing.requestStart;

  return (
    <div className="App" >
      <div className='timer-header'>
        Simple React Benchmark
      </div>
      <div id='timer'>Time from the start of the request to when the page loaded: {timer} ms</div>
    </div>
  );
}

export default App;
