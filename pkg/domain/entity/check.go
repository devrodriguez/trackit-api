package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Check struct {
	MgID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Address   string             `bson:"address" json:"address,omitempty" binding:"required"`
	Company   string             `bson:"company" json:"company,omitempty" binding:"required"`
	CompanyID string             `bson:"company_id" json:"company_id,omitempty" binding:"required"`
	Date      string             `bson:"date" json:"date,omitempty" binding:"required"`
	Email     string             `bson:"email" json:"email,omitempty" binding:"required,email"`
	Hour      string             `bson:"hour" json:"hour,omitempty" binding:"required"`
	Latitude  float32            `bson:"latitude" json:"latitude,omitempty" binding:"required"`
	Longitude float32            `bson:"longitude" json:"longitude,omitempty" binding:"required"`
	Type      string             `bson:"type" json:"type,omitempty" binding:"required"`
}
