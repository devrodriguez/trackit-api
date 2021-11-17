package handlers

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type UserHandler struct {
	userSrv service.IUserService
}

func NewUserHandler(userSrv service.IUserService) UserHandler {
	return UserHandler{
		userSrv,
	}
}

func (e *UserHandler) Create(c *gin.Context) {
	var user entity.User

	if err := c.BindJSON(&user); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	if err := e.userSrv.Create(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, nil)
}
