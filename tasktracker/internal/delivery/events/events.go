package events

import (
	"context"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/tasktracker"
)

type Producer interface {
	TaskCreate(ctx context.Context, message *tasktracker.TaskStreamMessageV2) error
	TaskUpdate(ctx context.Context, message *tasktracker.TaskStreamMessageV2) error
	TaskAdded(ctx context.Context, message *tasktracker.TaskAddedMessageV1) error
	TaskAssigned(ctx context.Context, message *tasktracker.TaskAssignedMessageV1) error
	TaskDone(ctx context.Context, message *tasktracker.TaskDoneMessageV1) error
}
