package service

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type ICheckService interface {
	Add(c *gin.Context, chk entity.Check) error
	ByEmployee(employee entity.Employee) (checks []entity.CheckQuery, err error)
}
