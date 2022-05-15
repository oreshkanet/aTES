package models

type User struct {
	PublicId string `db:"public_id" json:"publicId"`
	Name     string `db:"name" json:"name"`
	Password string `db:"password" json:"password"`
	Role     string `db:"role"`
}
