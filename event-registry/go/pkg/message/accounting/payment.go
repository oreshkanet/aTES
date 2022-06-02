package accounting

import "time"

var PaymentFinishedEvent = "payment-finished"

type PaymentFinishedV1 struct {
	PublicId     string    `json:"public_id"`
	Dt           time.Time `json:"dt"`
	UserPublicId string    `json:"user_public_id"`
	Amount       float32   `json:"amount"`
}
