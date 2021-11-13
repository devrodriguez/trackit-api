package application

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/service"
	"github.com/gin-gonic/gin"
)

type CheckSrv struct {
	repo repository.IChecks
}

func NewCheckService(repo repository.IChecks) service.ICheckService {
	return &CheckSrv{
		repo: repo,
	}
}

func (cs *CheckSrv) Add(c *gin.Context, chk entity.Check) error {
	if err := cs.repo.Create(chk); err != nil {
		panic(err)
	}

	return nil
}

func (cs *CheckSrv) GetByEmployee(c *gin.Context, employeeID string) ([]entity.Check, error) {
	return nil, nil
}
