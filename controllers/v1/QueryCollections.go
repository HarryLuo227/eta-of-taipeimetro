package v1

import (
	"context"
	"eta-of-taipeimetro/mongodb"
	"eta-of-taipeimetro/mongodb/model"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func QueryAllLineTransfer(c *gin.Context) {

	result := []model.LineTransfer{}

	cursor, err := mongodb.LineTransferCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("[Error] Get QueryAll cursor error.")
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		doc := model.LineTransfer{}
		if err := cursor.Decode(&doc); err != nil {
			log.Println("[Error] Decode cursor error.")
			log.Fatal(err)
		}
		result = append(result, doc)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}

func QueryAllS2STravelTime(c *gin.Context) {

	result := []model.S2STravelTime{}

	cursor, err := mongodb.S2STravelTimeCollection.Find(context.TODO(), bson.D{})
	if err != nil {
		log.Println("[Error] Get QueryAll cursor error.")
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		doc := model.S2STravelTime{}
		if err := cursor.Decode(&doc); err != nil {
			log.Println("[Error] Decode cursor error.")
			log.Fatal(err)
		}
		result = append(result, doc)
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}

	c.JSON(http.StatusOK, result)
}

func DurationTesting(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"data": "Hello, Golang gin-gonic!",
	// })
	start := c.Param("startStation")
	end := c.Param("endStation")
	c.String(http.StatusOK, "Start : %s\nEnd : %s\n", start, end)
}
