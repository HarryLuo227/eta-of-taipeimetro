package v1

import (
	"context"
	"eta-of-taipeimetro/mongodb"
	"eta-of-taipeimetro/mongodb/model"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
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
	// Parameters
	start := c.Param("startStation")
	end := c.Param("endStation")

	startLineID, endLineID := strings.Trim(start, "1234567890"), strings.Trim(end, "1234567890")
	startStationID, err := strconv.Atoi(strings.Trim(start, startLineID))
	if err != nil {
		log.Println("[Error] Parse end string to int error.")
		panic(err)
	}
	endStationID, err := strconv.Atoi(strings.Trim(end, endLineID))
	if err != nil {
		log.Println("[Error] Parse end string to int error.")
		panic(err)
	}

	if startStationID < endStationID {
		startStationID, endStationID = swap(startStationID, endStationID)
	}
	totalStation := sumOfStations(startLineID)

	pipeline := mongo.Pipeline{}
	if startLineID == endLineID {
		routeID := fmt.Sprintf("%s-1", startLineID)

		// Succeed in mongosh `db.S2STravelTime.aggregate([{$match: {RouteID: "BL-1"}}, {$unwind: "$TravelTimes"}, {$match: {$and: [{"TravelTimes.Sequence": {$gt: 18, $lte: 22}}]}}])`
		matchStage1 := bson.D{{"RouteID", routeID}}
		unwindKey := "$TravelTimes"
		matchStage2 := bson.D{{"$and", bson.A{
			bson.D{{"TravelTimes.Sequence", bson.D{{"$gt", totalStation - startStationID}, {"$lte", totalStation - endStationID}}}},
		}}}
		pipeline = mongo.Pipeline{{{"$match", matchStage1}}, {{"$unwind", unwindKey}}, {{"$match", matchStage2}}}
	}

	// specifyQuery(pipeline)
	c.JSON(http.StatusOK, calculateTime(pipeline))

}

func specifyQuery(pipeline mongo.Pipeline) []model.UnwindS2STravelTime {
	result := []model.UnwindS2STravelTime{}

	// Query documents
	filterCursor, err := mongodb.S2STravelTimeCollection.Aggregate(context.TODO(), pipeline)
	if err != nil {
		log.Println("[Error] Get cursor error.")
		panic(err)
	}
	for filterCursor.Next(context.TODO()) {
		doc := model.UnwindS2STravelTime{}
		if err := filterCursor.Decode(&doc); err != nil {
			log.Println("[Error] Decode cursor error.")
			log.Fatal(err)
		}
		result = append(result, doc)
	}
	log.Println("\nSpecifyQuery\n", result)

	return result
}

func sumOfStations(lineID string) int {
	result := model.StationOfLine{}

	// matchStage1 := bson.D{{"LineID", lineID}}
	// unwindKey := "$Stations"
	// pipeline := mongo.Pipeline{{{"$match", matchStage1}}, {{"$unwind", unwindKey}}}

	// cursor, err := mongodb.StationOfLineCollection.Aggregate(cotext.TODO(), pipeline)
	cursor, err := mongodb.StationOfLineCollection.Find(context.TODO(), bson.D{{"LineID", lineID}})
	if err != nil {
		log.Println("[Error] Get QueryAll cursor error.")
		panic(err)
	}

	for cursor.Next(context.TODO()) {
		doc := model.StationOfLine{}
		if err := cursor.Decode(&doc); err != nil {
			log.Println("[Error] Decode cursor error.")
			log.Fatal(err)
		}
		result = doc
	}
	if err := cursor.Err(); err != nil {
		log.Fatal(err)
	}
	log.Println("\nTotal station\n", result)

	return len(result.Stations)
}

func swap(a, b int) (int, int) {
	return b, a
}

func calculateTime(pipeline mongo.Pipeline) int32 {
	docs := specifyQuery(pipeline)

	var result int32 = 0
	for _, value := range docs {
		total := value.TravelTimes.RunTime + value.TravelTimes.StopTime
		result += total
	}

	return result
}
