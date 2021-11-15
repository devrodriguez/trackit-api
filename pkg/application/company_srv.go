package application

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"github.com/gin-gonic/gin"
)

type CompanyService struct {
	repo repository.ICompanies
}

func NewCompanyService(repo repository.ICompanies) *CompanyService {
	return &CompanyService{
		repo: repo,
	}
}

// GetAll ..
func (cs *CompanyService) GetAll() ([]entity.Company, error) {
	companies, err := cs.repo.GetAll()
	if err != nil {
		panic(err)
	}

	return companies, nil
}

func (cs *CompanyService) Create(c *gin.Context, company entity.Company) error {
	err := cs.repo.Create(company)

	if err != nil {
		return err
	}

	return nil
}
