package application

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/service"
)

type employeeService struct {
	repo repository.EmployeeRepository
}

func NewEmployeeService(repo repository.EmployeeRepository) service.EmployeeService {
	return &employeeService{
		repo,
	}
}

func (e *employeeService) Create(emp entity.Employee) error {
	err := e.repo.Create(emp)

	if err != nil {
		return err
	}

	return nil
}
