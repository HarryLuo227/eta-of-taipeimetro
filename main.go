package main

import (
	"eta-of-taipeimetro/configuration"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func init() {
	log.Println("[Read from configuration file]")
	configuration.LoadConfig("env")
}

func main() {
	router := gin.Default()

	// Service heartbeat
	router.GET("/heartbeat", heartbeat)
	router.GET("/hello", greeting)

	router.Run(fmt.Sprintf("%s:%v", configuration.Conf.Address, configuration.Conf.Port))
}

func heartbeat(c *gin.Context) {
	c.String(200, "alive")
}

func greeting(c *gin.Context) {
	c.JSON(200, gin.H{
		"data": "Hello, Golang gin-gonic!",
	})
}
