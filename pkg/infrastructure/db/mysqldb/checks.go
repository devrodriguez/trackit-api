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

func (c *ChecksAdapter) QueryByEmployee(employee entity.Employee) ([]entity.CheckQuery, error) {
	var checks []entity.CheckQuery
	selectQuery := `checks.id as check_id, checks.date, checks.address, checks.hour,
					check_types.id as check_type_id, check_types.name as check_type_name, 
					employees.id as employee_id, employees.name as employee_name, employees.email,
					companies.id as company_id, companies.name as company_name`

	rows, err := c.dbConn.
		Table("checks").
		Select(selectQuery).
		Joins("inner join check_types on checks.check_type_id = check_types.id").
		Joins("inner join employees on checks.employee_id = employees.id").
		Joins("inner join companies on checks.company_id = companies.id").
		Where("checks.employee_id = ?", employee.ID).
		Rows()

	if err != nil {
		return nil, err
	}

	for rows.Next() {
		c.dbConn.ScanRows(rows, &checks)
	}

	return checks, nil
}

func NewCheckAdapter(dbConn *gorm.DB) repository.IChecks {
	return &ChecksAdapter{dbConn}
}
