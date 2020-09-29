package rest

import (
	"fmt"
	"net/http"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/service"
	"github.com/gin-gonic/gin"
)

type CheckHandler struct {
	srv service.CheckService
}

func NewCheckHandler(srv service.CheckService) *CheckHandler {
	return &CheckHandler{
		srv: srv,
	}
}

// Implementation
func (ch *CheckHandler) GetChecks(c *gin.Context) {

	checks, err := ch.srv.GetAll()

	if err != nil {
		c.JSON(http.StatusNoContent, APIResponse{
			Message: "not data found",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusNoContent),
					Status: http.StatusNoContent,
				},
			},
		})

		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Data: checks,
	})
}

func (ch *CheckHandler) GetChecksBy(c *gin.Context) {
	email := c.Query("email")
	companyID := c.Query("company_id")
	date := c.Query("date")

	checks, err := ch.srv.GetBy(email, companyID, date)

	if err != nil {
		c.JSON(http.StatusNoContent, APIResponse{
			Message: "not data found",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusNoContent),
					Status: http.StatusNoContent,
				},
			},
		})

		return
	}

	if checks == nil {
		checks = []*entity.Check{}
	}

	c.JSON(http.StatusOK, APIResponse{
		Data: checks,
	})
}

func (ch *CheckHandler) Create(c *gin.Context) {
	var check entity.Check

	// Get data from request
	if err := c.BindJSON(&check); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "error binding data",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusBadRequest),
					Status: http.StatusBadRequest,
				},
			},
		})
		return
	}

	if err := ch.srv.Create(c, check); err != nil {
		c.JSON(http.StatusNotModified, APIResponse{
			Message: "error saving data",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusNotModified),
					Status: http.StatusNotModified,
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: `success`,
	})
}

func (ch *CheckHandler) Update(c *gin.Context) {
	var check entity.Check
	id := c.Params.ByName("id")

	if err := c.BindJSON(&check); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Message: "error binding data",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusBadRequest),
					Status: http.StatusBadRequest,
				},
			},
		})
		return
	}

	if err := ch.srv.Update(id, check); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotModified, APIResponse{
			Message: "error updating data",
			Errors: []APIError{
				{
					Title:  http.StatusText(http.StatusNotModified),
					Status: http.StatusNotModified,
				},
			},
		})
		return
	}

	c.JSON(http.StatusOK, APIResponse{
		Message: "check updated",
	})
}
