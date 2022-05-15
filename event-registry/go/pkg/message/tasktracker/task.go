package tasktracker

var TaskStreamTopic = "task-tracker.task-stream.v1"
var TaskAddedTopic = "task-tracker.task-added.v1"
var TaskAssignedTopic = "task-tracker.task-assigned.v1"
var TaskDoneTopic = "task-tracker.task-done.v1"

type TaskStreamMessageV1 struct {
	Operation   string `json:"operation"`
	PublicId    string `json:"public_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskAddedMessageV1 struct {
	PublicId string `json:"public_id"`
}

type TaskAssignedMessageV1 struct {
	PublicId     string `json:"public_id"`
	UserPublicId string `json:"user_public_id"`
}

type TaskDoneMessageV1 struct {
	PublicId     string `json:"public_id"`
	UserPublicId string `json:"user_public_id"`
}
