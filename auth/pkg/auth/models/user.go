package models

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Role     string
}