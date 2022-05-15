package events

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/domain"
	"github.com/oreshkanet/aTES/accounting/internal/services"
	"github.com/oreshkanet/aTES/accounting/internal/transport/mq"
)

type Consumer struct {
	broker mq.MessageBroker

	usersService services.UsersService
	taskService  services.TasksService
	accService   services.AccountService
}

func NewConsumer(usersService services.UsersService, taskService services.TasksService, accService services.AccountService) *Consumer {
	handler := &Consumer{
		usersService: usersService,
		taskService:  taskService,
		accService:   accService,
	}

	return handler
}

func (c *Consumer) Init(ctx context.Context, broker mq.MessageBroker) error {
	// TODO: запускаем косьюминг топиков

	// User stream
	msgCh := make(chan []byte)
	go broker.Consume(ctx, domain.UserStreamTopic, msgCh)
	go c.Handle(ctx, msgCh, c.HandleUserStream)

	// Task stream
	taskStreamCh := make(chan []byte)
	go broker.Consume(ctx, domain.TaskStreamTopic, taskStreamCh)
	go c.Handle(ctx, taskStreamCh, c.HandleTaskStream)

	// Task added
	taskAddedCh := make(chan []byte)
	go broker.Consume(ctx, domain.TaskAddedTopic, taskAddedCh)
	go c.Handle(ctx, taskAddedCh, c.HandleTaskAdded)

	// Task assigned
	taskAssignedCh := make(chan []byte)
	go broker.Consume(ctx, domain.TaskAssignedTopic, taskAssignedCh)
	go c.Handle(ctx, taskAssignedCh, c.HandleTaskAssigned)

	// Task done
	taskDoneCh := make(chan []byte)
	go broker.Consume(ctx, domain.TaskDoneTopic, taskDoneCh)
	go c.Handle(ctx, taskDoneCh, c.HandleTaskDone)

	return nil
}

func (c *Consumer) Handle(ctx context.Context, msgCh <-chan []byte, handleMessages func(rawMessage []byte) error) {
	for {
		select {
		case <-ctx.Done():
			break
		case msg := <-msgCh:
			if err := handleMessages(msg); err != nil {
				// TODO: пропускать такое сообщение или пытатьсяобрабатывать снова?
				break
			}
		}
	}
}
