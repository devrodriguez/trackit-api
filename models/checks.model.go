package models

type Check struct {
	Address   string  `json:"address,omitempty" binding:"required"`
	Company   string  `json:"company,omitempty" binding:"required"`
	Date      string  `json:"date,omitempty" binding:"required"`
	Email     string  `json:"email,omitempty" binding:"required,email"`
	Hour      string  `json:"hour,omitempty" binding:"required"`
	Latitude  float32 `json:"latitude,omitempty" binding:"required"`
	Longitude float32 `json:"longitude,omitempty" binding:"required"`
	Type      string  `json:"type,omitempty" binding:"required"`
}
