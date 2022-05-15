package domain

var TaskStreamTopic = "tasks.task.stream.0"
var TaskAddedTopic = "tasks.task.added.0"
var TaskAssignedTopic = "tasks.task.assigned.0"
var TaskDoneTopic = "tasks.task.done.0"

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

type TaskAddedMessage struct {
	PublicId string `json:"public_id"`
}

type TaskAssignedMessage struct {
	PublicId     string `json:"public_id"`
	UserPublicId string `json:"user_public_id"`
}

type TaskDoneMessage struct {
	PublicId     string `json:"public_id"`
	UserPublicId string `json:"user_public_id"`
}
