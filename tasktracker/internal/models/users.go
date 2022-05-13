package models

type User struct {
	Id       string `db:"id"`
	PublicId string `json:"public_id" db:"public_id"`
	Name     string `json:"name" db:"name"`
	Role     string `json:"role" db:"role"`
}
