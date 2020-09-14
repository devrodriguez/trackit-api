package service

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type CheckService interface {
	GetAll() ([]*entity.Check, error)
	Create(c *gin.Context, chk entity.Check) error
	Update(id string, chk entity.Check) error
}
