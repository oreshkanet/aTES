package accounting

type TransactionStreamV1 struct {
	PublicId     string  `json:"public_id"`
	UserPublicId string  `json:"user_public_id"`
	TaskPublicId string  `json:"task_public_id"`
	Credit       float32 `json:"credit"`
	Debit        float32 `json:"debit"`
	Status       uint8   `json:"status"`
}
