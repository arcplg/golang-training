package services

type User struct {
	ID   string
	Name string
}

func GetAllUsers() []User {
	return []User{
		{ID: "1", Name: "Alice"},
		{ID: "2", Name: "Bob"},
	}
}

func GetUserByID(id string) *User {
	users := GetAllUsers()
	for _, user := range users {
		if user.ID == id {
			return &user
		}
	}
	return nil
}
