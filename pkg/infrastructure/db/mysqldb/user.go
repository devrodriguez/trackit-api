package mysqldb

import (
	"errors"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
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
	var count int64
	tx := u.dbConn.Table("users").
		Where("email = ? and password = ?", email, password).
		Count(&count)

	if tx.Error != nil {
		return tx.Error
	}

	if count == 0 {
		return errors.New("not found")
	}

	return nil
}
