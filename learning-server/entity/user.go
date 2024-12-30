package entity

type User struct {
	ID    string `json:"id" bson:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
