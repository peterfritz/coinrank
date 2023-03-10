package database

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)


func ConnectDB() *mongo.Client  {
  if err := godotenv.Load(); err != nil {
    log.Println("No .env file found")
  }

  uri := os.Getenv("MONGODB_URI")

  if uri == "" {
    log.Fatal("You must set your 'MONGODB_URI' environmental variable. See\n\t https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
  }

  client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

  if err != nil {
    panic(err)
  }

  if err := client.Database("coinrank").CreateCollection(context.Background(), "coins"); err == nil {
    coins := []interface{}{
      bson.M{"name": "Bitcoin", "symbol": "BTC", "upvotes": 0, "downvotes": 0},
      bson.M{"name": "Ethereum", "symbol": "ETH", "upvotes": 0, "downvotes": 0},
      bson.M{"name": "Ripple", "symbol": "XRP", "upvotes": 0, "downvotes": 0},
    }

    client.Database("coinrank").Collection("coins").InsertMany(context.Background(), coins)
  }

  return client
}

var Client *mongo.Client = ConnectDB()
