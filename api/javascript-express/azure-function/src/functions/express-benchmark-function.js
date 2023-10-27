const { app } = require("@azure/functions");

async function benchmark(request, context) {
    context.log(`Http function processed request for url "${request.url}"`);

    return {
        jsonBody: {
            message: "Simple Azure Functions benchmark"
        }
    }
};

app.http('benchmark', {
    route: "benchmark",
    methods: ['GET'],
    authLevel: 'anonymous',
    handler: benchmark
});

module.exports = benchmark
