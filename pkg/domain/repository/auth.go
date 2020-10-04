package repository

import "github.com/devrodriguez/first-class-api-go/pkg/domain/entity"

type AuthRepository interface {
	Register(entity.User) error
	ByCredentials(string, string) (*entity.Employee, error)
	ByEmail(string) (*entity.Employee, error)
}
