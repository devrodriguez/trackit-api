package middlewares

import (
	"github.com/devrodriguez/trackit-go-api/cmd/api/server/handlers"
	"net/http"

	"github.com/gin-gonic/gin"
)

func ValidateAuth() gin.HandlerFunc {
	var resModel handlers.APIResponse

	return func(gCtx *gin.Context) {
		err := handlers.VerifyToken(gCtx.Request)

		if err != nil {
			resModel.Message = "you need to be authorized"
			resModel.Errors = []handlers.APIError{
				{
					Status:      http.StatusUnauthorized,
					Title:       http.StatusText(http.StatusUnauthorized),
					Description: err.Error(),
				},
			}

			gCtx.JSON(http.StatusUnauthorized, resModel)
			gCtx.Abort()
			return
		}

		gCtx.Next()
	}
}
