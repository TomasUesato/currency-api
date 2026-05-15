# Go Exchange Rates API (AWS Lambda)

A simple serverless API written in Go and deployed to AWS Lambda.  
The API returns mock JSON exchange rate data and is exposed through AWS API Gateway.

This project was built as a lightweight backend exercise to gain hands-on experience with:

- Go backend development
- AWS Lambda
- API Gateway
- Serverless deployments
- JSON APIs
- Debugging AWS event payloads

---

## Overview

The API exposes a simple endpoint that returns mock exchange rates.

Example response:

```json
{
  "base": "USD",
  "rates": {
    "ARS": 1200.5,
    "EUR": 0.92,
    "GBP": 0.78
  }
}
```

---

## Tech Stack

- Go
- AWS Lambda
- AWS API Gateway (HTTP API)
- JSON REST API

---

## Why AWS Lambda?

AWS Lambda is a **serverless compute service**.

Instead of managing a traditional server:

- no VM setup
- no infrastructure maintenance
- no manual scaling

AWS runs your code only when requests arrive.

In this project:

1. A request hits API Gateway
2. API Gateway triggers the Lambda function
3. The Go handler executes
4. The function returns JSON
5. API Gateway sends the HTTP response back to the client

This architecture is commonly used for:

- lightweight APIs
- microservices
- event-driven systems
- backend utilities
- prototypes and internal tools

---

## Project Structure

```bash
.
├── main.go
├── go.mod
├── go.sum
└── README.md
```

---

## Running Locally

Install dependencies:

```bash
go mod tidy
```

Run locally:

```bash
go run main.go
```

---

## Deploying to AWS Lambda

Build the binary for Linux:

```bash
GOOS=linux GOARCH=amd64 go build -o bootstrap main.go
```

Zip the executable:

```bash
zip function.zip bootstrap
```

Upload the ZIP to AWS Lambda and configure:

- Runtime: `provided.al2023`
- Architecture: `x86_64`

This project uses the custom Amazon Linux 2023 runtime for AWS Lambda.

Connect the Lambda function to API Gateway.

---

## Example Endpoint

```http
GET /rates
```

Response:

```json
{
  "base": "USD",
  "rates": {
    "ARS": 1200.5,
    "EUR": 0.92,
    "GBP": 0.78
  }
}
```

---

## AWS Gotcha I Ran Into

One interesting issue during development was debugging the mismatch between:

- API Gateway HTTP API (V2)
- the older REST API Lambda event format

The request path and event structure differ between versions, which caused routing issues initially.

After inspecting the Lambda event payloads and adjusting the handler logic, the API worked correctly.

This was a useful real-world AWS debugging experience because it's a common issue engineers encounter when working with Lambda integrations.

---

## What I Learned

- Building APIs in Go
- Returning structured JSON responses
- Deploying Go binaries to AWS Lambda
- Understanding API Gateway request flow
- Working with serverless architecture
- Debugging AWS event payload differences
- Packaging and deploying Lambda functions manually

---

## Future Improvements

Possible next steps:

- integrate a real exchange rates provider
- add environment variables
- add unit tests
- use Infrastructure as Code (Terraform or AWS SAM)
- add request logging and monitoring
- support multiple routes
