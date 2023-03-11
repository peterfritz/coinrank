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

func ConnectDB() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}

	uri := os.Getenv("MONGO_URL")

	if uri == "" {
		log.Fatal("You must set your 'MONGO_URL' environmental variable. See https://www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	if err := client.Database("coinrank").CreateCollection(context.Background(), "coins"); err == nil {
		ResetDB(client)
	}

	return client
}

var Client *mongo.Client = ConnectDB()

func ResetDB(client *mongo.Client) {
	client.Database("coinrank").Collection("coins").DeleteMany(context.Background(), bson.M{})

	coins := []interface{}{
		bson.M{"name": "Bitcoin", "symbol": "BTC", "upvotes": 0, "downvotes": 0},
		bson.M{"name": "Ethereum", "symbol": "ETH", "upvotes": 0, "downvotes": 0},
		bson.M{"name": "Klever", "symbol": "KLV", "upvotes": 0, "downvotes": 0},
	}

	client.Database("coinrank").Collection("coins").InsertMany(context.Background(), coins)
}
