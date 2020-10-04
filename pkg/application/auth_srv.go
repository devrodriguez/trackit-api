package application

import (
	"errors"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/service"
	"go.mongodb.org/mongo-driver/mongo"
)

var errUserExist = errors.New("user exist")

type authService struct {
	repo repository.AuthRepository
}

func NewAuthService(repo repository.AuthRepository) service.AuthService {
	return &authService{
		repo,
	}
}

func (as *authService) Register(user entity.User) error {

	xuser, err := as.repo.ByEmail(user.Email)

	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}

	if xuser != nil {
		return errUserExist
	}

	if err := as.repo.Register(user); err != nil {
		return err
	}

	return nil
}

func (as *authService) ValidateCredentials(email string, pass string) (bool, error) {
	emp, err := as.repo.ByCredentials(email, pass)

	if err != nil || emp == nil {
		return false, err
	}

	return true, nil
}
