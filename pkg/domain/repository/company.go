package repository

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/gin-gonic/gin"
)

type ICompanyRepository interface {
	DBGetAll() ([]*entity.Company, error)
	DBCreate(*gin.Context, entity.Company) error
}

type ICompanies interface {
	Get(id string) (entity.Company, error)
	GetAll() ([]entity.Company, error)
	Create(company entity.Company) error
}

type Companies struct {
	ID   int
	Name string
}
