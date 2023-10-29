# Master thesis

This repository contains everything regarding my master thesis.  
The thesis is titled "Comparing price to performance ratio of Serverless services in public cloud".

The premise of this thesis is to explore the world of serverless services and appropriate programming languages, frameworks and other tools used at the time of writting this thesis (Summer / Fall 2023).  
This is not a one size fits all solution, but I will give insight on how I felt using each platform and tool from a standpoint of a Systems engineer / DevOps engineer. 

## Curl testing used

Usage: curl -w "@curl-format.txt" -o /dev/null -s "https://link"  

## Serverless services tested

Below is the list of serverless services and providers included in this test.  
Services and providers are separated into 4 categories based on the usage of serverless services.

### Deploying a static website

  1. Vercel
  2. Netlify
  3. GitHub Pages
  4. CloudFlare
  5. AWS
     * Simple Storage Service
  6. Azure
     * Blob Storage

### Deploying an API 

  1. Railway.app
  2. Render.com
  3. Fly.io
  4. AWS  
     * Lambda  
  5. Azure  
     * Azure Functions  
     * Azure App Services  

### Deploying a databse

  1. Railway
  2. Planetscale
  3. CockroachDB
  4. Supabase
  5. Atlas
  6. AWS  
     * DynamoDB  
     * Aurora Serverless  
  7. GCP  
     * Cloud SQL  
  8. Azure  
     * Azure SQL Database  
     * Azure Cosmos DB  

### Deploying object storage

  1. Contabo
  2. Bunny.net
  3. Fybe
  4. AWS S3
  5. IBM Cloud object storage
  6. Google cloud storage
  7. Dell object storage
  8. Azure blob storage


## Frameworks and languages tested

Below is the list of programming languages and frameworks used in this test.  
Languages and frameworks are categorized the same way as services.

### Deploying a static website

  1. jQuery
  2. ReactJS
  3. AngularJS
  4. Vue.JS
  5. Svelte

### Deploying an API

  1. Flask
  2. ExpressJS
  3. Gin
  4. Fiber
  5. Axum


## Links for deployed services

### APIs

  1. Flask  
    - [Railway.app](https://flask-railway.matejbasic.com/benchmark)  
    - [Render.com](https://flask-render.matejbasic.com/benchmark)  
    - [Fly.io](https://flask-fly.matejbasic.com/benchmark)  
    - [Azure App Service](https://flask-azure.matejbasic.com/benchmark)  
    - [Azure Function](https://flask-function.matejbasic.com/benchmark)  
    - [AWS Lambda](https://flask-lambda.matejbasic.com/benchmark)  
  2. ExpressJS  
    - [Railway.app](https://express-railway.matejbasic.com/benchmark)  
    - [Render.com](https://express-render.matejbasic.com/benchmark)  
    - [Fly.io](https://express-fly.matejbasic.com/benchmark)  
    - [Azure App Service](https://express-azure.matejbasic.com/benchmark)  
    - [Azure Function](https://express-function.matejbasic.com/benchmark)  
    - [AWS Lambda](https://express-lambda.matejbasic.com/benchmark)  
  3. Gin  
    - [Railway.app](https://gin-railway.matejbasic.com/benchmark)  
    - [Render.com](https://gin-render.matejbasic.com/benchmark)  
    - [Fly.io](https://gin-fly.matejbasic.com/benchmark)  
    - [Azure App Service](https://gin-azure.matejbasic.com/benchmark)  
    - [Azure Function](https://gin-function.matejbasic.com/benchmark)  
    - [AWS Lambda](https://gin-lambda.matejbasic.com/benchmark)  
  4. Fiber  
    - [Railway.app](https://fiber-railway.matejbasic.com/benchmark)  
    - [Render.com](https://fiber-render.matejbasic.com/benchmark)  
    - [Fly.io](https://fiber-fly.matejbasic.com/benchmark)  
    - [Azure App Service](https://fiber-azure.matejbasic.com/benchmark)  
    - [Azure Function](https://fiber-function.matejbasic.com/benchmark)  
    - [AWS Lambda](https://fiber-lambda.matejbasic.com/benchmark)
  5. Axum  
    - [Railway.app](https://axum-railway.matejbasic.com/benchmark)  
    - [Render.com](https://axum-render.matejbasic.com/benchmark)  
    - [Fly.io](https://axum-fly.matejbasic.com/benchmark)  
    - [Azure App Service](https://axum-azure.matejbasic.com/benchmark)  
    - [Azure Function](https://axum-function.matejbasic.com/benchmark)  
    - [AWS Lambda](https://axum-lambda/matejbasic.com/benchmark)  