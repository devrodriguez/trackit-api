package mysqldb

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"gorm.io/gorm"
)

type ChecksAdapter struct {
	dbConn *gorm.DB
}

func (c *ChecksAdapter) Create(check entity.Check) error {
	tx := c.dbConn.Create(&check)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func NewCheckAdapter(dbConn *gorm.DB) repository.IChecks {
	return &ChecksAdapter{dbConn}
}
