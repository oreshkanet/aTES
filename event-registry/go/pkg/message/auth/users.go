package auth

var UserStreamTopic = "auth.user-stream.v1"
var UserChangeRoleTopic = "auth.user-changed-role.v1"

type UserStreamV1 struct {
	Operation string `json:"operation"`
	PublicId  string `json:"public_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
}
