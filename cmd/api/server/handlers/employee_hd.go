package handlers

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmployeeHandler struct {
	employeeSrv service.IEmployee
}

func NewEmployeeHandler(employeeSrv service.IEmployee) EmployeeHandler {
	return EmployeeHandler{
		employeeSrv,
	}
}

func (e *EmployeeHandler) Create(c *gin.Context) {
	var employee entity.Employee

	if err := c.BindJSON(&employee); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := e.employeeSrv.Add(employee); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, nil)
}
