package entity

import "go.mongodb.org/mongo-driver/bson/primitive"

type Employee struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name,omitempty" json:"name"`
	Email        string             `bson:"email,omitempty" json:"email,omitempty"`
	Position     string             `bson:"position,omitempty" json:"position,omitempty"`
	InternalUser User               `bson:"internal_user,omitempty" json:"internal_user,omitempty"`
}
