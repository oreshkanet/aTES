package event

import (
	"context"
	"github.com/oreshkanet/aTES/accounting/internal/service"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/tasktracker"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/schema-registry"
	"github.com/oreshkanet/aTES/packages/pkg/transport/mq"
)

type Consumer struct {
	mq.MessageBroker
	sReg *schemaregistry.EventSchemaRegistry

	cons map[string]mq.Consumer

	uSvc service.UsersService
	tSvc service.TasksService
	aSvc service.AccountService
}

func NewConsumer(
	broker mq.MessageBroker,
	schemaReg *schemaregistry.EventSchemaRegistry,
	usersService service.UsersService,
	taskService service.TasksService,
	accService service.AccountService,
) *Consumer {
	handler := &Consumer{
		MessageBroker: broker,
		sReg:          schemaReg,
		cons:          make(map[string]mq.Consumer),
		uSvc:          usersService,
		tSvc:          taskService,
		aSvc:          accService,
	}

	return handler
}

func (c *Consumer) Run(ctx context.Context) error {
	// User stream
	var err error
	c.cons[auth.UserStreamEvent], err = c.Consume(ctx,
		mq.NewTopic("auth", auth.UserStreamEvent, "1", c.sReg))
	if err != nil {
		return err
	}
	go c.Handle(ctx, auth.UserStreamEvent, c.HandleUserStream)

	// Task stream
	c.cons[tasktracker.TaskStreamEvent], err = c.Consume(ctx,
		mq.NewTopic("tasktracker", tasktracker.TaskStreamEvent, "2", c.sReg))
	if err != nil {
		return err
	}
	go c.Handle(ctx, tasktracker.TaskStreamEvent, c.HandleTaskStream)

	// Task added
	c.cons[tasktracker.TaskAddedEvent], err = c.Consume(ctx,
		mq.NewTopic("tasktracker", tasktracker.TaskAddedEvent, "1", c.sReg))
	if err != nil {
		return err
	}
	go c.Handle(ctx, tasktracker.TaskAddedEvent, c.HandleTaskAdded)

	// Task assigned
	c.cons[tasktracker.TaskAssignedEvent], err = c.Consume(ctx,
		mq.NewTopic("tasktracker", tasktracker.TaskAssignedEvent, "1", c.sReg))
	if err != nil {
		return err
	}
	go c.Handle(ctx, tasktracker.TaskAssignedEvent, c.HandleTaskAssigned)

	// Task done
	c.cons[tasktracker.TaskDoneEvent], err = c.Consume(ctx,
		mq.NewTopic("tasktracker", tasktracker.TaskDoneEvent, "1", c.sReg))
	if err != nil {
		return err
	}
	go c.Handle(ctx, tasktracker.TaskDoneEvent, c.HandleTaskDone)

	return nil
}

func (c *Consumer) Handle(ctx context.Context, eventName string, handler func(rawMessage []byte) error) {
	for {
		select {
		case <-ctx.Done():
			break
		case msg := <-c.cons[eventName].Read():
			if err := handler(msg); err != nil {
				// TODO: пропускать такое сообщение или пытатьсяобрабатывать снова?
				break
			}
		}
	}
}
