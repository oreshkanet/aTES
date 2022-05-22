package accounting

var TaskCostCalculatedEvent = "task-cost-calculated"

type TaskCostCalculatedV1 struct {
	TaskPublicId string  `json:"task_public_id"`
	AssignCost   float32 `json:"assign_cost"`
	DoneCost     float32 `json:"done_cost"`
}
