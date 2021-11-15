package handlers

import (
	"fmt"
	"net/http"

	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/service"
	"github.com/gin-gonic/gin"
)

type CheckHandler struct {
	srv service.ICheckService
}

func NewCheckHandler(srv service.ICheckService) *CheckHandler {
	return &CheckHandler{
		srv: srv,
	}
}

// GetChecks ..
func (ch *CheckHandler) GetChecks(c *gin.Context) {
	employeeID := c.Query("emp_id")
	checks, err := ch.srv.GetByEmployee(c, employeeID)

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

	if err := ch.srv.Add(c, check); err != nil {
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
