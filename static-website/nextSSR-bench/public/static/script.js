const currentTime = new Date();
const fullTime = currentTime - window.performance.timing.requestStart;
console.log(fullTime);
document.getElementById("timer").innerHTML = fullTime;