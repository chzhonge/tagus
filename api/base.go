package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Pong(c *gin.Context) {
	c.JSON(200, gin.H{"error": false})
}

func HealthCheck(c *gin.Context) {
	c.String(http.StatusOK, "# tag11us")
}
