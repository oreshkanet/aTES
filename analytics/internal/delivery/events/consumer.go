package events

import (
	"context"
	"github.com/oreshkanet/aTES/analytics/internal/service"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
)

type Consumer struct {
	broker mq.MessageBroker

	usersService    service.UsersService
	taskService     service.TasksService
	analyticService service.AnalyticService
}

func NewConsumer(usersService service.UsersService, taskService service.TasksService, analyticService service.AnalyticService) *Consumer {
	handler := &Consumer{
		usersService:    usersService,
		taskService:     taskService,
		analyticService: analyticService,
	}

	return handler
}

func (c *Consumer) Init(ctx context.Context, broker mq.MessageBroker) error {
	// TODO: запускаем косьюминг топиков

	// User stream
	msgCh := make(chan []byte)
	go broker.Consume(ctx, message.GetTopicName("auth", "user-stream", "1"), msgCh)
	go c.Handle(ctx, msgCh, c.HandleUserStream)

	// Task stream
	taskStreamCh := make(chan []byte)
	go broker.Consume(ctx, message.GetTopicName("task-tracker", "task-stream", "1"), taskStreamCh)
	go c.Handle(ctx, taskStreamCh, c.HandleTaskStream)

	// Task added
	/*
		taskAddedCh := make(chan []byte)
		go broker.Consume(ctx, message.GetTopicName("task-tracker", "task-add", "1"), taskAddedCh)
		go c.Handle(ctx, taskAddedCh, c.HandleTaskAdded)

		// Task assigned
		taskAssignedCh := make(chan []byte)
		go broker.Consume(ctx, message.GetTopicName("task-tracker", "task-assigned", "1"), taskAssignedCh)
		go c.Handle(ctx, taskAssignedCh, c.HandleTaskAssigned)

		// Task done
		taskDoneCh := make(chan []byte)
		go broker.Consume(ctx, message.GetTopicName("task-tracker", "task-done", "1"), taskDoneCh)
		go c.Handle(ctx, taskDoneCh, c.HandleTaskDone)

	*/

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
