package domain

type Task struct {
	Id           int    `db:"public_id"`
	PublicId     string `db:"public_id"`
	JiraId       string `db:"jira_id"`
	Title        string `db:"title"`
	Description  string `db:"description"`
	AssignedUser string `db:"assigned_user"`
}
