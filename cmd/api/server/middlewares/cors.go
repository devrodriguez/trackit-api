package middlewares

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func EnableCORS() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("c.Request.Method: ", c.Request.Method)
			c.AbortWithStatus(http.StatusOK)
			return
		}

		c.Next()
	}
}
