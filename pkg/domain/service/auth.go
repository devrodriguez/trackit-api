package service

import "github.com/devrodriguez/first-class-api-go/pkg/domain/entity"

type AuthService interface {
	Register(entity.User) error
	ValidateCredentials(string, string) (bool, error)
}
