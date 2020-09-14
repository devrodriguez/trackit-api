package rest

import (
	"net/http"

	"github.com/devrodriguez/first-class-api-go/pkg/application"
	"github.com/gin-gonic/gin"
)

type CompanyHandler struct {
	srv *application.CompanyService
}

func NewCompanyHandler(srv *application.CompanyService) *CompanyHandler {
	return &CompanyHandler{
		srv,
	}
}

// Implementation
func (ch *CompanyHandler) Get(c *gin.Context) {

	company, err := ch.srv.Get()

	if err != nil {
		panic(err)
	}

	c.JSON(http.StatusOK, company)
}
