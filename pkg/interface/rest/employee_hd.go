package rest

import (
	"net/http"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/service"
	"github.com/gin-gonic/gin"
)

type EmployeeHandler struct {
	srv service.EmployeeService
}

func NewEmployeeHandler(srv service.EmployeeService) *EmployeeHandler {
	return &EmployeeHandler{
		srv,
	}
}

func (eh *EmployeeHandler) Create(c *gin.Context) {
	var employee entity.Employee

	// Get data from request
	if err := c.BindJSON(&employee); err != nil {

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

	err := eh.srv.Create(employee)

	if err != nil {
		c.JSON(http.StatusBadRequest, err)
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}
