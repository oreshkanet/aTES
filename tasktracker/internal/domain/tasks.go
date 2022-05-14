package domain

var TaskStreamTopic = "tasks.task.stream.0"

type Task struct {
	Id           int    `db:"public_id"`
	PublicId     string `db:"public_id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	AssignedUser string `db:"assigned_user"`
}

type TaskStreamMessage struct {
	Operation   string `json:"operation"`
	PublicId    string `json:"public_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskAddMessage struct {
	PublicId string `json:"public_id"`
}

type TaskAssignMessage struct {
	PublicId             string `json:"public_id"`
	AssignedUserPublicId string `json:"assigned_user_public_id"`
}
