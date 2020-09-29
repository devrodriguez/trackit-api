package repository

import "github.com/devrodriguez/first-class-api-go/pkg/domain/entity"

type WorkplanRepository interface {
	GetById(string) (*entity.Workplan, error)
	GetByEmployee(entity.Employee) ([]*entity.Workplan, error)
	Create(entity.Workplan) error
}
