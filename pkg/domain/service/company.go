package service

import "github.com/devrodriguez/first-class-api-go/pkg/domain/entity"

type CompanyService interface {
	Get() entity.Company
}
