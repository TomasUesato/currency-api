package main

import (
	"context"
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
)

type Rate struct {
	Currency string  `json:"currency"`
	Rate     float64 `json:"rate"`
	Unit     string  `json:"unit"`
}

type RatesResponse struct {
	Base  string `json:"base"`
	Rates []Rate `json:"rates"`
}

type HealthResponse struct {
	Status  string `json:"status"`
	Service string `json:"service"`
}

func handleHealth() (events.APIGatewayV2HTTPResponse, error) {
	body, _ := json.Marshal(HealthResponse{Status: "ok", Service: "currency-api"})
	return response(200, string(body)), nil
}

func handleRates() (events.APIGatewayV2HTTPResponse, error) {
	rates := RatesResponse{
		Base: "ARS",
		Rates: []Rate{
			{Currency: "USD", Rate: 1200.50, Unit: "ARS per USD"},
			{Currency: "EUR", Rate: 1305.75, Unit: "ARS per EUR"},
			{Currency: "BRL", Rate: 215.30, Unit: "ARS per BRL"},
		},
	}
	body, _ := json.Marshal(rates)
	return response(200, string(body)), nil
}

func handler(ctx context.Context, req events.APIGatewayV2HTTPRequest) (events.APIGatewayV2HTTPResponse, error) {
	switch {
	case req.RequestContext.HTTP.Method == "GET" && req.RequestContext.HTTP.Path == "/default/health":
		return handleHealth()
	case req.RequestContext.HTTP.Method == "GET" && req.RequestContext.HTTP.Path == "/default/rates":
		return handleRates()
	default:
		return response(404, `{"error":"route not found"}`), nil
	}
}

func response(statusCode int, body string) events.APIGatewayV2HTTPResponse {
	return events.APIGatewayV2HTTPResponse{
		StatusCode: statusCode,
		Headers:    map[string]string{"Content-Type": "application/json"},
		Body:       body,
	}
}

func main() {
	lambda.Start(handler)
}
