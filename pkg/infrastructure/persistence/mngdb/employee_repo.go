package mngdb

import (
	"context"
	"fmt"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type employeeRepository struct {
	cli *mongo.Client
}

func NewEmployeeRepository(cli *mongo.Client) repository.EmployeeRepository {
	return &employeeRepository{
		cli,
	}
}

func (e *employeeRepository) Create(emp entity.Employee) error {
	docRef := e.cli.Database("locateme").Collection("employees")

	res, err := docRef.InsertOne(context.TODO(), emp)

	if err != nil {
		return err
	}

	fmt.Println("mngdb.employeeRepository.create: inserted ID ", res.InsertedID)

	return nil
}
