package handlers

import (
	"fmt"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
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
func (ch *CheckHandler) GetEmployeeChecked(c *gin.Context) {
	employee := c.Param("emp_id")
	company := c.Query("company")
	fromDate := c.Query("date")

	companyID, err := strconv.ParseUint(company, 10, 32)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	employeeID, err := strconv.Atoi(employee)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	checks, err := ch.srv.CheckedByEmployee(entity.Employee{
		ID: uint(employeeID),
	}, entity.Company{
		ID: uint(companyID),
	}, fromDate)

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

func (ch *CheckHandler) GetEmployeeChecks(c *gin.Context) {
	employee := c.Param("emp_id")
	employeeID, err := strconv.Atoi(employee)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	checks, err := ch.srv.ChecksByEmployee(entity.Employee{
		ID: uint(employeeID),
	})
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
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
