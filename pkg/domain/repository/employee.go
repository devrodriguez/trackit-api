package repository

import "github.com/devrodriguez/first-class-api-go/pkg/domain/entity"

type EmployeeRepository interface {
	Create(entity.Employee) error
	ByCredentials(string, string) (*entity.Employee, error)
}
