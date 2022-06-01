package domain

type Task struct {
	Id          int     `db:"id"`
	PublicId    string  `db:"public_id"`
	Title       string  `db:"title"`
	Description string  `db:"description"`
	AssignCost  float32 `db:"assign_cost"`
	DoneCost    float32 `db:"done_cost"`
}
