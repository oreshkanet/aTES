package domain

import (
	"golang.org/x/text/currency"
)

var TransactStreamTopic = "account.transact.stream.0"
var TransactPaymentTopic = "account.transact.payment.0"

type Transact struct {
	Id             int `db:"id"`
	PublicId       string
	UserId         int `db:"user_id"`
	UserPublicId   string
	TaskId         int `db:"task_id"`
	TaskPublicId   string
	AmountIncrease currency.Amount `db:"amount_increase"`
	AmountDecrease currency.Amount `db:"amount_decrease"`
	Comment        string          `db:"comment"`
}

type TransactStreamMessage struct {
	PublicId       string          `json:"public_id"`
	UserPublicId   string          `json:"user_public_id"`
	TaskPublicId   string          `json:"task_public_id"`
	AmountIncrease currency.Amount `json:"amount_increase"`
	AmountDecrease currency.Amount `json:"amount_decrease"`
	Comment        string          `json:"comment"`
}

type TransactPaymentMessage struct {
	PublicId     string          `json:"public_id"`
	UserPublicId string          `json:"user_public_id"`
	TaskPublicId string          `json:"task_public_id"`
	TaskTitle    string          `json:"task_title"`
	Amount       currency.Amount `json:"amount"`
	Comment      string          `json:"comment"`
}
