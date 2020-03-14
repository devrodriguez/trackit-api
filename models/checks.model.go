package models

type Check struct {
	Address   string  `json:"address,omitempty"`
	Company   string  `json:"company"`
	Date      string  `json:"date"`
	Email     string  `json:"email"`
	Hour      string  `json:"hour"`
	Latitude  float32 `json:"latitude"`
	Longitude float32 `json:"longitude"`
	Type      string  `json:"type"`
}
