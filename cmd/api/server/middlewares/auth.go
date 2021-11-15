package middlewares

import (
	"github.com/devrodriguez/trackit-go-api/cmd/api/server/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAuth() gin.HandlerFunc {
	var resModel handlers.APIResponse

	return func(c *gin.Context) {
		return
		err := handlers.VerifyToken(c.Request)

		if err != nil {
			resModel.Message = "you need to be authorized"
			resModel.Errors = []handlers.APIError{
				{
					Status:      http.StatusUnauthorized,
					Title:       http.StatusText(http.StatusUnauthorized),
					Description: err.Error(),
				},
			}

			c.JSON(http.StatusUnauthorized, resModel)
			c.Abort()
			return
		}

		c.Next()
	}
}
