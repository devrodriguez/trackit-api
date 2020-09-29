package entity

type User struct {
	Email    string `bson:"email" json:"email"`
	Name     string `bson:"name" json:"name"`
	Password string `bson:"password" json:"password,omitempty"`
}
