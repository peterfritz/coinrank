package controllers

import (
	"context"
	"regexp"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"ptr.red/coinrank/database"
)

type Coin struct {
	Name      string `json:"name" bson:"name"`
	Symbol    string `json:"symbol" bson:"symbol"`
	Upvotes   int    `json:"upvotes" bson:"upvotes"`
	Downvotes int    `json:"downvotes" bson:"downvotes"`
}

func GetCoins(c *gin.Context) {
	var coins []Coin

	collection := database.Client.Database("coinrank").Collection("coins")

	cursor, err := collection.Aggregate(
		context.Background(),
		[]bson.M{
			{
				"$addFields": bson.M{
					"diff": bson.M{
						"$subtract": []interface{}{
							"$upvotes",
							"$downvotes",
						},
					},
				},
			},
			{
				"$sort": bson.M{
					"diff": -1,
				},
			},
			{
				"$unset": "diff",
			},
		},
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	for cursor.Next(context.Background()) {
		var coin Coin
		cursor.Decode(&coin)
		coins = append(coins, coin)
	}

	cursor.Close(context.Background())

	if err := cursor.Err(); err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, coins)
}

func UpvoteCoin(c *gin.Context) {
	symbol := c.Param("symbol")

	matched, _ := regexp.MatchString(`^[A-Z0-9]+$`, symbol)

	if !matched {
		c.JSON(400, gin.H{"error": "Invalid symbol"})

		return
	}

	collection := database.Client.Database("coinrank").Collection("coins")

	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"symbol": symbol},
		bson.M{"$inc": bson.M{"upvotes": 1}},
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	if result.MatchedCount == 0 {
		c.JSON(404, gin.H{"error": "Coin not found"})

		return
	}

	var updatedCoin Coin

	collection.FindOne(
		context.Background(),
		bson.M{"symbol": symbol},
	).Decode(&updatedCoin)

	c.JSON(200, gin.H{"updatedCoin": updatedCoin})
}

func DownvoteCoin(c *gin.Context) {
	symbol := c.Param("symbol")

	matched, _ := regexp.MatchString(`^[A-Z0-9]+$`, symbol)

	if !matched {
		c.JSON(400, gin.H{"error": "Invalid symbol"})

		return
	}

	collection := database.Client.Database("coinrank").Collection("coins")

	result, err := collection.UpdateOne(
		context.Background(),
		bson.M{"symbol": symbol},
		bson.M{"$inc": bson.M{"downvotes": 1}},
	)

	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})

		return
	}

	if result.MatchedCount == 0 {
		c.JSON(404, gin.H{"error": "Coin not found"})

		return
	}

	var updatedCoin Coin

	collection.FindOne(
		context.Background(),
		bson.M{"symbol": symbol},
	).Decode(&updatedCoin)

	c.JSON(200, gin.H{"updatedCoin": updatedCoin})
}
