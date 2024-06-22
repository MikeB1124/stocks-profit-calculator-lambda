package db

import (
	"github.com/alpacahq/alpaca-trade-api-go/v3/alpaca"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PatternData struct {
	PatternType   string  `json:"patterntype" bson:"patterntype"`
	PatternName   string  `json:"patternname" bson:"patternname"`
	ProfitOne     float64 `json:"profit1" bson:"profit1"`
	DisplaySymbol string  `json:"displaysymbol" bson:"displaysymbol"`
	Symbol        string  `json:"symbol" bson:"symbol"`
	StopLoss      float64 `json:"stoploss" bson:"stoploss"`
	PatternUrl    string  `json:"url" bson:"url"`
	TimeFrame     string  `json:"timeframe" bson:"timeframe"`
	Status        string  `json:"status" bson:"status"`
	Entry         string  `json:"entry" bson:"entry"`
	PatternClass  string  `json:"patternclass" bsom:"patternclass"`
}

type AlpacaEntryOrder struct {
	ObjectID       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Order          *alpaca.Order      `json:"order" bson:"order"`
	TradeCompleted bool               `json:"tradeCompleted" bson:"tradeCompleted"`
	TradeProfit    float64            `json:"tradeProfit" bson:"tradeProfit"`
	PatternData    PatternData        `json:"patternData" bson:"patternData"`
}
