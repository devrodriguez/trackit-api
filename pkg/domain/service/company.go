package service

import "github.com/devrodriguez/trackit-go-api/pkg/domain/entity"

type CompanyService interface {
	GetAll() []*entity.Company
	Create(entity.Company) error
}
