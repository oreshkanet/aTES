package domain

var UserStreamTopic = "auth.user.cud.0"

type User struct {
	Id       string `db:"id"`
	PublicId string `db:"public_id"`
	Name     string `db:"name"`
	Role     string `db:"role"`
}

type UserStreamMessage struct {
	Operation string `json:"id"`
	PublicId  string `json:"public_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
}