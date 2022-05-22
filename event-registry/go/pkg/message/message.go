package message

import (
	"github.com/google/uuid"
	"time"
)

type EventMessage struct {
	TraceId      string    `json:"trace_id"`
	EventId      string    `json:"event_id"`
	EventName    string    `json:"event_name"`
	EventVersion string    `json:"event_version"`
	EventTime    time.Time `json:"event_time"`
	Producer     string    `json:"producer"`

	Data interface{} `json:"data"`
}

func NewEventMessage(traceId string, name string, version string, producer string, data interface{}) *EventMessage {
	return &EventMessage{
		TraceId:      traceId,
		EventId:      uuid.New().String(),
		EventName:    name,
		EventVersion: version,
		EventTime:    time.Now(),
		Producer:     producer,
		Data:         data,
	}
}
