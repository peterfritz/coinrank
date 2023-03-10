package controllers

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"ptr.red/coinrank/database"
)

type Coin struct {
  Name       string  `json:"name" bson:"name"`
  Symbol     string  `json:"symbol" bson:"symbol"`
  Upvotes    int     `json:"upvotes" bson:"upvotes"`
  Downvotes  int     `json:"downvotes" bson:"downvotes"`
}

func GetCoins(c *gin.Context) {
  var coins []Coin

  cursor, err := database.Client.Database("coinrank").Collection("coins").Find(context.Background(), bson.M{})

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

  result, err := database.Client.Database("coinrank").Collection("coins").UpdateOne(
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

  database.Client.Database("coinrank").Collection("coins").FindOne(
    context.Background(), 
    bson.M{"symbol": symbol},
  ).Decode(&updatedCoin)

  c.JSON(200, gin.H{"updatedCoin": updatedCoin})
}

func DownvoteCoin(c *gin.Context) {
  symbol := c.Param("symbol")

  result, err := database.Client.Database("coinrank").Collection("coins").UpdateOne(
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

  database.Client.Database("coinrank").Collection("coins").FindOne(
    context.Background(), 
    bson.M{"symbol": symbol},
  ).Decode(&updatedCoin)

  c.JSON(200, gin.H{"updatedCoin": updatedCoin})
}
