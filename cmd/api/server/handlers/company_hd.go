package handlers

import (
	"net/http"

	"github.com/devrodriguez/trackit-go-api/pkg/application"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	srv *application.CompanySrv
}

func NewCompanyHandler(srv *application.CompanySrv) *CompanyHandler {
	return &CompanyHandler{
		srv,
	}
}

// Implementation
func (ch *CompanyHandler) GetAll(c *gin.Context) {

	companies, err := ch.srv.GetAll()

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
		Data: companies,
	})
}

// Implementation
func (ch *CompanyHandler) Create(c *gin.Context) {
	var company entity.Company

	// Get data from request
	if err := c.BindJSON(&company); err != nil {

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

	if err := ch.srv.Create(c, company); err != nil {
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
