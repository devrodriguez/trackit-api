package repository

import "github.com/devrodriguez/first-class-api-go/pkg/domain/entity"

type CompanyRepository interface {
	DBGet() (entity.Company, error)
}
