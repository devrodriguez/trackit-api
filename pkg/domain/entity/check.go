package entity

import (
	"gorm.io/gorm"
)

type Check struct {
	gorm.Model
	ID          uint    `json:"id,omitempty"`
	Description string  `json:"name"`
	Address     string  `json:"address"`
	Date        string  `json:"date"`
	Hour        string  `json:"hour"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
	CompanyID   uint    `json:"company_id"`
	EmployeeID  uint    `json:"employee_id"`
	CheckTypeID uint    `json:"check_type_id"`
}

type CheckType struct {
	gorm.Model
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name"`
}

type CheckQuery struct {
	Address       string `gorm:"column:address" json:"address"`
	CheckID       uint   `gorm:"column:check_id" json:"check_id"`
	CompanyID     uint   `gorm:"column:company_id" json:"company_id"`
	CompanyName   string `gorm:"column:company_name" json:"company_name"`
	EmployeeID    uint   `gorm:"column:employee_id" json:"employee_id"`
	EmployeeName  string `gorm:"column:employee_name" json:"employee_name"`
	CheckTypeID   uint   `gorm:"column:check_type_id" json:"check_type_id"`
	CheckTypeName string `gorm:"column:check_type_name" json:"check_type_name"`
	Date          string `gorm:"column:date" json:"date"`
	Hour          string `gorm:"hour" json:"hour"`
	Latitude      string `gorm:"latitude" json:"latitude"`
	Longitude     string `gorm:"longitude" json:"longitude"`
	Description   string `gorm:"column:description" json:"description"`
}
