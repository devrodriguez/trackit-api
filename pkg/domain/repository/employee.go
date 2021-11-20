package repository

import "github.com/devrodriguez/trackit-go-api/pkg/domain/entity"

type IEmployeeRepository interface {
	Insert(employee entity.Employee) error
}
