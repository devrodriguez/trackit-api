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

func (cs *CheckSrv) ByEmployee(employee entity.Employee) (checks []entity.CheckQuery, err error) {
	checks, err = cs.repo.QueryByEmployee(employee)
	if err != nil {
		return nil, err
	}

	return checks, nil
}
