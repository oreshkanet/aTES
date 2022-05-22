package auth

var (
	UserStreamEvent      = "user-stream"
	UserRoleChangedEvent = "user-role-changed"
)

type UserStreamV1 struct {
	PublicId string `json:"public_id"`
	Name     string `json:"name"`
	Role     string `json:"role"`
}

type UserRoleChangedV1 struct {
	PublicId string `json:"public_id"`
	Role     string `json:"role"`
}
