package main_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"

	"ptr.red/coinrank/controllers"
	"ptr.red/coinrank/database"
)

func TestMain(m *testing.M) {
	gin.SetMode(gin.TestMode)

	database.Client.Database("coinrank").Collection("coins").Drop(context.Background())

	database.ResetDB(database.Client)

	defer database.ResetDB(database.Client)

	os.Exit(m.Run())
}

func TestGetCoins(t *testing.T) {
	router := gin.Default()
	router.GET("/coins", controllers.GetCoins)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/coins", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Contains(t, w.Body.String(), "KLV")
}

func TestUpvoteCoin(t *testing.T) {
	router := gin.Default()
	router.POST("/coins/:symbol/upvote", controllers.UpvoteCoin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/coins/KLV/upvote", nil)
	router.ServeHTTP(w, req)

	var result bson.M

	database.Client.Database("coinrank").Collection("coins").FindOne(
		context.Background(),
		bson.M{"symbol": "KLV"},
	).Decode(&result)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, int32(1), result["upvotes"])
}

func TestDownvoteCoin(t *testing.T) {
	router := gin.Default()
	router.POST("/coins/:symbol/downvote", controllers.DownvoteCoin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/coins/ETH/downvote", nil)
	router.ServeHTTP(w, req)

	var result bson.M

	database.Client.Database("coinrank").Collection("coins").FindOne(
		context.Background(),
		bson.M{"symbol": "ETH"},
	).Decode(&result)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, int32(1), result["downvotes"])
}

func TestUpvoteCoinInvalidSymbol(t *testing.T) {
	router := gin.Default()
	router.POST("/coins/:symbol/upvote", controllers.UpvoteCoin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/coins/invalid/upvote", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestDownvoteCoinInvalidSymbol(t *testing.T) {
	router := gin.Default()
	router.POST("/coins/:symbol/downvote", controllers.DownvoteCoin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/coins/invalid/downvote", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 400, w.Code)
}

func TestUpvoteCoinNotFound(t *testing.T) {
	router := gin.Default()
	router.POST("/coins/:symbol/upvote", controllers.UpvoteCoin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/coins/NOTFOUND/upvote", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}

func TestDownvoteCoinNotFound(t *testing.T) {
	router := gin.Default()
	router.POST("/coins/:symbol/downvote", controllers.DownvoteCoin)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/coins/NOTFOUND/downvote", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 404, w.Code)
}
