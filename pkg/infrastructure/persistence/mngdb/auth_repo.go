package mngdb

import (
	"context"
	"fmt"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type authRepository struct {
	Cli *mongo.Client
}

func NewAuthRepository(cli *mongo.Client) repository.AuthRepository {
	return &authRepository{
		cli,
	}
}

func (ar *authRepository) Register(user entity.User) error {
	docRef := ar.Cli.Database("locateme").Collection("users")
	res, err := docRef.InsertOne(context.TODO(), user)

	if err != nil {
		return err
	}

	fmt.Println("mngdb.Register: ", res.InsertedID)

	return nil
}

func (ar *authRepository) ByEmail(email string) (*entity.Employee, error) {
	var employee entity.Employee

	docRef := ar.Cli.Database("locateme").Collection("users")
	filter := bson.M{
		"email": bson.M{"$eq": email},
	}

	res := docRef.FindOne(context.TODO(), filter)
	err := res.Decode(&employee)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func (ar *authRepository) ByCredentials(usr string, pass string) (*entity.Employee, error) {
	var employee entity.Employee

	docRef := ar.Cli.Database("locateme").Collection("users")
	filter := bson.M{
		"email":    bson.M{"$eq": usr},
		"password": bson.M{"$eq": pass},
	}

	res := docRef.FindOne(context.TODO(), filter)
	err := res.Decode(&employee)

	if err != nil {
		return nil, err
	}

	return &employee, nil
}
