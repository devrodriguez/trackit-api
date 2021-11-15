package entity

import "gorm.io/gorm"

type Company struct {
	gorm.Model
	ID   uint   `json:"id,omitempty"`
	Name string `json:"name,omitempty""`
}
