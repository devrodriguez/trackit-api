package repository

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type CheckRepository interface {
	DBGetAll() ([]*entity.Check, error)
	DBCreate(*gin.Context, entity.Check) error
	DBUpdate(string, entity.Check) error
}

type IChecks interface {
	Create(check entity.Check) error
}

type Check struct {
	ID   string
	Name string
}
