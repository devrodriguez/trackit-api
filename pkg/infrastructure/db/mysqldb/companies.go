package mysqldb

import (
	"database/sql"
	"fmt"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/devrodriguez/trackit-go-api/pkg/domain/repository"
)

type CompaniesAdapter struct {
	dbConn *sql.DB
}

func (c *CompaniesAdapter) GetAll() ([]entity.Company, error) {
	companies := make([]entity.Company, 0, 10)

	rows, err := c.dbConn.Query("select * from enterprises")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var company entity.Company

		if err := rows.Scan(&company); err != nil {
			return nil, err
		}

		companies = append(companies, company)
	}

	return companies, nil
}

func (c *CompaniesAdapter) Get(id string) (entity.Company, error) {
	var company entity.Company

	row := c.dbConn.QueryRow("select * from companies where id = ?", id)
	if err := row.Scan(&company); err != nil {
		if err == sql.ErrNoRows {
			return company, fmt.Errorf("company id %s not found", id)
		}
		return company, fmt.Errorf("company with id %s get error %v", id, err)
	}

	return company, nil
}

func (c *CompaniesAdapter) Create() error {
	_, err := c.dbConn.Exec("")
	if err != nil {
		return err
	}

	return nil
}

func NewCompaniesAdapter(dbConn *sql.DB) repository.ICompanies {
	return &CompaniesAdapter{dbConn}
}
