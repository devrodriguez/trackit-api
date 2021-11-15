package mysqldb

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"gorm.io/gorm"
)

type EmployeeAdapter struct {
	dbConn *gorm.DB
}

func NewEmployeeAdapter(dbConn *gorm.DB) repository.IEmployee {
	return &EmployeeAdapter{
		dbConn,
	}
}

func (e *EmployeeAdapter) Create(employee entity.Employee) error {
	tx := e.dbConn.Create(&employee)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
