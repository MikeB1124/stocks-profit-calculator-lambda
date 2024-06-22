package controllers

import (
	"context"
	"log"

	"github.com/MikeB1124/stocks-profit-calculator-lambda/db"
	"github.com/aws/aws-lambda-go/events"
)

type Response struct {
	StatusCode int    `json:"statusCode"`
	Message    string `json:"message"`
}

func CalulateTradeProfits(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Complete Trade For Expired or Cancelled Orders
	expireUpdateResult, err := db.UpdateAllExpiredOrders()
	if err != nil {
		log.Printf("Error updating expired orders: %s", err.Error())
		return createResponse(Response{Message: err.Error(), StatusCode: 500})
	}
	log.Printf("Updated %d expired orders", expireUpdateResult.ModifiedCount)

	cancelUpdateResult, err := db.UpdateAllCancelledOrders()
	if err != nil {
		log.Printf("Error updating cancelled orders: %s", err.Error())
		return createResponse(Response{Message: err.Error(), StatusCode: 500})
	}
	log.Printf("Updated %d cancelled orders", cancelUpdateResult.ModifiedCount)

	return createResponse(Response{Message: "OK", StatusCode: 200})
}
