package router

import (
	"github.com/gin-gonic/gin"
	"tagus/api"
)

func SetRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", api.Pong)
	r.GET("/check", api.HealthCheck)
	r.POST("/signon", api.SignOn)
	r.POST("/signin", api.SignIn)

	return r
}
