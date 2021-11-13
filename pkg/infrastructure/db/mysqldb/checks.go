package mysqldb

import (
	"database/sql"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
)

type ChecksAdapter struct {
	dbConn *sql.DB
}

func (c *ChecksAdapter) Create(check entity.Check) error {
	_, err := c.dbConn.Exec("insert into checks (name, employee_id, company_id) values(?, ?, ?)", "")
	if err != nil {
		return err
	}

	return nil
}

func NewCheckAdapter(dbConn *sql.DB) repository.IChecks {
	return &ChecksAdapter{dbConn}
}
