package middlewares

import "github.com/gin-gonic/gin"

func CacheControl() gin.HandlerFunc{
	return func(c *gin.Context) {
		c.Writer.Header().Set("Cache-Control", "max-age=1000")
	}
}
