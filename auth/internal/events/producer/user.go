package producer

import (
	"context"
	"encoding/json"
	"github.com/oreshkanet/aTES/auth/internal/domain"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message"
	"github.com/oreshkanet/aTES/event-registry/go/pkg/message/auth"
)

func (p *Producer) UserCreated(ctx context.Context, user *domain.User) error {
	traceId := ctx.Value("TraceID").(string)
	userMsg := &auth.UserStreamV1{
		PublicId: user.PublicId,
		Name:     user.Name,
		Role:     user.Role,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Created",
		"1",
		"Auth",
		userMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.producers[auth.UserStreamEvent].Publish(msg)
}

func (p *Producer) UserUpdated(ctx context.Context, user *domain.User) error {
	traceId := ctx.Value("TraceID").(string)
	userMsg := &auth.UserStreamV1{
		PublicId: user.PublicId,
		Name:     user.Name,
		Role:     user.Role,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"Updated",
		"1",
		"Auth",
		userMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.producers[auth.UserStreamEvent].Publish(msg)
}

func (p *Producer) UserRoleChanged(ctx context.Context, user *domain.User) error {
	traceId := ctx.Value("TraceID").(string)
	userMsg := &auth.UserRoleChangedV1{
		PublicId: user.PublicId,
		Role:     user.Role,
	}
	eventMsg := message.NewEventMessage(
		traceId,
		"RoleChanged",
		"1",
		"Auth",
		userMsg,
	)

	msg, err := json.Marshal(eventMsg)
	if err != nil {
		return err
	}

	return p.producers[auth.UserStreamEvent].Publish(msg)
}
