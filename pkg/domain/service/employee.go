package service

import "github.com/devrodriguez/first-class-api-go/pkg/domain/entity"

type EmployeeService interface {
	Create(entity.Employee) error
	ValidateCredentials(string, string) (bool, error)
}
