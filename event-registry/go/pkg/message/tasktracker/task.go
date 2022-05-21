package tasktracker

type TaskStreamMessageV1 struct {
	Operation   string `json:"operation"`
	PublicId    string `json:"public_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type TaskStreamMessageV2 struct {
	Operation   string `json:"operation"`
	PublicId    string `json:"public_id"`
	JiraId      string `json:"jira_id"`
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
