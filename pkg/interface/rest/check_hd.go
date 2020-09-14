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
		c.JSON(http.StatusNotFound, gin.H{"message": "not found"})
	}

	c.JSON(http.StatusOK, checks)
}

func (ch *CheckHandler) Create(c *gin.Context) {
	var check entity.Check
	// Get data from request
	if err := c.BindJSON(&check); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error binding data"})
		return
	}

	if err := ch.srv.Create(c, check); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error saving data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ch *CheckHandler) Update(c *gin.Context) {
	var check entity.Check
	id := c.Params.ByName("id")

	if err := c.BindJSON(&check); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error binding data"})
		return
	}

	if err := ch.srv.Update(id, check); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "error updating data"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "updated"})
}
