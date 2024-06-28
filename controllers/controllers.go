package controllers

import (
	"context"
	"log"
	"strconv"

	stockslambdautils "github.com/MikeB1124/stocks-lambda-utils/v2"
	"github.com/MikeB1124/stocks-profit-calculator-lambda/configuration"
	"github.com/aws/aws-lambda-go/events"
)

func CalulateTradeProfits(ctx context.Context, event events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {
	// Complete Trade For Expired or Cancelled Orders
	expireUpdateResult, err := configuration.MongoClient.UpdateAllExpiredOrders()
	if err != nil {
		log.Printf("Error updating expired and canceled orders: %s", err.Error())
		return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
	}
	log.Printf("Updated %d expired and canceled orders", expireUpdateResult.ModifiedCount)

	// Get filled trades from DB
	filledTrades, err := configuration.MongoClient.GetFilledTradesFromDB()
	if err != nil {
		log.Printf("Error getting filled trades from DB: %s", err.Error())
		return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
	}

	// Calculate profits for each trade
	var calculatedTrades []stockslambdautils.AlpacaTrade
	for _, trade := range filledTrades {
		entryQty, err := strconv.ParseFloat(trade.Order.FilledQty, 64)
		if err != nil {
			log.Printf("Error parsing entry quantity: %s", err.Error())
			return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
		}
		entryAvgPrice, err := strconv.ParseFloat(trade.Order.FilledAvgPrice, 64)
		if err != nil {
			log.Printf("Error parsing entry average price: %s", err.Error())
			return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
		}

		var exitQty float64
		var exitAvgPrice float64
		var errParse error
		if trade.Order.Legs[0].Status == "filled" {
			exitQty, errParse = strconv.ParseFloat(trade.Order.Legs[0].FilledQty, 64)
			if errParse != nil {
				log.Printf("Error parsing exit quantity: %s", err.Error())
				return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
			}
			exitAvgPrice, errParse = strconv.ParseFloat(trade.Order.Legs[0].FilledAvgPrice, 64)
			if errParse != nil {
				log.Printf("Error parsing exit average price: %s", err.Error())
				return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
			}
		} else {
			exitQty, errParse = strconv.ParseFloat(trade.Order.Legs[1].FilledQty, 64)
			if errParse != nil {
				log.Printf("Error parsing exit quantity: %s", err.Error())
				return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
			}
			exitAvgPrice, errParse = strconv.ParseFloat(trade.Order.Legs[1].FilledAvgPrice, 64)
			if errParse != nil {
				log.Printf("Error parsing exit average price: %s", err.Error())
				return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
			}
		}

		profit := (exitAvgPrice * exitQty) - (entryAvgPrice * entryQty)
		trade.TradeProfit = profit
		trade.TradeCompleted = true
		calculatedTrades = append(calculatedTrades, trade)
	}

	if len(calculatedTrades) > 0 {
		// Bulk update trade profits
		bulkUpdateResults, err := configuration.MongoClient.BulkUpdateTradeProfits(calculatedTrades)
		if err != nil {
			log.Printf("Error updating trade profits: %s", err.Error())
			return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: err.Error(), StatusCode: 500})
		}
		log.Printf("%+v", bulkUpdateResults)

	} else {
		log.Println("No trades to update")
	}
	return stockslambdautils.CreateResponse(stockslambdautils.Response{Message: "OK", StatusCode: 200})
}
