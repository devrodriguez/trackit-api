package repository

import "github.com/devrodriguez/trackit-go-api/pkg/domain/entity"

type IEmployee interface {
	Create(employee entity.Employee) error
}
