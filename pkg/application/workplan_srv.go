package application

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/service"
)

type workplanService struct {
	repo repository.WorkplanRepository
}

func NewWorkplanService(repo repository.WorkplanRepository) service.WorkplanService {
	return &workplanService{
		repo,
	}
}

func (ws *workplanService) GetById(id string) (*entity.Workplan, error) {
	workplan, err := ws.repo.GetById(id)

	if err != nil {
		return nil, err
	}

	return workplan, nil
}

func (ws *workplanService) GetByEmployee(emp entity.Employee) ([]*entity.Workplan, error) {
	workplans, err := ws.repo.GetByEmployee(emp)

	if err != nil {
		return nil, err
	}

	return workplans, nil
}

func (ws *workplanService) Create(wp entity.Workplan) error {
	err := ws.repo.Create(wp)

	if err != nil {
		return err
	}

	return nil

}
