package entity

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Workplan struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Company   Company            `bson:"company" json:"company" binding:"required"`
	Address   string             `bson:"address" json:"address" binding:"required"`
	Latitude  float64            `bson:"latitude" json:"latitude" binding:"required"`
	Longitude float64            `bson:"longitude" json:"longitude" binding:"required"`
	StartTime time.Time          `bson:"start_time" json:"start_time" binding:"required"`
	EndTime   time.Time          `bson:"end_time" json:"end_time" binding:"required"`
	Employee  Employee           `bson:"employee" json:"employee" binding:"required"`
	Date      string             `bson:"date" json:"date" binding:"required"`
	Task      string             `bson:"task" json:"task"`
	Status    string             `bson:"status" json:"status"`
}
