package handlers

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/service"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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

	hashPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), 8)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
	}

	user.Password = string(hashPass)

	if err := e.userSrv.Create(user); err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	c.JSON(http.StatusOK, nil)
}
