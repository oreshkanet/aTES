package domain

var TaskStreamTopic = "task.stream.0"

type Task struct {
	Id           int    `db:"public_id"`
	PublicId     string `db:"public_id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	AssignedUser string `db:"assigned_user"`
}

type TaskStreamMessage struct {
	Operation    string `json:"operation"`
	PublicId     string `json:"public_id"`
	Title        string `json:"title"`
	Description  string `json:"description"`
	AssignedUser string `json:"assigned_user"`
}

type TaskAddMessage struct {
	PublicId     string `json:"public_id"`
	AssignedUser string `json:"assigned_user"`
}
