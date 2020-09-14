package mngdb

import (
	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type CompanyRepository struct {
	mongoCli *mongo.Client
}

func NewCompanyMongoRepo(cli *mongo.Client) *CompanyRepository {
	return &CompanyRepository{
		mongoCli: cli,
	}
}

// Implementation
func (cr *CompanyRepository) DBGet() (entity.Company, error) {
	return entity.Company{}, nil
}
