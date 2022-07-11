package main

import (
	"eta-of-taipeimetro/configuration"
	"eta-of-taipeimetro/controllers"
	"eta-of-taipeimetro/mongodb"
	"eta-of-taipeimetro/routes"
	"fmt"
	"log"

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
	routes.RoutesHandler(router)

	// defer mongodb.CloseConnection()
	router.Run(fmt.Sprintf("%s:%v", configuration.Conf.Address, configuration.Conf.Port))
}
