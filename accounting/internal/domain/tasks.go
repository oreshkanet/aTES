package domain

var TaskStreamTopic = "task-tracker.task.stream.0"
var TaskAddedTopic = "task-tracker.task.added.0"
var TaskAssignedTopic = "task-tracker.task.assigned.0"
var TaskDoneTopic = "task-tracker.task.done.0"

type Task struct {
	Id         int     `db:"public_id"`
	PublicId   string  `db:"public_id"`
	Title      string  `db:"title"`
	AssignCost float32 `db:"assign_cost"`
	DoneCost   float32 `db:"done_cost"`
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
