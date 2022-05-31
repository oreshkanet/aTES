package accounting

import "github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"

var PaymentFinishedEvent = "payment-finished"

type PaymentFinishedV1 struct {
	PublicId string            `json:"public_id"`
	Dt       string            `json:"dt"`
	User     auth.UserStreamV1 `json:"user"`
	Amount   float32           `json:"amount"`
}
