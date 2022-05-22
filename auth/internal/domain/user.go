package domain

type User struct {
	PublicId string `db:"public_id"`
	Name     string `db:"name"`
	Password string `db:"password"`
	Role     string `db:"role"`
}
