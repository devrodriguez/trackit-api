package mngdb

import (
	"context"
	"fmt"

	"github.com/devrodriguez/trackit-go-api/pkg/domain/entity"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type CompanyRepository struct {
	cli *mongo.Client
}

func NewCompanyMongoRepo(cli *mongo.Client) *CompanyRepository {
	return &CompanyRepository{
		cli,
	}
}

// Implementation
func (cr *CompanyRepository) DBGetAll() ([]*entity.Company, error) {
	var companies []*entity.Company

	docRef := cr.cli.Database("locateme").Collection("companies")
	opts := options.Find()

	cursor, err := docRef.Find(context.TODO(), bson.M{}, opts)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var company entity.Company

		if err := cursor.Decode(&company); err != nil {
			panic(err)
		}

		companies = append(companies, &company)
	}

	return companies, nil
}

func (cr *CompanyRepository) DBCreate(c *gin.Context, company entity.Company) error {
	docRef := cr.cli.Database("locateme").Collection("companies")

	res, err := docRef.InsertOne(c, company)

	if err != nil {
		return err
	}

	fmt.Println("Insert ID: ", res.InsertedID)

	return nil
}
