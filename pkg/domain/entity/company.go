package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Company struct {
	ID   primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name string             `bsdon:"name" json:"name"`
}
