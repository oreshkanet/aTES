package domain

type User struct {
	Id       string  `db:"id"`
	PublicId string  `db:"public_id"`
	Name     string  `db:"name"`
	Role     string  `db:"role"`
	Balance  float32 `db:"balance"`
}
