package mysqldb

import (
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
	"gorm.io/gorm"
)

type CompaniesAdapter struct {
	dbConn *gorm.DB
}

func (c *CompaniesAdapter) GetAll() ([]entity.Company, error) {
	companies := make([]entity.Company, 0, 10)

	txn := c.dbConn.Find(&companies)
	if txn.Error != nil {
		return nil, txn.Error
	}

	return companies, nil
}

func (c *CompaniesAdapter) Get(id string) (entity.Company, error) {
	var company entity.Company

	txn := c.dbConn.First(&company, id)
	if txn.Error != nil {
		return company, txn.Error
	}

	return company, nil
}

func (c *CompaniesAdapter) Create(company entity.Company) error {
	txn := c.dbConn.Create(&company)
	if txn.Error != nil {
		return txn.Error
	}

	return nil
}

func NewCompaniesAdapter(dbConn *gorm.DB) repository.ICompanies {
	return &CompaniesAdapter{dbConn}
}
