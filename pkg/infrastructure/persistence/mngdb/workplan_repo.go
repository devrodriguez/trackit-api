package mngdb

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"go.mongodb.org/mongo-driver/mongo"
)

type workplanRepository struct {
	cli *mongo.Client
}

func NewWorkplanRepository(cli *mongo.Client) repository.WorkplanRepository {
	return &workplanRepository{
		cli,
	}
}

func (wr *workplanRepository) GetById(id string) (*entity.Workplan, error) {
	var wp *entity.Workplan

	wpID, err := primitive.ObjectIDFromHex(id)

	docRef := wr.cli.Database("locateme").Collection("workplans")
	filter := bson.M{"_id": bson.M{"$eq": wpID}}
	result := docRef.FindOne(context.TODO(), filter)
	err = result.Decode(&wp)

	if err != nil {
		return nil, err
	}

	return wp, nil
}

func (wr *workplanRepository) GetByEmployee(emp entity.Employee) ([]*entity.Workplan, error) {
	var wps []*entity.Workplan

	docRef := wr.cli.Database("locateme").Collection("workplans")
	filter := bson.M{"employee._id": bson.M{"$eq": emp.ID}}
	cursor, err := docRef.Find(context.TODO(), filter)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var workplan entity.Workplan

		if err := cursor.Decode(&workplan); err != nil {
			panic(err)
		}

		wps = append(wps, &workplan)
	}

	return wps, nil
}

func (wr *workplanRepository) Create(wp entity.Workplan) error {
	docRef := wr.cli.Database("locateme").Collection("workplans")
	res, err := docRef.InsertOne(context.TODO(), wp)

	if err != nil {
		return err
	}

	fmt.Println("workplan.create.InsertedID: ", res.InsertedID)

	return nil

}
