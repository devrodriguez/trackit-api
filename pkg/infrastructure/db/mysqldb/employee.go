package mysqldb

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"gorm.io/gorm"
)

type EmployeeAdapter struct {
	dbConn *gorm.DB
}

func NewEmployeeAdapter(dbConn *gorm.DB) repository.IEmployeeRepository {
	return &EmployeeAdapter{
		dbConn,
	}
}

func (e *EmployeeAdapter) Insert(employee entity.Employee) error {
	tx := e.dbConn.Create(&employee)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}
