const express = require('express');
const app = express();
const port = process.env.PORT;

app.get('/', (req, res) => {
    res.send('Simple ExpressJS Benchmark');
})

module.exports = app