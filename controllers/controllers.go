package controllers

import (
	"context"
	"log"

	stockslambdautils "github.com/MikeB1124/stocks-lambda-utils"
	"github.com/MikeB1124/stocks-profit-calculator-lambda/clients"
	"github.com/aws/aws-lambda-go/events"
)

func CalulateTradeProfits(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Complete Trade For Expired or Cancelled Orders
	expireUpdateResult, err := clients.MongoClient.UpdateAllExpiredOrders()
	if err != nil {
		log.Printf("Error updating expired orders: %s", err.Error())
		return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
	}
	log.Printf("Updated %d expired orders", expireUpdateResult.ModifiedCount)

	cancelUpdateResult, err := clients.MongoClient.UpdateAllCancelledOrders()
	if err != nil {
		log.Printf("Error updating cancelled orders: %s", err.Error())
		return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
	}
	log.Printf("Updated %d cancelled orders", cancelUpdateResult.ModifiedCount)

	return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: "OK", StatusCode: 200})
}
