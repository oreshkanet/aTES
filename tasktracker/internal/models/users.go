package models

type User struct {
	id   string `json:"id" db:"id"`
	name string `json:"name" db:"name"`
	role string `json:"role" db:"role"`
}
