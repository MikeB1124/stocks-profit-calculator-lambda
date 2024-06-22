package main

import (
	"github.com/MikeB1124/stocks-profit-calculator-lambda/controllers"
	"github.com/aquasecurity/lmdrouter"
	"github.com/aws/aws-lambda-go/lambda"
)

var router *lmdrouter.Router

func init() {
	router = lmdrouter.NewRouter("")
	router.Route("POST", "/sync/profit", controllers.CalulateTradeProfits)
}

func main() {
	lambda.Start(router.Handler)
}
