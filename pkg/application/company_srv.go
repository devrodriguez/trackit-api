package application

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"github.com/gin-gonic/gin"
)

type CompanySrv struct {
	repo repository.ICompanies
}

func NewCompanyService(repo repository.ICompanies) *CompanySrv {
	return &CompanySrv{
		repo: repo,
	}
}

// GetAll ..
func (cs *CompanySrv) GetAll() ([]entity.Company, error) {
	companies, err := cs.repo.GetAll()
	if err != nil {
		panic(err)
	}

	return companies, nil
}

func (cs *CompanySrv) Create(c *gin.Context, company entity.Company) error {
	err := cs.repo.Create(company)

	if err != nil {
		return err
	}

	return nil
}
