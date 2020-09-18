package middlewares

import (
	"net/http"

	"github.com/devrodriguez/first-class-api-go/pkg/interface/rest"
	"github.com/gin-gonic/gin"
)

func ValidateAuth() gin.HandlerFunc {
	var resModel rest.APIResponse

	return func(gCtx *gin.Context) {
		err := rest.VerifyToken(gCtx.Request)

		if err != nil {
			resModel.Message = "you need to be authorized"
			resModel.Errors = []rest.APIError{
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
