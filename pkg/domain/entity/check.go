package entity

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
)

type Check struct {
	gorm.Model
	ID          uint    `json:"id,omitempty"`
	Name        string  `json:"name"`
	Address     string  `json:"address"`
	Date        string  `json:"date"`
	Hour        string  `json:"hour"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	CompanyID   uint    `json:"company_id"`
	EmployeeID  uint    `json:"employee_id"`
	CheckTypeID uint    `json:"check_type_id"`
}

type CheckDB struct {
	MgID      primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Address   string             `bson:"address" json:"address,omitempty" binding:"required"`
	Company   string             `bson:"company" json:"company,omitempty" binding:"required"`
	Date      string             `bson:"date" json:"date,omitempty" binding:"required"`
	Email     string             `bson:"email" json:"email,omitempty" binding:"required,email"`
	Hour      string             `bson:"hour" json:"hour,omitempty" binding:"required"`
	Latitude  float32            `bson:"latitude" json:"latitude,omitempty" binding:"required"`
	Longitude float32            `bson:"longitude" json:"longitude,omitempty" binding:"required"`
	Type      string             `bson:"type" json:"type,omitempty" binding:"required"`
}
