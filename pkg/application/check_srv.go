package application

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/service"
	"github.com/gin-gonic/gin"
)

type checkService struct {
	repo repository.CheckRepository
}

func NewCheckService(repo repository.CheckRepository) service.CheckService {
	return &checkService{
		repo,
	}
}

// Implemantation
func (c *checkService) GetAll() ([]*entity.Check, error) {
	checks, err := c.repo.DBGetAll()

	if err != nil {
		return nil, err
	}

	return checks, nil
}

func (c *checkService) GetBy(email, companyID, date string) ([]*entity.Check, error) {
	checks, err := c.repo.DBGetBy(email, companyID, date)

	if err != nil {
		return nil, err
	}

	return checks, nil
}

func (cs *checkService) Create(c *gin.Context, chk entity.Check) error {
	if err := cs.repo.DBCreate(c, chk); err != nil {
		return err
	}

	return nil
}

func (c *checkService) Update(id string, chk entity.Check) error {
	if err := c.repo.DBUpdate(id, chk); err != nil {
		return err
	}

	return nil
}
