package service

import "github.com/devrodriguez/trackit-go-api/pkg/domain/entity"

type IUserService interface {
	Create(user entity.User) error
}
