package service

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type ICheckService interface {
	Add(c *gin.Context, chk entity.Check) error
	GetByEmployee(c *gin.Context, employeeID string) ([]entity.Check, error)
}
