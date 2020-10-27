package rest

import (
	"net/http"
	"net/url"

	"github.com/devrodriguez/first-class-api-go/pkg/util"
	"github.com/gin-gonic/gin"
)

type geoHandler struct {
}

func NewGeoHandler() geoHandler {
	return geoHandler{}
}

func (gh *geoHandler) AddressPredictions(c *gin.Context) {
	var res APIResponse
	var preds []string

	term := url.QueryEscape(c.Query("term"))
	ga, err := util.GetAddresPredictions(term)

	if err != nil {
		res = APIResponse{
			Message: "error getting predictions",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusBadRequest),
					Status: http.StatusBadRequest,
				},
			},
		}

		c.JSON(http.StatusBadRequest, res)
		return
	}

	for _, p := range ga.Predictions {
		preds = append(preds, p.Description)
	}

	res.Data = preds

	c.JSON(http.StatusOK, preds)

}
