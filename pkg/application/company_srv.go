package application

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
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
func (cs *CompanyService) Get() (entity.Company, error) {
	company, err := cs.repo.DBGet()

	if err != nil {
		panic(err)
	}

	return company, nil
}
