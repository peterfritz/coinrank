package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"ptr.red/coinrank/controllers"
	"ptr.red/coinrank/database"
)

type Coin struct {
	Name      string `json:"name" bson:"name"`
	Symbol    string `json:"symbol" bson:"symbol"`
	Upvotes   int    `json:"upvotes" bson:"upvotes"`
	Downvotes int    `json:"downvotes" bson:"downvotes"`
}

func main() {
	router := gin.Default()

	database.ConnectDB()

	config := cors.DefaultConfig()
	config.AllowAllOrigins = true

	router.Use(cors.New(config))

	router.GET("/coins", controllers.GetCoins)
	router.POST("/coins/:symbol/upvote", controllers.UpvoteCoin)
	router.POST("/coins/:symbol/downvote", controllers.DownvoteCoin)

	router.Run()
}
