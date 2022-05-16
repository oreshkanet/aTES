package domain

type Task struct {
	Id         int     `db:"public_id"`
	PublicId   string  `db:"public_id"`
	Title      string  `db:"title"`
	AssignCost float32 `db:"assign_cost"`
	DoneCost   float32 `db:"done_cost"`
}
