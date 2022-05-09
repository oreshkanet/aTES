package models

type User struct {
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role"`
}
