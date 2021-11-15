package entity

import "gorm.io/gorm"

type Employee struct {
	gorm.Model
	ID       uint   `json:"id,omitempty"`
	Name     string `json:"name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
}
