package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"tagus/cache"
)

func TokenAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header["Authorization"]

		fmt.Println(token)
		if len(token) == 0 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": false, "message": "authorization required"})
			return
		}

		_, found := cache.Manga.Get(token[0])
		if !found {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": false, "message": "Unauthorized"})
			return
		}

		c.Next()
	}
}
