package mngdb

import (
	"context"
	"fmt"

	"github.com/devrodriguez/first-class-api-go/pkg/domain/entity"
	"github.com/devrodriguez/first-class-api-go/pkg/domain/repository"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type checkRepository struct {
	cli *mongo.Client
}

func NewCheckMongoRepo(cli *mongo.Client) repository.CheckRepository {
	return &checkRepository{
		cli,
	}
}

func (cr *checkRepository) DBGetAll() ([]*entity.Check, error) {
	var checks []*entity.Check

	findOptions := options.Find()
	docRef := cr.cli.Database("locateme").Collection("checks")
	cursor, err := docRef.Find(context.TODO(), bson.D{{}}, findOptions)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var check entity.Check

		if err := cursor.Decode(&check); err != nil {
			panic(err)
		}

		checks = append(checks, &check)
	}

	return checks, nil
}

func (cr *checkRepository) DBGetBy(email, companyID, date string) ([]*entity.Check, error) {
	var checks []*entity.Check

	findOptions := options.Find()
	docRef := cr.cli.Database("locateme").Collection("checks")
	filter := bson.M{
		"email":      bson.M{"$eq": email},
		"company_id": bson.M{"$eq": companyID},
		"date":       bson.M{"$eq": date},
	}
	cursor, err := docRef.Find(context.TODO(), filter, findOptions)

	if err != nil {
		return nil, err
	}

	for cursor.Next(context.TODO()) {
		var check entity.Check

		if err := cursor.Decode(&check); err != nil {
			panic(err)
		}

		checks = append(checks, &check)
	}

	return checks, nil
}

func (cr *checkRepository) DBCreate(c *gin.Context, chk entity.Check) error {
	// chk.Company.ID = primitive.NewObjectID()

	docRef := cr.cli.Database("locateme").Collection("checks")
	res, err := docRef.InsertOne(c, chk)

	if err != nil {
		return err
	}

	fmt.Println("Insert ID: ", res.InsertedID)

	return nil
}

func (cr *checkRepository) DBUpdate(id string, chk entity.Check) error {
	docRef := cr.cli.Database("locateme").Collection("checks")
	opts := options.Update().SetUpsert(true)
	hid, err := primitive.ObjectIDFromHex(id)

	// Check for MongoDB ID ObjectIDFromHex errors
	if err != nil {
		panic(err)
	}

	filter := bson.M{"_id": bson.M{"$eq": hid}}

	update := bson.M{
		"$set": chk,
	}

	res, err := docRef.UpdateOne(
		context.Background(),
		filter,
		update,
		opts,
	)

	fmt.Println("ID: ", id, "modified: ", res.ModifiedCount, " matched: ", res.MatchedCount)

	if err != nil {
		return err
	}

	return nil
}
