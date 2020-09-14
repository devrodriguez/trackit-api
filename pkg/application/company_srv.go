package application

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"github.com/gin-gonic/gin"
)

type CompanyService struct {
	repo repository.CompanyRepository
}

func NewCompanyService(repo repository.CompanyRepository) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

// Implementation
func (cs *CompanyService) GetAll() ([]*entity.Company, error) {
	company, err := cs.repo.DBGetAll()

	if err != nil {
		panic(err)
	}

	return company, nil
}

func (cs *CompanyService) Create(c *gin.Context, company entity.Company) error {
	err := cs.repo.DBCreate(c, company)

	if err != nil {
		return err
	}

	return nil
}
