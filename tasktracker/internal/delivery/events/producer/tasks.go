package producer

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/tasktracker"
	"github.com/oreshkanet/aTES/tasktracker/internal/domain"
)

func (p *Producer) TaskCreate(ctx context.Context, task *domain.Task) error {
	traceId := ctx.Value("TraceID").(string)
	taskMsg := &tasktracker.TaskStreamMessageV2{
		PublicId:    task.PublicId,
		JiraId:      task.JiraId,
		Title:       task.Title,
		Description: task.Description,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Created",
		"2",
		"TaskTracker",
		taskMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.producers[tasktracker.TaskStreamEvent].Publish(msg)
}

func (p *Producer) TaskUpdate(ctx context.Context, task *domain.Task) error {
	traceId := ctx.Value("TraceID").(string)
	taskMsg := &tasktracker.TaskStreamMessageV2{
		PublicId:    task.PublicId,
		JiraId:      task.JiraId,
		Title:       task.Title,
		Description: task.Description,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Updated",
		"2",
		"TaskTracker",
		taskMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.producers[tasktracker.TaskStreamEvent].Publish(msg)
}

func (p *Producer) TaskAdded(ctx context.Context, task *domain.Task) error {
	traceId := ctx.Value("TraceID").(string)
	taskMsg := &tasktracker.TaskAddedMessageV1{
		PublicId: task.PublicId,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Added",
		"1",
		"TaskTracker",
		taskMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.producers[tasktracker.TaskAddedEvent].Publish(msg)
}

func (p *Producer) TaskAssigned(ctx context.Context, task *domain.Task) error {
	traceId := ctx.Value("TraceID").(string)
	taskMsg := &tasktracker.TaskAssignedMessageV1{
		PublicId:    task.PublicId,
		UserPublicId: task.AssignedUser
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Assigned",
		"1",
		"TaskTracker",
		taskMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.producers[tasktracker.TaskAssignedEvent].Publish(msg)
}

func (p *Producer) TaskDone(ctx context.Context, task *domain.Task) error {
	traceId := ctx.Value("TraceID").(string)
	taskMsg := &tasktracker.TaskDoneMessageV1{
		PublicId:    task.PublicId,
		UserPublicId: task.AssignedUser,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Done",
		"1",
		"TaskTracker",
		taskMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.producers[tasktracker.TaskDoneEvent].Publish(msg)
}
