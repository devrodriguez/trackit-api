package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `json:"id,omitempty""`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Password string `json:"password"`
}
