package db

import (
	"context"
	"fmt"
	"time"

	"github.com/MikeB1124/stocks-profit-calculator-lambda/configuration"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoClient *mongo.Client

func init() {
	config := configuration.GetConfig()
	// Connect to MongoDB
	opts := options.Client().ApplyURI(fmt.Sprintf("mongodb+srv://%s:%s@cluster0.du0vf.mongodb.net", config.MongoDB.Username, config.MongoDB.Password))
	client, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}
	mongoClient = client
}

func UpdateAllExpiredOrders() (*mongo.UpdateResult, error) {
	// Update all expired orders
	collection := mongoClient.Database("Stocks").Collection("orders")
	filter := bson.M{"order.status": "expired", "tradeCompleted": false}
	update := bson.M{
		"$set": bson.M{
			"tradeCompleted":  true,
			"recordUpdatedAt": time.Now().UTC(),
			"tradeProfit":     0,
		},
	}
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func UpdateAllCancelledOrders() (*mongo.UpdateResult, error) {
	// Update all cancelled orders
	collection := mongoClient.Database("Stocks").Collection("orders")
	filter := bson.M{"order.status": "canceled", "tradeCompleted": false}
	update := bson.M{
		"$set": bson.M{
			"tradeCompleted":  true,
			"recordUpdatedAt": time.Now().UTC(),
			"tradeProfit":     0,
		},
	}
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	if err != nil {
		return nil, err
	}
	return result, nil
}
