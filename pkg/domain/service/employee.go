package service

import "github.com/devrodriguez/trackit-go-api/pkg/domain/entity"

type IEmployee interface {
	Add(employee entity.Employee) error
}
