package model

type User struct {
	Name string `json:"name"`
}

func NewUser() *User {
	return &User{
		Name: "",
	}
}
