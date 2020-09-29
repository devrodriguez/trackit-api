package mngdb

import (
	"context"
	"fmt"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
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

func (e *employeeRepository) ByCredentials(usr string, pass string) (*entity.Employee, error) {
	var employee entity.Employee

	docRef := e.cli.Database("locateme").Collection("employees")
	filter := bson.M{
		"internal_user.email":    bson.M{"$eq": usr},
		"internal_user.password": bson.M{"$eq": pass},
	}

	res := docRef.FindOne(context.TODO(), filter)
	err := res.Decode(&employee)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}
