package application

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"github.com/gin-gonic/gin"
)

type CheckSrv struct {
	repo repository.CheckRepository
}

func NewCheckService(repo repository.CheckRepository) *CheckSrv {
	return &CheckSrv{
		repo: repo,
	}
}

// Implemantation
func (c *CheckSrv) GetAll() ([]*entity.Check, error) {
	checks, err := c.repo.DBGetAll()

	if err != nil {
		panic(err)
	}

	return checks, nil
}

func (cs *CheckSrv) Create(c *gin.Context, chk entity.Check) error {
	if err := cs.repo.DBCreate(c, chk); err != nil {
		panic(err)
	}

	return nil
}

func (c *CheckSrv) Update(id string, chk entity.Check) error {
	if err := c.repo.DBUpdate(id, chk); err != nil {
		return err
	}

	return nil
}
