package main

import (
	"eta-of-taipeimetro/configuration"
	"eta-of-taipeimetro/controllers"
	"eta-of-taipeimetro/mongodb"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("[Info] Read configuration file.")
	configuration.LoadConfig("env")

	mongodb.Initialize()
}

func main() {
	router := gin.Default()

	// Service heartbeat
	router.GET("/heartbeat", controllers.Heartbeat)

	// api
	// api := router.Group("/api/v1")

	router.GET("/api/v1/LineTransfer", mongodb.QueryAllLineTransfer)
	// router.GET("/api/v1/S2STravelTime", controllers.v1)

	router.GET("/hello/:startStation/:endStation", greeting)

	// defer mongodb.CloseConnection()
	router.Run(fmt.Sprintf("%s:%v", configuration.Conf.Address, configuration.Conf.Port))
}

func greeting(c *gin.Context) {
	// c.JSON(200, gin.H{
	// 	"data": "Hello, Golang gin-gonic!",
	// })
	start := c.Param("startStation")
	end := c.Param("endStation")
	c.String(http.StatusOK, "Start : %s\nEnd : %s\n", start, end)
}
