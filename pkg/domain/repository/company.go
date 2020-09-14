package repository

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type CompanyRepository interface {
	DBGetAll() ([]*entity.Company, error)
	DBCreate(*gin.Context, entity.Company) error
}
