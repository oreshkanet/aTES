package domain

import "time"

var TransactStreamTopic = "account.transact.stream.0"
var TransactPaymentTopic = "account.transact.payment.0"

type Transaction struct {
	Id           int       `db:"id"`
	Dt           time.Time `db:"dt"`
	PublicId     string    `db:"public_id"`
	UserId       int       `db:"user_id"`
	UserPublicId string
	TaskId       int `db:"task_id"`
	TaskPublicId string
	Credit       float32 `db:"credit"`
	Debit        float32 `db:"debit"`
	Status       uint8   `db:"status"`
}

type Payment struct {
	Id           int       `db:"id"`
	Dt           time.Time `db:"dt"`
	PublicId     string    `db:"public_id"`
	UserId       string    `db:"user_id"`
	UserPublicId string
	Amount       float32 `db:"amount"`
	Status       uint8   `db:"status"`
}
