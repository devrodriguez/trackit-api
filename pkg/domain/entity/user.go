package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID       uint   `gorm:"column:id" json:"id,omitempty"`
	Name     string `gorm:"column:name" json:"name"`
	Email    string `gorm:"column:email" json:"email"`
	Password string `gorm:"column:password" json:"password"`
}
