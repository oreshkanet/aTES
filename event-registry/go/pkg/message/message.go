package message

import (
	"time"
)

type EventMessage struct {
	EventId      string    `json:"event_id"`
	EventVersion string    `json:"event_version"`
	EventName    string    `json:"event_name"`
	EventTime    time.Time `json:"event_time"`
	Producer     string    `json:"producer"`

	Data interface{} `json:"data"`
}
