const express = require('express');
const app = express();
const port = 3000;

app.get('/benchmark', (req, res) => {
    res.send('Simple ExpressJS Benchmark');
})

app.listen(port, () => {
    console.log("Simple ExpressJS Benchmark running on port " + port);
})