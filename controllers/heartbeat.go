package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Heartbeat(c *gin.Context) {
	c.String(http.StatusOK, "Alive")
}
