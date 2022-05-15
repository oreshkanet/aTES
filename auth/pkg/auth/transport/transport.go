package transport

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/oreshkanet/aTES/auth/pkg/auth/models"
	"github.com/oreshkanet/aTES/auth/pkg/kafka"
)

var TopicUserStream = "auth.user.stream.0"

type UserStreamMessage struct {
	Operation string `json:"operation"`
	PublicId  string `json:"public_id"`
	Name      string `json:"name"`
	Role      string `json:"role"`
}

type KafkaMessage struct {
	Topic     string
	Routing   string
	Payload   []byte
	Partition int
	Key       string
}

type Transports struct {
	UserTopic *kafka.Producer
}

func (q *Transports) Close() {
	q.UserTopic.Close()
}

func (q *Transports) PubUser(ctx context.Context, user *models.User) error {
	userMessage := &UserStreamMessage{
		Operation: "C",
		PublicId:  user.PublicId,
		Name:      user.Name,
		Role:      user.Role,
	}
	message, err := json.Marshal(userMessage)
	if err != nil {
		return fmt.Errorf("Pub User: %v", err)
	}
	return q.UserTopic.PubMessage(ctx, user.Name, string(message))
}

func CreateTransport(kafkaAddress string) *Transports {
	userConn := kafka.NewProducer(kafkaAddress, TopicUserStream)

	return &Transports{
		UserTopic: userConn,
	}
}
