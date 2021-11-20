package repository

import "github.com/devrodriguez/trackit-go-api/pkg/domain/entity"

type IUserRepository interface {
	Insert(user entity.User) error
	Check(email string, password string) error
}
