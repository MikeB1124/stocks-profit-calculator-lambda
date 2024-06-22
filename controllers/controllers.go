package controllers

import (
	"context"

	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func CalulateTradeProfits(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	return createResponse(Response{Message: "OK", StatusCode: 200})
}
