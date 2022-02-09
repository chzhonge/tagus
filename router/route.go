package router

import (
	"github.com/gin-gonic/gin"
	"tagus/api"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", api.Pong)
	r.GET("/check", api.HealthCheck)

	return r
}
