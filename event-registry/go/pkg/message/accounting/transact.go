package accounting

var TransactStreamTopic = "account.transact-stream.0"
var TransactPaymentTopic = "account.transact-payment.0"

type TransactStreamMessageV1 struct {
	PublicId       string  `json:"public_id"`
	UserPublicId   string  `json:"user_public_id"`
	TaskPublicId   string  `json:"task_public_id"`
	AmountIncrease float32 `json:"amount_increase"`
	AmountDecrease float32 `json:"amount_decrease"`
	Comment        string  `json:"comment"`
}

type TransactPaymentMessageV1 struct {
	PublicId     string  `json:"public_id"`
	UserPublicId string  `json:"user_public_id"`
	TaskPublicId string  `json:"task_public_id"`
	TaskTitle    string  `json:"task_title"`
	Amount       float32 `json:"amount"`
	Comment      string  `json:"comment"`
}
