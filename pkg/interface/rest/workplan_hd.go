package rest

import (
	"net/http"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/service"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type WorkplanHandler struct {
	srv service.WorkplanService
}

func NewWorkplanHandler(srv service.WorkplanService) *WorkplanHandler {
	return &WorkplanHandler{
		srv,
	}
}

func (wh *WorkplanHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	_, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Errors: []APIError{
				{
					Title:       http.StatusText(http.StatusBadRequest),
					Description: err.Error(),
					Status:      http.StatusBadRequest,
				},
			},
		})

		return
	}

	workplan, err := wh.srv.GetById(id)

	c.JSON(http.StatusOK, APIResponse{
		Data: workplan,
	})
}

func (wh *WorkplanHandler) GetByEmployee(c *gin.Context) {
	var employee entity.Employee

	id := c.Param("id")
	empID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Errors: []APIError{
				{
					Title:       http.StatusText(http.StatusBadRequest),
					Description: err.Error(),
					Status:      http.StatusBadRequest,
				},
			},
		})

		return
	}

	employee.ID = empID

	workplanes, err := wh.srv.GetByEmployee(employee)

	c.JSON(http.StatusOK, APIResponse{
		Data: workplanes,
	})
}

func (wh *WorkplanHandler) Create(c *gin.Context) {
	var workplan entity.Workplan

	if err := c.BindJSON(&workplan); err != nil {
		c.JSON(http.StatusBadRequest, APIResponse{
			Errors: []APIError{
				{
					Title:       http.StatusText(http.StatusBadRequest),
					Description: err.Error(),
					Status:      http.StatusBadRequest,
				},
			},
		})

		return
	}

	if err := wh.srv.Create(workplan); err != nil {
		c.JSON(http.StatusInternalServerError, APIResponse{
			Errors: []APIError{
				{
					Title:       http.StatusText(http.StatusInternalServerError),
					Description: err.Error(),
					Status:      http.StatusInternalServerError,
				},
			},
		})

		return
	}

	c.JSON(http.StatusOK, APIResponse{})
}
