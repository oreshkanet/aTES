package domain

var UserStreamTopic = "auth.user.stream.0"
var UserChangeRoleTopic = "auth.user.changed-role.0"

// TODO: изменение роли может быть бизнес-событием, но пока не реализовано

type User struct {
	Id       string `db:"id"`
	PublicId string `db:"public_id"`
	Name     string `db:"name"`
	Role     string `db:"role"`
}

type UserStreamMessage struct {
	Operation string `json:"operation"`
	PublicId  string `json:"public_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
}
