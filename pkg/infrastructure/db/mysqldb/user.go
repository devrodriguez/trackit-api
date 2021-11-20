package mysqldb

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserAdapter struct {
	dbConn *gorm.DB
}

func NewUserAdapter(dbConn *gorm.DB) repository.IUserRepository {
	return &UserAdapter{
		dbConn,
	}
}

func (u *UserAdapter) Insert(user entity.User) error {
	tx := u.dbConn.Create(&user)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func (u *UserAdapter) Check(email string, password string) error {
	var user entity.User

	tx := u.dbConn.Table("users").
		Where("email = ?", email).
		First(&user)

	if tx.Error != nil {
		return tx.Error
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return err
	}

	return nil
}
