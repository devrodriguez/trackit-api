package repository

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type CheckRepository interface {
	DBGetAll() ([]*entity.Check, error)
	DBGetBy(string, string, string) ([]*entity.Check, error)
	DBCreate(*gin.Context, entity.Check) error
	DBUpdate(string, entity.Check) error
}
