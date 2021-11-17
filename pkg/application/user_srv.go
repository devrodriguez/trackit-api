package application

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/service"
)

type UserSrv struct {
	repo repository.IUserRepository
}

func NewUserSrv(repo repository.IUserRepository) service.IUserService {
	return &UserSrv{
		repo,
	}
}

func (u *UserSrv) Create(user entity.User) error {
	if err := u.repo.Insert(user); err != nil {
		return err
	}

	return nil
}
